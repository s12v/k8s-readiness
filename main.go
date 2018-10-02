package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var healthy = true
var ready = true
var timeoutStart = time.Unix(0, 0)

func main() {

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if timeout() {
			time.Sleep(6 * time.Second)
		}

		if ! healthy {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		if timeout() {
			time.Sleep(6 * time.Second)
		}

		if ! ready {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/set/not/ready", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			ready = false
		}
	})

	http.HandleFunc("/set/ready", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			ready = true
		}
	})

	http.HandleFunc("/set/not/healthy", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			healthy = false
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})

	http.HandleFunc("/set/healthy", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			healthy = true
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})

	http.HandleFunc("/set/timeout", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			timeoutStart = time.Now()
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})

	http.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, os.Getenv("HOSTNAME"))
		fmt.Fprint(w, "hi")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	http.ListenAndServe(":8080", nil)
}

func timeout() bool {
	return time.Since(timeoutStart) < 1 * time.Minute
}
