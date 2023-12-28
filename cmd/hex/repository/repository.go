package repository

import "github.com/rjribeiro/hex/cmd/hex/domain"

type CustomerRepository interface {
	FindAll() ([]domain.Customer, error)
}
