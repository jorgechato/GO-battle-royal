package metrics

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//URLPatterns package routers
func URLPatterns(router *mux.Router) {
	router.Handle("", promhttp.Handler()).Name("Prometheus")
}
