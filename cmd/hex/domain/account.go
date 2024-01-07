package domain

import (
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

func (c Account) MapStatusFromBool() string {
	if c.Status {
		return "active"
	}
	return "inactive"
}

func MapStatusToBool(status string) bool {
	if status == "active" {
		return true
	}
	return false
}
