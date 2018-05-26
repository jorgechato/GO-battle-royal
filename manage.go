package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/jorgechato/battle-royale/battleroyale"
	"github.com/jorgechato/battle-royale/health"
	"github.com/jorgechato/battle-royale/metrics"
)

var (
	version = "1.0"
	tag     = ""
	author  = ""
	port    = "8080"
	host    = ""
	address = fmt.Sprintf("%v:%v", host, port)
	info    = map[string]string{
		"Build User":   "@" + author,
		"Version":      "v" + version,
		"Version desc": "v" + tag,
		"Server":       "http://" + address,
	}
	start = time.Now()
)

func urlPatterns(r *mux.Router) {
	battleroyale.URLPatterns(r.PathPrefix("/").Subrouter())
	health.URLPatterns(r.PathPrefix("/health").Subrouter())
	metrics.URLPatterns(r.PathPrefix("/metrics").Subrouter())
}

func main() {
	r := mux.NewRouter()
	r.Host(host)
	r.Headers("Content-Type", "application/(text|json)")

	urlPatterns(r)

	b, _ := json.MarshalIndent(info, "", "  ")
	fmt.Println(string(b))

	server := &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	metrics.RampUpTime.Set(
		float64(time.Since(start).Nanoseconds()),
	)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 2)
	defer cancel()
	server.Shutdown(ctx)
	os.Exit(0)
}
