package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// We need to create a router
	rt := mux.NewRouter().StrictSlash(true)

	// Add the "index" or root path
	rt.HandleFunc("/", Index)

	// Fire up the server
	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", rt))
}

// Index is the "index" handler
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World from %q", html.EscapeString(r.URL.Path))
}
