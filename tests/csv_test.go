package tests

import (
	"github.com/thiago1623/banck_transactions_api/models"
	"github.com/thiago1623/banck_transactions_api/utils"
	"testing"
	"time"
)

// TestParseCSV test for csv parser
func TestParseCSV(t *testing.T) {
	filePath := "./file_csv_test.csv"
	transactions, err := utils.ParseCSV(filePath)
	if err != nil {
		t.Errorf("Error al analizar el archivo CSV: %v", err)
	}
	expectedNumTransactions := 14
	if len(transactions) != expectedNumTransactions {
		t.Errorf("Wrong number of transactions: expected %d but got %d", expectedNumTransactions, len(transactions))
	}
	expectedFirstDate := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC).Truncate(24 * time.Hour)
	expectedFirstTransaction := models.Transaction{Date: expectedFirstDate, Transaction: 50.00}
	if !transactions[1].Date.Equal(expectedFirstTransaction.Date) || transactions[1].Transaction != expectedFirstTransaction.Transaction {
		t.Errorf("The transaction does not match what was expected")
	}
}
