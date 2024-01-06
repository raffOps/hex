package errs

import "fmt"

type AppError struct {
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

type CustomerNotFoundError struct {
	Id string
}

func (e CustomerNotFoundError) Error() string {
	return "Customer id " + e.Id + " not found"
}

type AccountNotFoundError struct {
	Id string
}

func (e AccountNotFoundError) Error() string {
	return "Account not found for id " + e.Id
}

type InvalidAccountConfigError struct {
	Err error
}

func (e InvalidAccountConfigError) Error() string {
	return "Invalid account config: " + e.Err.Error()
}

type RepositoryError struct {
	Err error
}

func (e RepositoryError) Error() string {
	return "Repository error: " + e.Err.Error()
}

type DuplicateAccountError struct {
	AccountId string
}

func (e DuplicateAccountError) Error() string {
	return "Duplicate account for id " + e.AccountId
}

type InvalidAccountError struct {
	Err error
}

func (e InvalidAccountError) Error() string {
	return e.Err.Error()
}

type InvalidAmountError struct {
	Amount float64
}

func (e InvalidAmountError) Error() string {
	return fmt.Sprintf("Amount must be greater than 0. Amount received: %f", e.Amount)
}
