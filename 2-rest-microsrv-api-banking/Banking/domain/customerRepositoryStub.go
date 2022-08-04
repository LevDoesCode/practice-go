package domain

// Adapter to the Secondary Port
// CustomerRepositoryStub is the stub adapter for the server side
type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Lev", City: "Lima", ZipCode: "15103", DOB: "08-08-1985", Status: "1"},
		{Id: "1001", Name: "Roger", City: "Springfield", ZipCode: "10001", DOB: "10-10-1985", Status: "1"},
	}
	return CustomerRepositoryStub{customers}
}
