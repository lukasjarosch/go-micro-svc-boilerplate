package metrics

import (
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"fmt"
	"github.com/sirupsen/logrus"
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
	return &Prometheus{log: log, HttpEndpoint:endpoint, HttpPort:port}
}

func (m *Prometheus) Register() error {
	return nil
}

func (m *Prometheus) Serve() {
	http.Handle(m.HttpEndpoint, promhttp.Handler())
	if err := http.ListenAndServe(fmt.Sprintf(":%d", m.HttpPort), nil); err != nil {
		m.log.WithError(err).Error("failed to serve metrics")
	}
	m.log.Info("asdf")
}