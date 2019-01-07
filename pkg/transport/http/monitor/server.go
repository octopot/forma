package monitor

import (
	"expvar"
	"net"
	"net/http"

	"github.com/kamilsk/form-api/pkg/config"
	"github.com/kamilsk/form-api/pkg/transport"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// New TODO issue#173
func New(_ config.MonitoringConfig) transport.Server {
	return &server{}
}

type server struct{}

// Serve TODO issue#173
func (*server) Serve(listener net.Listener) error {
	defer func() { _ = listener.Close() }()
	mux := &http.ServeMux{}
	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/vars", expvar.Handler())
	return http.Serve(listener, mux)
}
