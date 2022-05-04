package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

var (
	listenAddr   string
	healthy      int32
	helloMessage = "I'm a little teapot"
)

func main() {
	flag.StringVar(&listenAddr, "listen-addr", ":8000", "server listen address")
	flag.Parse()

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	logger.Println("Linkfire Go Server")

	message := os.Getenv("CUSTOM_MESSAGE")

	if len(message) != 0 {
		helloMessage = message
	}

	router := http.NewServeMux()
	router.HandleFunc("/hello", Middle(logger, hello))
	router.HandleFunc("/healthz", Middle(logger, healthz))

	server := &http.Server{
		Addr:         listenAddr,
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	logger.Println("Server is ready to handle requests at", listenAddr)
	atomic.StoreInt32(&healthy, 1)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
	}

	logger.Println("Server stopped")
}

func Middle(l *log.Logger, f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l.Println(r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
		f(w, r)
	}
}
