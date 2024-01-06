package domain

import (
	"github.com/rjribeiro/hex/cmd/hex/dto"
	"time"
)

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate time.Time
	AccountType string
	Amount      float64
	Status      bool
}

func ToDtoAccount(account Account) dto.Account {
	return dto.Account{
		AccountId:   account.AccountId,
		CustomerId:  account.CustomerId,
		OpeningDate: account.OpeningDate.String(),
		AccountType: account.AccountType,
		Amount:      account.Amount,
		Status:      account.MapStatus(),
	}
}

func ToDtoAccountResponse(account Account) dto.AccountResponse {
	return dto.AccountResponse{
		AccountId:   account.AccountId,
		CustomerId:  account.CustomerId,
		OpeningDate: account.OpeningDate.String(),
		AccountType: account.AccountType,
		Amount:      account.Amount,
		Status:      account.MapStatus(),
	}
}

func FromDtoAccount(account dto.Account) Account {
	openingDate, _ := time.Parse("2006-01-02 15:04:05", account.OpeningDate)
	return Account{
		AccountId:   account.AccountId,
		CustomerId:  account.CustomerId,
		OpeningDate: openingDate,
		AccountType: account.AccountType,
		Amount:      account.Amount,
		Status:      MapStatusToBool(account.Status),
	}
}

func (c Account) MapStatus() string {
	if c.Status {
		return "active"
	}
	return "inactive"
}

func (c Account) ToDtoAccountResponse() interface{} {
	return dto.AccountResponse{
		AccountId:   c.AccountId,
		CustomerId:  c.CustomerId,
		OpeningDate: c.OpeningDate.String(),
		AccountType: c.AccountType,
		Amount:      c.Amount,
		Status:      c.MapStatus(),
	}
}

func MapStatusToBool(status string) bool {
	if status == "active" {
		return true
	}
	return false
}
