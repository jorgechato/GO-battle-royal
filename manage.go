package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jorgechato/battle-royale/battleroyale"
	"github.com/jorgechato/battle-royale/health"
)

var (
	version = "1.0"
	tag     = ""
	author  = ""
	port    = "8080"
	host    = ""
)

func main() {
	r := mux.NewRouter()
	r.Host(host)
	r.Headers("X-Requested-With", "XMLHttpRequest")
	r.Headers("Content-Type", "application/(text|json)")

	battleroyale.URLPatterns(r.PathPrefix("/").Subrouter())
	health.URLPatterns(r.PathPrefix("/health").Subrouter())

	address := fmt.Sprintf("%v:%v", host, port)
	info := map[string]string{
		"Build User":   "@" + author,
		"Version":      "v" + version,
		"Version desc": "v" + tag,
		"Server":       "http://" + address,
	}
	b, _ := json.MarshalIndent(info, "", "  ")
	fmt.Println(string(b))

	log.Fatal(
		http.ListenAndServe(
			address,
			r,
		),
	)
}
