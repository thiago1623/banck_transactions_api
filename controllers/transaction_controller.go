package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/thiago1623/banck_transactions_api/config"
	"github.com/thiago1623/banck_transactions_api/models"
	"github.com/thiago1623/banck_transactions_api/services"
	"github.com/thiago1623/banck_transactions_api/utils"
	"net/http"
)

func GetTransactions(c *gin.Context) {
	var transactions []models.Transaction
	config.DB.Find(&transactions)

	c.JSON(200, transactions)
}

// TransactionController contiene los métodos para manejar las solicitudes relacionadas con transacciones.
type TransactionController struct{}

// NewTransactionController crea una nueva instancia de TransactionController.
func NewTransactionController() *TransactionController {
	return &TransactionController{}
}

// ProcessCSV procesa el archivo CSV que contiene transacciones y las guarda en la base de datos.
func (tc *TransactionController) ProcessCSV(c *gin.Context) {
	// Obtener el archivo CSV del formulario
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Guardar el archivo en el servidor
	filePath := "uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el archivo"})
		return
	}
	// Parse the file CSV
	transactions, err := utils.ParseCSV(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar el archivo CSV"})
		return
	}
	// save transactions inside database
	services.SaveTransactions(c, transactions)
	// generate the file CSV with the summary
	filePath, err = utils.CreateSummaryCSV(transactions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar el archivo CSV con el resumen"})
		return
	}
	// send email with the summary
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
