// Server3 is an "echo" server that displays request parameters.
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)

	timeoutTime := 3

	server := &http.Server{
		Addr:              "localhost:8000",
		ReadHeaderTimeout: time.Duration(timeoutTime) * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	if _, errPrint := fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto); errPrint != nil {
		return
	}

	for k, v := range r.Header {
		if _, err := fmt.Fprintf(w, "Header[%q] = %q\n", k, v); err != nil {
			return
		}
	}

	if _, err := fmt.Fprintf(w, "Host = %q\n", r.Host); err != nil {
		return
	}

	if _, err := fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr); err != nil {
		return
	}

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		if _, err := fmt.Fprintf(w, "Form[%q] = %q\n", k, v); err != nil {
			return
		}
	}
}
