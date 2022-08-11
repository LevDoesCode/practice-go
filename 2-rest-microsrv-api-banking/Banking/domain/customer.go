package domain

import "Banking/errs"

// Customer is the domain object
type Customer struct {
	Id      string `json:"idee" xml:"idee"`
	Name    string
	City    string
	ZipCode string
	DOB     string
	Status  string
}

// CustomerRepository is the Secondary port that will connect to the server side
// All adapters need to conform to this interface
// as CustomerRepository (port) works as a protocol
type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
