package repository

import (
	"github.com/rjribeiro/hex/cmd/hex/errs"
	"time"
)
import "github.com/rjribeiro/hex/cmd/hex/domain"

type CustomerRepositoryStub struct {
	customers []domain.Customer
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []domain.Customer{
		{
			Id:          "1",
			Name:        "John Doe",
			City:        "New York",
			ZipCode:     "10001",
			DateOfBirth: time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC),
			Status:      true,
		},
		{
			Id:          "2",
			Name:        "Jane Doe",
			City:        "Los Angeles",
			ZipCode:     "90001",
			DateOfBirth: time.Date(1985, 1, 1, 0, 0, 0, 0, time.UTC),
			Status:      false,
		},
	}

	return CustomerRepositoryStub{customers: customers}
}

func (c CustomerRepositoryStub) FindAll() ([]domain.Customer, error) {
	return c.customers, nil
}

func (c CustomerRepositoryStub) FindById(id string) (*domain.Customer, error) {
	var domainCustomer domain.Customer
	for _, customer := range c.customers {
		if id == customer.Id {
			return &customer, nil
		}
	}
	return &domainCustomer, errs.CustomerNotFoundError{Id: id}
}
