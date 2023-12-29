package main

import (
	"github.com/rjribeiro/hex/cmd/hex/app/rest"
	"github.com/rjribeiro/hex/cmd/hex/repository"
)
import "github.com/rjribeiro/hex/cmd/hex/database"

func main() {
	postgresConn := database.GetPostgresConnection()
	repo := repository.NewCustomerRepositoryPostgres(postgresConn)
	rest.Start(repo)
}
