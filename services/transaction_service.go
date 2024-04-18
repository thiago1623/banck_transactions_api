package services

import (
	"github.com/gin-gonic/gin"
	"github.com/thiago1623/banck_transactions_api/config"
	"github.com/thiago1623/banck_transactions_api/models"
	"github.com/thiago1623/banck_transactions_api/repositories"
	"net/http"
)

func SaveTransactions(c *gin.Context, transactions []models.Transaction) {
	db := config.DB
	transactionRepo := repositories.NewTransactionRepository(db)
	for _, transaction := range transactions {
		if err := transactionRepo.SaveTransaction(&transaction); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "save data error"})
			return
		}
	}
}
