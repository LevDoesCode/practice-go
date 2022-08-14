package domain

import "Banking/errs"

// Customer is the domain object
type Customer struct {
	//Id      string `json:"idee" xml:"idee"`
	Id      string `db:"customer_id"`
	Name    string `db:"name"`
	DOB     string `db:"dob"`
	City    string `db:"city"`
	ZipCode string `db:"zipcode"`
	Status  string `db:"status"`
}

// CustomerRepository is the Secondary port that will connect to the server side
// All adapters need to conform to this interface
// as CustomerRepository (port) works as a protocol
type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
