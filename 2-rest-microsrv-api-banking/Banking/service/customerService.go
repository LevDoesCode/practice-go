package service

import (
	"Banking/domain"
	"Banking/dto"
	"Banking/errs"
)

type CustomerService interface {
	GetAllCustomers(string) ([]dto.CustomerResponse, *errs.AppError)
	//GetCustomer(string) (*domain.Customer, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	//var res []dto.CustomerResponse
	//return s.repo.FindAll(status)

	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}

	// Create the slice here
	var allDTO []dto.CustomerResponse
	for _, customer := range customers {
		if status == "" {
			allDTO = append(allDTO, customer.ToDTO())
		} else if customer.Status == status {
			allDTO = append(allDTO, customer.ToDTO())
		}
	}
	
	// OR have Customer handle the logic
	allDTOs := domain.AllToDTO(customers, status)

	return allDTOs, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	res := c.ToDTO()
	return &res, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
