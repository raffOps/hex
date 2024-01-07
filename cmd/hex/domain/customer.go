package domain

import (
	"time"
)

type Customer struct {
	Id          string
	Name        string
	City        string
	ZipCode     string
	DateOfBirth time.Time
	Status      bool
}

func (c Customer) MapStatus() string {
	if c.Status {
		return "active"
	}
	return "inactive"
}
