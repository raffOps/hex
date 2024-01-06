package repository

import (
	"github.com/rjribeiro/hex/cmd/hex/domain"
)

type AccountRepositoryStub struct {
	accounts []domain.Account
}

func NewAccountRepositoryStub() AccountRepositoryStub {
	return AccountRepositoryStub{}
}

func (s AccountRepositoryStub) Save(account domain.Account) (*domain.Account, error) {
	s.accounts = append(s.accounts, account)
	return &account, nil
}

func (s AccountRepositoryStub) FindById(id string) (*domain.Account, error) {
	for _, account := range s.accounts {
		if account.AccountId == id {
			return &account, nil
		}
	}
	return nil, nil
}

func (s AccountRepositoryStub) UpdateAmount(accountId string, amount float64) (*domain.Account, error) {
	for _, account := range s.accounts {
		if account.AccountId == accountId {
			account.Amount = amount
			return &account, nil
		}
	}
	return nil, nil
}

func (s AccountRepositoryStub) BeginTransaction() interface{} {
	return nil
}

func (s AccountRepositoryStub) RollbackTransaction() interface{} {
	return nil
}

func (s AccountRepositoryStub) CommitTransaction() interface{} {
	return nil
}
