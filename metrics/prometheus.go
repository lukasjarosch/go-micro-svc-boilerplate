package metrics

import (
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics interface {
	Register() error
	Serve()
}

type Prometheus struct {
	log *logrus.Logger
	HttpEndpoint string
	HttpPort int
}

func NewPrometheus(log *logrus.Logger, endpoint string, port int) *Prometheus {
	// TODO: ensure slash prefix
	if endpoint == "" {
		endpoint = "/metrics"
	}
	if port == 0 {
		port = 8080
	}
	return &Prometheus{log: log, HttpEndpoint:endpoint, HttpPort:port}
}

func (m *Prometheus) Register() error {

	publishedEventCount := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "published_event_count",
			Help: "The total amount of published events",
		},
		[]string{"topic"},
	)
	if err := prometheus.Register(publishedEventCount); err != nil {
		m.log.WithError(err).Warn("failed to register 'published_events_count'")
	}
	m.log.Debugf("registered metric: published_event_count")


	receivedEventCount := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "received_event_count",
			Help: "The total amount of published events",
		},
		[]string{"topic"},
	)
	if err := prometheus.Register(receivedEventCount); err != nil {
		m.log.WithError(err).Warn("failed to register 'received_event_count'")
	}
	m.log.Debugf("registered metric: received_event_count")

	return nil
}

func (m *Prometheus) Serve() {
	http.Handle(m.HttpEndpoint, promhttp.Handler())
	m.log.Infof("serving metrics at 0.0.0.0:%d%s", m.HttpPort, m.HttpEndpoint)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", m.HttpPort), nil); err != nil {
		m.log.WithError(err).Error("failed to serve metrics")
	}
}