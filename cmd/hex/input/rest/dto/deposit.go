package dto

import "github.com/rjribeiro/hex/cmd/hex/domain"

type DepositResponse struct {
	AccountId   string  `json:"account_id"`
	CustomerId  string  `json:"-"`
	OpeningDate string  `json:"-"`
	AccountType string  `json:"-"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"-"`
}

func ToDepositResponse(account domain.Account) DepositResponse {
	return DepositResponse{
		AccountId:   account.AccountId,
		CustomerId:  account.CustomerId,
		OpeningDate: account.OpeningDate.Format("2006-01-02 15:04:05"),
		AccountType: account.AccountType,
		Amount:      account.Amount,
		Status:      account.MapStatusFromBool(),
	}
}
