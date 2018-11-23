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

	"github.com/jorgechato/battle-royal/api"
	"github.com/jorgechato/battle-royal/metrics"
	. "github.com/jorgechato/battle-royal/utils"
)

var (
	idleTimeout       = time.Second * 60
	writeTimeout      = time.Second * 10
	readHeaderTimeout = time.Second * 1
	maxHeaderBytes    = http.DefaultMaxHeaderBytes
	address           = fmt.Sprintf("%v:%v", HOST, PORT)
	hiddenAddress     = fmt.Sprintf("%v:%v", HOST, HIDDENPORT)
	info              = map[string]string{
		"Build User":   "@" + AUTHOR,
		"Version":      "v" + VERSION,
		"Version desc": "v" + TAG,
		"Server":       "http://" + address,
	}
	start = time.Now()
)

func main() {
	b, _ := json.MarshalIndent(info, "", "  ")
	fmt.Println(string(b))

	server := &http.Server{
		Addr:              address,
		Handler:           api.Build(),
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
		MaxHeaderBytes:    maxHeaderBytes,
	}

	hiddenServer := &http.Server{
		Addr:              hiddenAddress,
		Handler:           metrics.Build(),
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
		MaxHeaderBytes:    maxHeaderBytes,
	}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	go func() {
		log.Fatal(hiddenServer.ListenAndServe())
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
	hiddenServer.Shutdown(ctx)
	os.Exit(0)
}
