package dto

import "github.com/rjribeiro/hex/cmd/hex/domain"

type CustomerResponse struct {
	Name        string `json:"name"`
	City        string `json:"city"`
	ZipCode     string `json:"zip_code"`
	DateOfBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}

func ToDtoCustomerResponse(customer domain.Customer) CustomerResponse {
	return CustomerResponse{
		Name:        customer.Name,
		City:        customer.City,
		ZipCode:     customer.ZipCode,
		DateOfBirth: customer.DateOfBirth.Format("2006-01-02"),
		Status:      customer.MapStatus(),
	}
}
