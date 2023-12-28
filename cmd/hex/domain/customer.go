package domain

import "time"

type Customer struct {
	Id          string
	Name        string
	City        string
	ZipCode     string
	DateOfBirth time.Time
	Status      string
}
