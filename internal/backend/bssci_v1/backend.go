package bssci_v1

import (
	"crypto/tls"
	"crypto/x509"


	"net"
	"os"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"

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
	return nil
}