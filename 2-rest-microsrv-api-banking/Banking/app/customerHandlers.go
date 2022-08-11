package app

import (
	"Banking/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// By the default, the json response will use the struct definition for the keys
// If we want to change them we can add the tags in backticks ``
// We don't need this struct anymore as it's implemented in customer.go now
//type Customer struct {
//	Name    string `json:"full_name" xml:"full_name"`
//	City    string `json:"city" xml:"city"`
//	ZipCode string `json:"zip_code" xml:"zip_code"`
//}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	//customers := []Customer{
	//	{Name: "Lev", City: "Lima", ZipCode: "15103"},
	//	{Name: "Roger", City: "Springfield", ZipCode: "10001"},
	//}
	customers, err := ch.service.GetAllCustomers()
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else if r.Header.Get("Content-Type") == "application/json" {
		writeResponse(w, http.StatusOK, customers)
		//w.Header().Add("Content-Type", "application/json")
		//json.NewEncoder(w).Encode(customers)
	}
	// Set the correct response header for the response writer from plain text to json or xml
	//fmt.Printf("%s\n", r.Header.Get("Content-Type"))

}

func (ch *CustomerHandler) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	// This is the appropriate order to write the responses
	// w.Header().Add("Content-Type", "application/json")
	// w.WriteHeader()
	// json.NewEncoder(w).Encode(err.AsMessage())
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
