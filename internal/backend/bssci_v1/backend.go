package bssci_v1

import (
	"context"
	"crypto/tls"
	"crypto/x509"

	"net"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs/messages"
	"mioty-bssci-adapter/internal/common"
	"mioty-bssci-adapter/internal/config"
)

type Backend struct {
	sync.RWMutex

	caCert  string
	tlsCert string
	tlsKey  string

	listener net.Listener
	scheme   string
	isClosed bool

	basestations basestations

	statsInterval time.Duration
	pingInterval  time.Duration
	readTimeout   time.Duration
	writeTimeout  time.Duration

	// downlinkTxAckFunc           func(*gw.DownlinkTxAck)
	// uplinkFrameFunc             func(*gw.UplinkFrame)
	// basestationStatsFunc        func(*gw.GatewayStats)
	// rawPacketForwarderEventFunc func(*gw.RawPacketForwarderEvent)

	// Cache to store diid to UUIDs.
	diidCache *cache.Cache
}

// NewBackend creates a new Backend.
func NewBackend(conf config.Config) (backend *Backend, err error) {
	b := Backend{
		scheme: "ssl",

		basestations: basestations{
			basestations: make(map[common.EUI64]*connection),
		},

		caCert:  conf.Backend.BssciV1.CACert,
		tlsCert: conf.Backend.BssciV1.TLSCert,
		tlsKey:  conf.Backend.BssciV1.TLSKey,

		statsInterval: conf.Backend.BssciV1.StatsInterval,
		pingInterval:  conf.Backend.BssciV1.PingInterval,
		readTimeout:   conf.Backend.BssciV1.ReadTimeout,
		writeTimeout:  conf.Backend.BssciV1.WriteTimeout,

		diidCache: cache.New(time.Minute, time.Minute),
	}

	// create the listener
	b.listener, err = net.Listen("tcp", conf.Backend.BssciV1.Bind)
	if err != nil {
		return nil, errors.Wrap(err, "create listener error")
	}

	// if the CA and TLS cert is configured, setup client certificate verification.
	if b.tlsCert == "" && b.tlsKey == "" && b.caCert == "" {
		rawCACert, err := os.ReadFile(b.caCert)
		if err != nil {
			return nil, errors.Wrap(err, "read ca cert error")
		}
		tlsCert, err := tls.LoadX509KeyPair(b.tlsCert, b.tlsKey)
		if err != nil {
			return nil, errors.Wrap(err, "read tls cert error")
		}

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(rawCACert)

		// wrap the tcp listener in a tls listener
		b.listener = tls.NewListener(b.listener, &tls.Config{
			Certificates: []tls.Certificate{tlsCert},
			ClientCAs:    caCertPool,
			ClientAuth:   tls.RequireAndVerifyClientCert,
		})

	}
	backend = &b
	return
}

// Stop stops the backend.
func (b *Backend) Stop() error {
	b.isClosed = true
	return b.listener.Close()
}

// Start starts the backend.
func (b *Backend) Start() error {

	go func() {
		for !b.isClosed {
			// accept a new connection
			conn, err := b.listener.Accept()

			if err != nil {
				log.Error().Stack().Err(err).Msg("tls accept failed")

			}
			logger := log.With().Str("remote", conn.RemoteAddr().String()).Logger()
			logger.Info().Msg("accepted new tls connection")

			// try to read Con message
			conn.SetReadDeadline(time.Now().Add(b.readTimeout))
			cmdHeader, raw, err := ReadBssciMessage(conn)

			if err != nil {
				logger.Error().Stack().Err(err).Msg("codec error")
				conn.Close()
			} else {
				// first message after connecting should always be Con
				cmd := cmdHeader.GetCommand()

				if cmd == structs.MsgCon {

					var con messages.Con
					_, err = con.UnmarshalMsg(raw)
					if err != nil {
						logger.Error().Stack().Err(err).Str("command", string(cmd)).Msg("unmarshal msgp error")
						conn.Close()
					} else {
						logger.Info().Str("command", string(cmd)).Msg("initializing basestation connection")
						ctx := context.Background()
						ctx = logger.WithContext(ctx)

						b.initBasestation(ctx, con, conn, b.handleBasestation)
					}
				} else {
					logger.Warn().Str("command", string(cmd)).Msg("expected con command")

					conn.Close()
				}
			}
		}
	}()
	return nil
}

