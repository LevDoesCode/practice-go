package app

import (
	"log"
	"net/http"
)

func StartServer() {
	// We register the pattern and the handler function
	// The default multiplexer is used by HandleFunc
	//http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello, world!")
	//}) // replace by a using a function name as parameter
	// we create a new multiplexer that can be customized instead of the default one
	mux := http.NewServeMux()

	//  define routes
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)
	// starting server
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
	// ListenAndServe returns an error we can check if there's an error starting the server
	// http.ListenAndServe("localhost:8080", nil)
	// we pass a nil handler because we're not creating our own
}
