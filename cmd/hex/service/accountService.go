package service

import (
	"github.com/rjribeiro/hex/cmd/hex/domain"
	"github.com/rjribeiro/hex/cmd/hex/errs"
	"github.com/rjribeiro/hex/cmd/hex/logger"
)

type AccountService interface {
	NewAccount(account domain.Account) (*domain.Account, error)
	Deposit(accountId string, amountToDeposit float64) (*domain.Account, error)
}

type DefaultAccountService struct {
	accountRepo  domain.AccountRepository
	customerRepo domain.CustomerRepository
}

func NewDefaultAccountService(accountRepo domain.AccountRepository, customerRepo domain.CustomerRepository) DefaultAccountService {
	return DefaultAccountService{accountRepo: accountRepo, customerRepo: customerRepo}
}

func (s DefaultAccountService) NewAccount(account domain.Account) (*domain.Account, error) {
	customer, err := s.customerRepo.FindById(account.CustomerId)
	if customer == nil {
		return nil, errs.CustomerNotFoundError{Id: account.CustomerId}
	}

	accountResponse, err := s.accountRepo.Save(account)
	if err != nil {
		return nil, err
	}

	return accountResponse, nil
}

func (s DefaultAccountService) Deposit(accountId string, amountToDeposit float64) (*domain.Account, error) {
	if amountToDeposit <= 0 {
		return nil, errs.InvalidAmountError{Amount: amountToDeposit}
	}
	account, err := s.accountRepo.FindById(accountId)
	if err != nil {
		logger.Debug("account service: Error while searching account")
		return nil, err
	}
	amountAfterDeposit := account.Amount + amountToDeposit
	account, err = s.accountRepo.UpdateAmount(accountId, amountAfterDeposit)
	if err != nil {
		return nil, err
	}

	return account, nil
}
