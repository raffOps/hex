package service

import (
	"github.com/rjribeiro/hex/cmd/hex/domain"
	"github.com/rjribeiro/hex/cmd/hex/dto"
	"github.com/rjribeiro/hex/cmd/hex/errs"
	"github.com/rjribeiro/hex/cmd/hex/logger"
)

type AccountService interface {
	NewAccount(account dto.Account) (*dto.AccountResponse, error)
	Deposit(accountId string, amountToDeposit float64) (*dto.DepositResponse, error)
}

type DefaultAccountService struct {
	accountRepo  domain.AccountRepository
	customerRepo domain.CustomerRepository
}

func NewDefaultAccountService(accountRepo domain.AccountRepository, customerRepo domain.CustomerRepository) DefaultAccountService {
	return DefaultAccountService{accountRepo: accountRepo, customerRepo: customerRepo}
}

func (s DefaultAccountService) NewAccount(account dto.Account) (*dto.AccountResponse, error) {
	err := account.Validate()
	if err != nil {
		return nil, errs.InvalidAccountError{Err: err}
	}

	customer, err := s.customerRepo.FindById(account.CustomerId)
	if customer == nil {
		return nil, errs.CustomerNotFoundError{Id: account.CustomerId}
	}

	accountToSave := domain.FromDtoAccount(account)
	domainAccount, err := s.accountRepo.Save(accountToSave)
	if err != nil {
		return nil, err
	}

	accountResponse := domain.ToDtoAccountResponse(*domainAccount)
	return &accountResponse, nil
}

func (s DefaultAccountService) Deposit(accountId string, amountToDeposit float64) (*dto.DepositResponse, error) {
	if amountToDeposit <= 0 {
		return nil, errs.InvalidAmountError{Amount: amountToDeposit}
	}
	domainAccount, err := s.accountRepo.FindById(accountId)
	if err != nil {
		logger.Debug("account service: Error while searching account")
		return nil, err
	}
	amountAfterDeposit := domainAccount.Amount + amountToDeposit
	domainAccount, err = s.accountRepo.UpdateAmount(accountId, amountAfterDeposit)
	if err != nil {
		return nil, err
	}

	accountResponse := domain.ToDtoAccountResponse(*domainAccount)
	return (*dto.DepositResponse)(&accountResponse), nil
}
