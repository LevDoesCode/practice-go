package app

import (
	"Banking/domain"
	"Banking/service"
	"fmt"
	"github.com/gorilla/mux"
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
	//mux := http.NewServeMux()
	myRouter := mux.NewRouter() // mux from gorilla-mux

	//  define routes
	// .Methods(http.MethodGet) to make explicit it's methods matcher
	// otherwise it's an http request by default
	myRouter.HandleFunc("/greet", greet).Methods(http.MethodGet)
	// wiring
	//ch := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	myRouter.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	myRouter.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	//myRouter.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomerById).Methods(http.MethodGet)
	myRouter.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	// starting server
	log.Fatal(http.ListenAndServe("localhost:8080", myRouter))
	// ListenAndServe returns an error we can check if there's an error starting the server
	// http.ListenAndServe("localhost:8080", nil)
	// we pass a nil handler because we're not creating our own
}

func getCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Post Request received!")
}
