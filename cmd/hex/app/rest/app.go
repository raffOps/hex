package rest

import (
	"github.com/gorilla/mux"
	"github.com/rjribeiro/hex/cmd/hex/repository"
	"github.com/rjribeiro/hex/cmd/hex/service"
	"log"
	"net/http"
)

func Start(customerRepo repository.CustomerRepository) {
	router := mux.NewRouter()
	handlers := NewCustomerHandlers(
		service.NewDefaultCustomerService(
			customerRepo,
		),
	)
	//handlers := NewCustomerHandlers(
	//	service.NewDefaultCustomerService(
	//		repository.NewCustomerRepositoryStub(),
	//		),
	//	)

	router.HandleFunc("/customers", handlers.getAllCustomers).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
