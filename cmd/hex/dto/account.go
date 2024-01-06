package dto

import (
	"github.com/go-playground/validator/v10"
)

const MinimumAmount = 500

type Account struct {
	AccountId   string  `json:"account_id"`
	CustomerId  string  `json:"customer_id"`
	OpeningDate string  `json:"opening_date" validate:"required,validateDate"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount" validate:"required,validateAmount"`
	Status      string  `json:"status"`
}

type AccountResponse struct {
	AccountId   string  `json:"account_id"`
	CustomerId  string  `json:"-"`
	OpeningDate string  `json:"-"`
	AccountType string  `json:"-"`
	Amount      float64 `json:"-"`
	Status      string  `json:"-"`
}

type DepositResponse struct {
	AccountId   string  `json:"account_id"`
	CustomerId  string  `json:"-"`
	OpeningDate string  `json:"-"`
	AccountType string  `json:"-"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"-"`
}

func (a *Account) Validate() error {
	validate := validator.New()
	_ = validate.RegisterValidation("validateDate", validateDate)
	_ = validate.RegisterValidation("validateAmount", validateAmount)
	return validate.Struct(a)
}

func validateAmount(fl validator.FieldLevel) bool {
	return fl.Field().Float() > MinimumAmount
}
