package handler

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type metricServer struct {
	host string
}

func NewMetricServer(host string) metricServer {
	return metricServer{
		host: host,
	}
}

func (m metricServer) Start() error {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		http.ListenAndServe(m.host, nil)
	}()

	return nil
}

func (m metricServer) Stop() {

}
