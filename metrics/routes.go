package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Build the needed endpoints
func Build() (mux *http.ServeMux) {
	// Register handlers to routes.
	mux = http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	return
}
