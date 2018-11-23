package api

import (
	"net/http"

	"github.com/jorgechato/battle-royal/api/compute"
)

// Build the needed endpoints
func Build() (mux *http.ServeMux) {
	// Register handlers to routes.
	mux = http.NewServeMux()
	mux.Handle("/graphql", graphqlHandler())
	mux.Handle("/compute", compute.Handler())
	mux.Handle("/health", healthHandler())
	mux.Handle("/hello", helloHandler())

	return
}
