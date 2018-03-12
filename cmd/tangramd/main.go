package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

type apiServer struct {
	server *http.Server
}

// Application exit status codes
const (
	successExitStatus            = 0
	errorStartingHTTPServer      = 1
	errorLoadingConfig           = 2
	errorStopintServerStatusCode = 3
)

// Configuration values
const (
	defaultAddress          = ":2018"
	defaultShutdownTimeout  = 2 * time.Millisecond
	defaultHTTPReadTimeout  = 500 * time.Millisecond
	defaultHTTPWriteTimeout = 500 * time.Second
)

// Application build details
var (
	version   = "development"
	build     = "undefined"
	buildDate = "undefined"
)

func main() {
	log.Println("The Tangram Composer")
	log.Printf("\tversion:      %s\n", version)
	log.Printf("\tbuild:        %s\n", build)
	log.Printf("\tbuild date:   %s\n", buildDate)
	log.Printf("\tstartup date: %s\n", time.Now().Format(time.RFC3339))

	apiserver := apiServer{}
	apiserver.startHTTPServer()
	apiserver.waitAndShutdown(defaultShutdownTimeout)
}

func (s *apiServer) startHTTPServer() {
	// configure HTTP server, register application status entrypoints and routes
	r := mux.NewRouter()
	s.server = &http.Server{
		Handler:      r,
		Addr:         defaultAddress,
		ReadTimeout:  defaultHTTPReadTimeout,
		WriteTimeout: defaultHTTPWriteTimeout,
	}

	r.HandleFunc("/healthy", healthyHandler).Methods("GET")
	r.HandleFunc("/", handler).Methods("GET")

	go func() {
		log.Printf("Listening on %s\n", defaultAddress)
		if err := s.server.ListenAndServe(); err != nil {
			log.Printf("Cannot start HTTP server. Error: %s", err)
			os.Exit(errorStartingHTTPServer)
		}
	}()
}

func (s *apiServer) waitAndShutdown(timeout time.Duration) {
	// deal with Ctrl+C (SIGTERM) and graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	log.Printf("Shutting down, with a timeout of %s...\n", timeout)
	if err := s.server.Shutdown(ctx); err != nil {
		log.Printf("Error stopping http server. Error: %v\n", err)
		os.Exit(errorStopintServerStatusCode)
	}
	log.Println("The Tangram Composer stoped")
	os.Exit(successExitStatus)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Not implemented yet")
}

func healthyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}
