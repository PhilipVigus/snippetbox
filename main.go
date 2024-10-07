package main

import (
	"log"
	"net/http"
)

// writes a byte slice containing the string "Hello from SnippetBox" as the response body
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from SnippetBox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	mux := http.NewServeMux()
	// Register the home function as the handler for the "/" URL pattern
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on :4000")
	// Start the web server using the default ServeMux as the router
	// we aren't specifying a host, so the server will listen on all available interfaces on the machine
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
