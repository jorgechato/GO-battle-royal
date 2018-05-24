package battleroyale

import (
	"github.com/gorilla/mux"
)

//URLPatterns package routers
func URLPatterns(router *mux.Router) {
	router.HandleFunc("/", homeController).Methods("GET").Name("Home")
}
