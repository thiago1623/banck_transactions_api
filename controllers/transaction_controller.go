package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/thiago1623/banck_transactions_api/config"
	"github.com/thiago1623/banck_transactions_api/services"
	"github.com/thiago1623/banck_transactions_api/utils"
	"net/http"
)

// TransactionController contiene the methods for handling transaction-related requests.
type TransactionController struct{}

// NewTransactionController creates a new TransactionController instance.
func NewTransactionController() *TransactionController {
	return &TransactionController{}
}

// ProcessCSV processes the CSV file containing transactions and saves them to the database.
func (tc *TransactionController) ProcessCSV(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	filePath := "uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el archivo"})
		return
	}
	transactions, err := utils.ParseCSV(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar el archivo CSV"})
		return
	}
	services.SaveTransactions(c, transactions)
	filePath, err = utils.CreateSummaryCSV(transactions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar el archivo CSV con el resumen"})
		return
	}
	emailService := services.NewEmailService()
	cfg := config.LoadSettings()
	serverSection := cfg.Section("Server")
	err = emailService.SendEmailWithTemplate(serverSection.Key("RecipientEmail").String(), "Resumen de transacciones", "utils/emails/transactions_email.html", filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error sending email"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transactions processed successfully"})
}
