package main

import (
	"github.com/rjribeiro/hex/cmd/hex/app/rest"
	"github.com/rjribeiro/hex/cmd/hex/logger"
	"github.com/rjribeiro/hex/cmd/hex/repository"
)

func main() {
	logger.Info("Starting application")

	postgresConn := repository.GetPostgresConnection()
	defer repository.ClosePostgresConnection(postgresConn)

	customerRepo := repository.NewCustomerRepositoryPostgres(postgresConn)
	accountRepo := repository.NewAccountRepositoryPostgres(postgresConn)
	//accountRepo := repository.NewAccountRepositoryStub()

	rest.Start(customerRepo, accountRepo)
}
