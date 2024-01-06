package domain

import (
	"github.com/rjribeiro/hex/cmd/hex/dto"
	"time"
)

type Customer struct {
	Id          string
	Name        string
	City        string
	ZipCode     string
	DateOfBirth time.Time
	Status      bool
}

func ToDtoCustomer(customer Customer) dto.CustomerResponse {
	return dto.CustomerResponse{
		Name:        customer.Name,
		City:        customer.City,
		ZipCode:     customer.ZipCode,
		DateOfBirth: customer.DateOfBirth.Format("2006-01-02"),
		Status:      customer.MapStatus(),
	}
}

func (c Customer) MapStatus() string {
	if c.Status {
		return "active"
	}
	return "inactive"
}