func (b *Backend) initBasestation(ctx context.Context, con messages.Con, conn net.Conn, handler func(ctx context.Context, eui common.EUI64, conn *connection)) {
	defer conn.Close()

	eui := con.GetEui()

	logger := zerolog.Ctx(ctx)
	logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
		return c.Str("gw_eui", eui.String())
	})

	newUuid := uuid.New()
	bsConnection := connection{
		conn: conn,
		// stats:      stats.NewCollector(),
		lastActive: time.Now(),
		opId:       -1,
		version:    con.Version,
		SnBsUuid:   con.SnBsUuid.ToUuid(),
		SnScUuid:   newUuid,
	}

	conRsp := messages.NewConRsp(con.OpId, con.Version, newUuid)

	// check for existing connection
	oldConnection, err := b.basestations.get(eui)
	if err == nil && oldConnection != nil {
		// found existing connection for this basestation,
		if oldConnection.SnBsUuid == con.SnBsUuid.ToUuid() {
			// check if session uuid matched with con request
			logger.Info().Msg("resuming previous basestation connection")
			bsConnection.SnScUuid = oldConnection.SnScUuid

			if con.SnScOpId != nil {
				bsConnection.opId = *con.SnScOpId
			}
			conRsp.ResumeConnection(oldConnection.SnScUuid)

		} else {
			// remove old connection if session uuid does not match
			logger.Warn().Msg("removing previous basestation connection")
			_ = b.basestations.remove(eui)
		}
	}

	// set the gateway connection
	if err := b.basestations.set(eui, &bsConnection); err != nil {
		logger.Error().Stack().Err(err).Msg("failed to set connection")
	}

	logger.Info().Msg("basestation connected")

	done := make(chan struct{})

	// remove the basestation on return
	defer func() {
		done <- struct{}{}
		b.basestations.remove(eui)
		logger.Info().Msg("basestation disconnected")
	}()

	pingTicker := time.NewTicker(b.pingInterval)
	defer pingTicker.Stop()
	statusTicker := time.NewTicker(b.statsInterval)
	defer pingTicker.Stop()

	go func() {
		for {
			select {
			case <-pingTicker.C:
				opId := bsConnection.DecrementOpId()
				msg := messages.NewPing(opId)
				err := bsConnection.Write(&msg, b.writeTimeout)

				if err != nil {
					logger.Error().Stack().Err(err).Str("command", string(msg.GetCommand())).Msg("failed to send scheduled ping message")
					bsConnection.conn.Close()
					return
				}
				logger.Info().Msg("sent scheduled ping message")
			case <-statusTicker.C:
				opId := bsConnection.DecrementOpId()
				msg := messages.NewStatus(opId)
				err := bsConnection.Write(&msg, b.writeTimeout)

				if err != nil {
					logger.Error().Stack().Err(err).Str("command", string(msg.GetCommand())).Msg("failed to send scheduled status message")
					bsConnection.conn.Close()
					return
				}
				logger.Info().Msg("sent scheduled status message")
			case <-done:
				return
			}
		}
	}()

	connectCounter().Inc()
	//send ConRsp
	bsConnection.Write(&conRsp, b.writeTimeout)

	handler(ctx, eui, &bsConnection)
	done <- struct{}{}

}

