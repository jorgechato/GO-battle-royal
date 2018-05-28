package health

import (
	"github.com/gorilla/mux"
)

//URLPatterns package routers
func URLPatterns(router *mux.Router) {
	router.HandleFunc("", getHealthCtrl).Methods("GET").Name("Health")
}
