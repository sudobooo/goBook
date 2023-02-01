// Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler

	timeoutTime := 3

	server := &http.Server{
		Addr:              "localhost:8000",
		ReadHeaderTimeout: time.Duration(timeoutTime) * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	if err != nil {
		return
	}
}