// handle all messages coming from a client
func (b *Backend) handleBasestation(ctx context.Context, eui common.EUI64, conn *connection) {
	logger := zerolog.Ctx(ctx)
	for {
		var response messages.Message
		cmdHeader, raw, err := conn.Read(b.readTimeout)

		if err != nil {
			logger.Error().Stack().Err(err).Msg("failed to read message")
			// terminate this connection
			return
		}

		opId := cmdHeader.GetOpId()
		cmd := cmdHeader.GetCommand()

		// update logging context
		logger := log.With().Str("command", string(cmd)).Int64("op_id", opId).Logger()

		if logger.Debug().Enabled() {
			logger.Debug().Bytes("msgp", raw).Msg("received message")
		} else {
			logger.Info().Msg("received message")
		}

		// only match ClientMsg... messages
		switch cmd {
		case structs.ClientMsgCon:
			// handle con message
			var msg messages.Con
			_, err = msg.UnmarshalMsg(raw)
			if err != nil {
				logger.Error().Stack().Err(err).Msg("unmarshal msgp error")
				continue
			}
			response = b.handleConMessage(ctx, conn, eui, msg)
		case structs.ClientMsgAtt:
			// handle attach message
			var msg messages.Att
			_, err = msg.UnmarshalMsg(raw)
			if err != nil {
				logger.Error().Stack().Err(err).Msg("unmarshal msgp error")
				continue
			}
			response = b.handleAttMessage(ctx, eui, msg)
		case structs.ClientMsgDet:
			// handle detach message
			var msg messages.Det
			_, err = msg.UnmarshalMsg(raw)
			if err != nil {
				logger.Error().Stack().Err(err).Msg("unmarshal msgp error")
				continue
			}
			response = b.handleDetMessage(ctx, eui, msg)
		case structs.ClientMsgUlData:
			// handle uplink data message
			var msg messages.UlData
			_, err = msg.UnmarshalMsg(raw)
			if err != nil {
				logger.Error().Stack().Err(err).Msg("unmarshal msgp error")
				continue
			}

			response = b.handleUlDataMessage(ctx, eui, msg)
		case structs.ClientMsgDlDataRes:
			// handle downlink data result response
			var msg messages.DlDataRes
			_, err = msg.UnmarshalMsg(raw)
			if err != nil {
				logger.Error().Stack().Err(err).Msg("unmarshal msgp error")
				continue
			}
			response = b.handleDlDataResMessage(ctx, msg, eui)

		case structs.ClientMsgDlRxStat:
			// handle downlink rx status data message
			var msg messages.DlRxStat
			_, err = msg.UnmarshalMsg(raw)
			if err != nil {
				logger.Error().Stack().Err(err).Msg("unmarshal msgp error")
				continue
			}
			response = b.handleDlRxStatMessage(ctx, eui, msg)
		case structs.ClientMsgStatusRsp:
			// handle status response message
			var msg messages.StatusRsp
			_, err = msg.UnmarshalMsg(raw)
			if err != nil {
				logger.Error().Stack().Err(err).Msg("unmarshal msgp error")
				continue
			}
			response = b.handleStatusRspMessage(ctx, eui, msg)
		case structs.ClientMsgPing:
			// handle ping message
			defaultResponse := messages.NewPingRsp(opId)
			response = &defaultResponse
		case structs.ClientMsgPingRsp:
			// handle ping response (pong) message
			defaultResponse := messages.NewPingCmp(opId)
			response = &defaultResponse
		case structs.ClientMsgDlDataRevRsp:
			// handle downlink data revoke response message
			defaultResponse := messages.NewDlDataRevCmp(opId)
			response = &defaultResponse
		case structs.ClientMsgDlDataQueRsp:
			// handle downlink data queue response message
			defaultResponse := messages.NewDlDataQueCmp(opId)
			response = &defaultResponse
		case structs.ClientMsgDlRxStatQryRsp:
			// handle downlink rx status query response message
			defaultResponse := messages.NewDlRxStatQryCmp(opId)
			response = &defaultResponse
		case structs.ClientMsgAttPrpRsp:
			// handle attach propagate response message
			defaultResponse := messages.NewAttPrpCmp(opId)
			response = &defaultResponse
		case structs.ClientMsgDetPrpRsp:
			// handle detach propagate response message
			defaultResponse := messages.NewDetPrpCmp(opId)
			response = &defaultResponse
		case structs.ClientMsgError:
			// handle error message
			var msg messages.BssciError
			_, err = msg.UnmarshalMsg(raw)
			if err != nil {
				logger.Error().Stack().Err(err).Msg("unmarshal msgp error")
				continue
			}
			logger.Warn().Uint32("err_code", msg.Code).Str("err_msg", msg.Message).Msg("received bssci error message")
			defaultResponse := messages.NewBssciErrorAck(opId)
			response = &defaultResponse
		case structs.ClientMsgErrorAck:
			// Equivalent to ...Cmp message
			continue
		case structs.ClientMsgPingCmp:
			// ...Cmp messages need no further handling
			continue
		case structs.ClientMsgUlDataCmp:
			continue
		case structs.ClientMsgDlDataResCmp:
			continue
		case structs.ClientMsgAttCmp:
			continue
		case structs.ClientMsgConCmp:
			continue
		case structs.ClientMsgDlRxStatCmp:
			continue
		case structs.ClientMsgDetCmp:
			continue
		default:
			logger.Warn().Msg("received unsupported command")

			// maybe send Error Message?
			continue
		}

		if response != nil {
			err := conn.Write(response, b.writeTimeout)
			if err != nil {
				logger.Error().Stack().Err(err).Msg("failed to write message")
				// terminate this connection
				return
			}
			logger.Info().Str("response", string(response.GetCommand())).Msg("sent response")
		}
	}
}

