package rest

import (
	"encoding/json"
	"github.com/rjribeiro/hex/cmd/hex/service"
	"net/http"
)

type CustomerResponse struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func NewCustomerHandlers(service service.CustomerService) CustomerHandlers {
	return CustomerHandlers{service: service}
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()

	var customerResponse []CustomerResponse

	for _, customer := range customers {
		customerResponse = append(customerResponse, CustomerResponse{
			Name:    customer.Name,
			City:    customer.City,
			ZipCode: customer.ZipCode,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(customerResponse)
}
