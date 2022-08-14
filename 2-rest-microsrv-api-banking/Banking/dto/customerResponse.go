package dto

type CustomerResponse struct {
	Id      string `json:"customer_id"`
	Name    string `json:"full_name"`
	DOB     string `json:"date_of_birth"`
	City    string `json:"city"`
	ZipCode string `json:"zipcode"`
	Status  string `json:"status"`
}
