package service

import (
	"github.com/rjribeiro/hex/cmd/hex/domain"
	"github.com/rjribeiro/hex/cmd/hex/dto"
	"github.com/rjribeiro/hex/cmd/hex/errs"
)

type CustomerService interface {
	GetAllCustomer() ([]dto.CustomerResponse, error)
	GetCustomerById(string) (*dto.CustomerResponse, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func NewDefaultCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repo}
}

func (s DefaultCustomerService) GetAllCustomer() ([]dto.CustomerResponse, error) {
	customer, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	var customerResponse []dto.CustomerResponse
	for _, c := range customer {
		customerResponse = append(customerResponse, domain.ToDtoCustomer(c))
	}
	return customerResponse, nil
}

func (s DefaultCustomerService) GetCustomerById(id string) (*dto.CustomerResponse, error) {
	c, err := s.repo.FindById(id)
	if c == nil && err == nil {
		return nil, errs.CustomerNotFoundError{Id: id}
	}
	customer := domain.ToDtoCustomer(*c)
	return &customer, nil
}
