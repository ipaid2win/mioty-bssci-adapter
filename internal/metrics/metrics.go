package metrics

import (
	"mioty-bssci-adapter/internal/config"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
)

// Setup configures the metrics package.
func Setup(conf config.Config) error {
	if !conf.Metrics.Prometheus.EndpointEnabled {
		return nil
	}

	log.Info().Str("bind", conf.Metrics.Prometheus.Bind).Msg("starting prometheus metrics server")

	server := http.Server{
		Handler: promhttp.Handler(),
		Addr:    conf.Metrics.Prometheus.Bind,
	}

	go func() {
		err := server.ListenAndServe()
		log.Error().Stack().Err(err).Msg("prometheus metrics server error")
	}()

	return nil
}
