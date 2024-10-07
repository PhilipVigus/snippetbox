package main

import (
	"log"
	"net/http"
)

// writes a byte slice containing the string "Hello from SnippetBox" as the response body
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from SnippetBox"))
}

func main() {
	mux := http.NewServeMux()
	// Register the home function as the handler for the "/" URL pattern
	mux.HandleFunc("/", home)
	log.Print("Starting server on :4000")
	// Start the web server using the default ServeMux as the router
	// we aren't specifying a host, so the server will listen on all available interfaces on the machine
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