func (b *Backend) handleDlDataResMessage(ctx context.Context, msg messages.DlDataRes, eui common.EUI64) messages.Message {
	b.RLock()
	defer b.RUnlock()



	// txack, err := msg.IntoProto(eui)
	// if err != nil {
	// 	log.WithError(err).WithFields(log.Fields{
	// 		"gateway_id": eui,
	// 	}).Error("backend/bssci: error converting DlDataRes to protobuf message")
	// 	return nil
	// }

	// if v, ok := b.diidCache.Get(fmt.Sprintf("%d", txack.GetDownlinkId())); ok {
	// 	pl := v.(*gw.DownlinkFrame)
	// 	txack.DownlinkId = pl.DownlinkId

	// 	if conn, err := b.basestations.get(eui); err == nil {
	// 		conn.stats.CountDownlink(pl, &txack)
	// 	}
	// }

	// log.WithFields(log.Fields{
	// 	"gateway_id":  eui,
	// 	"downlink_id": txack.GetDownlinkId(),
	// }).Info("backend/bssci: DlDataRes message received")

	// if b.downlinkTxAckFunc != nil {
	// 	b.downlinkTxAckFunc(&txack)
	// }

	// handle with b.downlinkTxAckFunc

	// TODO do something with the data
	// clean up cache / queue
	defaultResponse := messages.NewDlDataResCmp(msg.GetOpId())
	return &defaultResponse

}

func (b *Backend) handleConMessage(ctx context.Context, conn *connection, eui common.EUI64, msg messages.Con) messages.Message {
	logger := zerolog.Ctx(ctx)

	resume, snScUuid := conn.ResumeConnection(msg.SnBsUuid.ToUuid(), msg.SnScOpId)

	conRsp := messages.NewConRsp(msg.GetOpId(), msg.Version, snScUuid)

	// check if session uuid is identical to current session
	if resume {
		logger.Info().Msg("Resuming current session")
		// set resume flag
		conRsp.SnResume = true

	} else {
		logger.Warn().Msg("Failed to resume current session")
	}
	return &conRsp
}

func (b *Backend) handleAttMessage(ctx context.Context, eui common.EUI64, msg messages.Att) messages.Message {
	// TODO response builder
	// TODO do something with the data
	// send to broker and get session key
	// propagate to all basestations

	return nil
}

func (b *Backend) handleDetMessage(ctx context.Context, eui common.EUI64, msg messages.Det) messages.Message {
	// TODO do something with the data
	// send to broker
	// propagate to all basestations

	return msg.Respond()
}

func (b *Backend) handleUlDataMessage(ctx context.Context, eui common.EUI64, msg messages.UlData) messages.Message {

	// TODO do something with the data
	// send to broker
	// handle downlink if rx is open
	defaultResponse := messages.NewUlDataRsp(msg.GetOpId())
	return &defaultResponse
}

func (b *Backend) handleStatusRspMessage(ctx context.Context, eui common.EUI64, msg messages.StatusRsp) messages.Message {

	// TODO do something with the data
	// send to broker
	defaultResponse := messages.NewStatusCmp(msg.GetOpId())
	return &defaultResponse
}

func (b *Backend) handleDlRxStatMessage(ctx context.Context, eui common.EUI64, msg messages.DlRxStat) messages.Message {

	// TODO do something with the data
	// send to broker
	defaultResponse := messages.NewDlRxStatRsp(msg.GetOpId())
	return &defaultResponse
}
