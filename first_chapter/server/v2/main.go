// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)

	timeoutTime := 3

	server := &http.Server{
		Addr:              "localhost:8000",
		ReadHeaderTimeout: time.Duration(timeoutTime) * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()

	_, err := fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	if err != nil {
		return
	}
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, _ *http.Request) {
	mu.Lock()

	_, err := fmt.Fprintf(w, "Count %d\n", count)
	if err != nil {
		return
	}

	mu.Unlock()
}
