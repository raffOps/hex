package rest

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rjribeiro/hex/cmd/hex/app/rest/handlersApp"
	"github.com/rjribeiro/hex/cmd/hex/domain"
	"github.com/rjribeiro/hex/cmd/hex/logger"
	"github.com/rjribeiro/hex/cmd/hex/service"
	"net/http"
	"os"
)

func sanityCheck() {
	if _, ok := os.LookupEnv("APP_HOST"); !ok {
		panic("APP_HOST environment variable is not set")
	}

	if _, ok := os.LookupEnv("APP_PORT"); !ok {
		panic("APP_PORT environment variable is not set")
	}
}

// Start starts the REST API application. It need the folowing environment variables to be set:
// APP_HOST and APP_PORT
func Start(customerRepo domain.CustomerRepository, accountRepo domain.AccountRepository) {
	sanityCheck()
	router := mux.NewRouter()
	customerHandler := handlersApp.NewCustomerHandlers(
		service.NewDefaultCustomerService(
			customerRepo,
		),
	)
	accountHandler := handlersApp.NewAccountHandlers(
		service.NewDefaultAccountService(
			accountRepo,
			customerRepo,
		),
		service.NewDefaultCustomerService(
			customerRepo,
		),
	)

	router.HandleFunc("/customers", customerHandler.GetAllCustomers).Methods("GET")
	router.HandleFunc("/customers/{id:[0-9]*}", customerHandler.GetCustomerById).Methods("GET")

	router.HandleFunc("/accounts", accountHandler.Save).Methods("POST")
	router.HandleFunc("/accounts/{id:\\d*}/deposit/{amount:[+-]?(?:[0-9]*[.])?[0-9]+}", accountHandler.Deposit).Methods("POST")

	appHost := os.Getenv("APP_HOST")
	appPort := os.Getenv("APP_PORT")
	logger.Info(fmt.Sprintf("Starting API on %s:%s", appHost, appPort))
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", appHost, appPort), router)
	if err != nil {
		logger.Error(fmt.Sprintf("Error starting application: %s", err.Error()))
	}
}
