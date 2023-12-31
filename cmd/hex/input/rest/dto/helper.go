package dto

import (
	"github.com/go-playground/validator/v10"
	"time"
)

func validateDate(fl validator.FieldLevel) bool {
	_, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}
	return true
}
