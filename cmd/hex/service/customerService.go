package service

import (
	"github.com/rjribeiro/hex/cmd/hex/domain"
	"github.com/rjribeiro/hex/cmd/hex/errs"
)

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func NewDefaultCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repo}
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	customers, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (s DefaultCustomerService) GetCustomerById(id string) (*domain.Customer, error) {
	customer, err := s.repo.FindById(id)
	if customer == nil && err == nil {
		return nil, errs.CustomerNotFoundError{Id: id}
	}
	return customer, nil
}
