package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/rjribeiro/hex/cmd/hex/domain"
	"time"
)

const MinimumAmount = 500

type AccountRequest struct {
	AccountId   string  `json:"account_id"`
	CustomerId  string  `json:"customer_id"`
	OpeningDate string  `json:"opening_date" validate:"required,validateDate"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount" validate:"required,validateAmount"`
	Status      string  `json:"status"`
}

func (a *AccountRequest) Validate() error {
	validate := validator.New()
	_ = validate.RegisterValidation("validateDate", validateDate)
	_ = validate.RegisterValidation("validateAmount", validateAmount)
	return validate.Struct(a)
}

func validateAmount(fl validator.FieldLevel) bool {
	return fl.Field().Float() > MinimumAmount
}

func ToDtoAccountRequest(account domain.Account) AccountRequest {
	return AccountRequest{
		AccountId:   account.AccountId,
		CustomerId:  account.CustomerId,
		OpeningDate: account.OpeningDate.String(),
		AccountType: account.AccountType,
		Amount:      account.Amount,
		Status:      account.MapStatusFromBool(),
	}
}

type AccountResponse struct {
	AccountId   string  `json:"account_id"`
	CustomerId  string  `json:"-"`
	OpeningDate string  `json:"-"`
	AccountType string  `json:"-"`
	Amount      float64 `json:"-"`
	Status      string  `json:"-"`
}

func ToDtoAccountResponse(account domain.Account) AccountResponse {
	return AccountResponse{
		AccountId:   account.AccountId,
		CustomerId:  account.CustomerId,
		OpeningDate: account.OpeningDate.String(),
		AccountType: account.AccountType,
		Amount:      account.Amount,
		Status:      account.MapStatusFromBool(),
	}
}

func FromDtoAccountRequest(account AccountRequest) domain.Account {
	openingDate, _ := time.Parse("2006-01-02 15:04:05", account.OpeningDate)
	return domain.Account{
		AccountId:   account.AccountId,
		CustomerId:  account.CustomerId,
		OpeningDate: openingDate,
		AccountType: account.AccountType,
		Amount:      account.Amount,
		Status:      domain.MapStatusToBool(account.Status),
	}
}
