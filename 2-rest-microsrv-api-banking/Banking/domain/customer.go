package domain

// Customer is the domain object
type Customer struct {
	Id      string
	Name    string
	City    string
	ZipCode string
	DOB     string
	Status  string
}

// CustomerRepository is the Secondary port that will connect to the server side
type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
