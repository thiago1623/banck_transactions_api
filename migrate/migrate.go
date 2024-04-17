package main

import (
	"github.com/thiago1623/banck_transactions_api/config"
	"github.com/thiago1623/banck_transactions_api/models"
)

func init() {
	config.ConnectDB()
}

func main() {
	config.DB.AutoMigrate(&models.Transaction{})
}
