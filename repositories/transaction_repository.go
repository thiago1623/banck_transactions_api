package repositories

import (
	"github.com/thiago1623/banck_transactions_api/models"
	"gorm.io/gorm"
)

// TransactionRepository proporciona métodos para interactuar con la tabla de transacciones en la base de datos.
type TransactionRepository struct {
	DB *gorm.DB
}

// NewTransactionRepository crea una nueva instancia de TransactionRepository.
func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

// SaveTransaction guarda una nueva transacción en la base de datos.
func (tr *TransactionRepository) SaveTransaction(transaction *models.Transaction) error {
	result := tr.DB.Create(transaction)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
