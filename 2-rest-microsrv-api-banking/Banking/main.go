package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// We register the pattern and the handler function
	// The default multiplexer is used by HandleFunc
	//http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello, world!")
	//})
	//  define routes
	http.HandleFunc("/greet", greet)
	// starting server
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
	// ListenAndServe returns an error we can check if there's an error starting the server
	// http.ListenAndServe("localhost:8080", nil)
	// we pass a nil handler because we're not creating our own
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}
