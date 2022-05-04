package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

func toggleHealth(w http.ResponseWriter, r *http.Request) {
	if atomic.LoadInt32(&healthy) == 1 {
		fmt.Println("Marking server as unhealthy")
		atomic.StoreInt32(&healthy, 0)
	} else {
		fmt.Println("Marking server as healthy")
		atomic.StoreInt32(&healthy, 1)
	}
	w.WriteHeader(http.StatusOK)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		toggleHealth(w, r)
	} else if r.Method == http.MethodGet {
		if atomic.LoadInt32(&healthy) == 1 {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTeapot)
	fmt.Fprintln(w, helloMessage)
}
