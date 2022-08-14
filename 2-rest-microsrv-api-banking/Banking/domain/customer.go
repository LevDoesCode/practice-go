package domain

import (
	"Banking/dto"
	"Banking/errs"
)

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

func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Customer) ToDTO() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:      c.Id,
		Name:    c.Name,
		DOB:     c.DOB,
		City:    c.City,
		ZipCode: c.ZipCode,
		Status:  c.statusAsText(),
	}
}

func AllToDTO(customers []Customer, status string) []dto.CustomerResponse {
	var allDTO []dto.CustomerResponse
	for _, customer := range customers {
		if status == "" {
			allDTO = append(allDTO, customer.ToDTO())
		} else if customer.Status == status {
			allDTO = append(allDTO, customer.ToDTO())
		}
	}
	return allDTO
}

// CustomerRepository is the Secondary port that will connect to the server side
// All adapters need to conform to this interface
// as CustomerRepository (port) works as a protocol
type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
