package app

import (
	"Banking/domain"
	"Banking/service"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func sanityCheck() {
	var missing []string
	if os.Getenv("SERVER_ADDRESS") == "" {
		missing = append(missing, "SERVER_ADDRESS")
	}
	if os.Getenv("SERVER_PORT") == "" {
		missing = append(missing, "SERVER_PORT")
	}
	if os.Getenv("DB_USER") == "" {
		missing = append(missing, "DB_USER")
	}
	if os.Getenv("DB_PASS") == "" {
		missing = append(missing, "DB_PASS")
	}
	if os.Getenv("DB_ADDR") == "" {
		missing = append(missing, "DB_ADDR")
	}
	if os.Getenv("DB_PORT") == "" {
		missing = append(missing, "DB_PORT")
	}
	if os.Getenv("DB_NAME") == "" {
		missing = append(missing, "DB_NAME")
	}
	var err string
	for _, v := range missing {
		err += fmt.Sprintf("Environment variable %s not defined\n", v)
	}
	if len(missing) > 0 {
		log.Fatal(err)
	}
	//log.Fatal("Environment variable not defined")
}

func StartServer() {
	// We register the pattern and the handler function
	// The default multiplexer is used by HandleFunc
	//http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello, world!")
	//}) // replace by a using a function name as parameter
	// we create a new multiplexer that can be customized instead of the default one
	//mux := http.NewServeMux()
	sanityCheck()
	myRouter := mux.NewRouter() // mux from gorilla-mux

	//  define routes
	// .Methods(http.MethodGet) to make explicit it's methods matcher
	// otherwise it's an http request by default
	myRouter.HandleFunc("/greet", greet).Methods(http.MethodGet)
	// wiring
	//ch := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	myRouter.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	myRouter.HandleFunc("/customers/", ch.getAllCustomers).Methods(http.MethodGet)
	myRouter.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	//myRouter.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomerById).Methods(http.MethodGet)
	myRouter.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	// starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Println(port)
	log.Println(address)
	// log.Fatal(http.ListenAndServe("localhost:8080", myRouter)) // old baked in call
	// log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), myRouter)) // also works
	log.Fatal(http.ListenAndServe(address+":"+port, myRouter))
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
