package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

// By the default, the json response will use the struct definition for the keys
// If we want to change them we can add the tags in backticks ``
type Customer struct {
	Name    string `json:"full_name" xml:"full_name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code"`
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
