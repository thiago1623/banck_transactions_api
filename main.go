package main

import (
	"github.com/thiago1623/banck_transactions_api/config"
	"github.com/thiago1623/banck_transactions_api/routes"
)

func main() {
	config.ConnectDB()
	r := routes.SetupRouter()
	r.Run(":8080")
}
