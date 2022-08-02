package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

// By the default, the json response will use the struct definition for the keys
// If we want to change them we can add the tags in backticks ``
type Customer struct {
	Name    string `json:"full_name" xml:"full_name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code"`
}

func main() {
	// We register the pattern and the handler function
	// The default multiplexer is used by HandleFunc
	//http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello, world!")
	//}) // replace by a using a function name as parameter
	//  define routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)
	// starting server
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
	// ListenAndServe returns an error we can check if there's an error starting the server
	// http.ListenAndServe("localhost:8080", nil)
	// we pass a nil handler because we're not creating our own
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Lev", City: "Lima", ZipCode: "15103"},
		{Name: "Roger", City: "Springfield", ZipCode: "10001"},
	}
	// Set the correct response header for the response writer from plain text to json or xml
	//fmt.Printf("%s\n", r.Header.Get("Content-Type"))
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
