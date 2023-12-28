package main

import "github.com/rjribeiro/hex/cmd/hex/app/rest"
import "github.com/rjribeiro/hex/cmd/hex/database"

func main() {
	postgresConn := database.GetPostgresConnection()
	rest.Start(postgresConn)
}
