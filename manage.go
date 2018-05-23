package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jorgechato/battle-royale/battleroyale"
	"github.com/jorgechato/battle-royale/help"
)

var (
	port = "8080"
	host = "localhost"
)

func main() {
	r := mux.NewRouter()

	battleroyale.URLPatterns(r.PathPrefix("/").Subrouter())
	help.URLPatterns(r.PathPrefix("/help").Subrouter())

	r.Host(host)
	r.Headers("X-Requested-With", "XMLHttpRequest")
	r.Headers("Content-Type", "application/(text|json)")

	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf("%v:%v", host, port),
			r,
		),
	)
}
