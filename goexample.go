package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	log.Print("Starting go docker test file, listen on 8081")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q, just a basic test.", html.EscapeString(r.URL.Path))
	})
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi there")
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
