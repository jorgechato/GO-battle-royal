package help

import (
	"github.com/gorilla/mux"
)

//URLPatterns package routers
func URLPatterns(router *mux.Router) {
	router.StrictSlash(true) // Default will redirect the endpoint to the right tail

	router.HandleFunc("/", getHelpCtrl).Methods("GET").Name("HelpHome")
}
