package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/thiago1623/banck_transactions_api/config"
	"github.com/thiago1623/banck_transactions_api/models"
	"github.com/thiago1623/banck_transactions_api/repositories"
	"github.com/thiago1623/banck_transactions_api/utils"
	"log"
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

	// Parsear el archivo CSV
	transactions, err := utils.ParseCSV(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar el archivo CSV"})
		return
	}

	// Guardar las transacciones en la base de datos
	db := config.DB
	transactionRepo := repositories.NewTransactionRepository(db)
	for _, transaction := range transactions {
		if err := transactionRepo.SaveTransaction(&transaction); err != nil {
			log.Println("Error al guardar la transacción:", err)
			// Podrías manejar el error de alguna manera apropiada, como registrar, enviar una respuesta JSON, etc.
		}
	}

	// Si todo salió bien, enviar una respuesta de éxito
	c.JSON(http.StatusOK, gin.H{"message": "Transacciones procesadas exitosamente"})
}
