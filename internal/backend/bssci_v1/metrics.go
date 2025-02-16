package bssci_v1

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	ppc = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "backend_bssci_ping_count",
		Help: "The number of BSSCI Ping/Pong requests sent and received (from server/client).",
	}, []string{"src", "bs"})

	rec = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "backend_bssci_received_count",
		Help: "The number of BSSCI messages received by the backend (per msgtype).",
	}, []string{"msgtype", "bs"})

	sent = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "backend_bssci_sent_count",
		Help: "The number of BSSCI messages sent by the backend (per msgtype).",
	}, []string{"msgtype", "bs"})

	bsc = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "backend_bssci_basestation_connect_count",
		Help: "The number of basestation connections received by the backend.",
	}, []string{"bs"})

	bsd = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "backend_bssci_basestation_disconnect_count",
		Help: "The number of basestations that disconnected from the backend.",
	}, []string{"bs"})
)

func pingPongCounter(src string, bs string,) prometheus.Counter {
	return ppc.With(prometheus.Labels{"bs": bs, "src": src})
}

func messageReceiveCounter(bs string, msgtype string) prometheus.Counter {
	return rec.With(prometheus.Labels{"bs": bs, "msgtype": msgtype})
}

func messagSendCounter(bs string, msgtype string) prometheus.Counter {
	return sent.With(prometheus.Labels{"bs": bs, "msgtype": msgtype})
}

func connectCounter(bs string) prometheus.Counter {
	return bsc.With(prometheus.Labels{"bs": bs})
}

func disconnectCounter(bs string) prometheus.Counter {
	return bsd.With(prometheus.Labels{"bs": bs})
}
