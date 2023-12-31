package controllers

import (
	"github.com/gorilla/mux"
	"github.com/rjribeiro/hex/cmd/hex/domain"
	"net/http"
)

type CustomerHandlers struct {
	service domain.CustomerService
}

func NewCustomerHandlers(service domain.CustomerService) CustomerHandlers {
	return CustomerHandlers{service: service}
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := ch.service.GetAllCustomer()

	if err != nil {
		HandleError(w, err)
		return
	}

	writeResponse(w, http.StatusOK, customers)
}

func (ch *CustomerHandlers) GetCustomerById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	customer, err := ch.service.GetCustomerById(id)
	if err != nil {
		HandleError(w, err)
		return
	}
	writeResponse(w, http.StatusOK, customer)
}
