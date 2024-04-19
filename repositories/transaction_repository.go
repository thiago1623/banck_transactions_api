package repositories

import (
	"github.com/thiago1623/banck_transactions_api/models"
	"gorm.io/gorm"
)

// TransactionRepository provides methods to interact with the transaction table in the database.
type TransactionRepository struct {
	DB *gorm.DB
}

// NewTransactionRepository creates a new TransactionRepository instance.
func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

// SaveTransaction saves a new transaction to the database.
func (tr *TransactionRepository) SaveTransaction(transaction *models.Transaction) error {
	result := tr.DB.Create(transaction)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
