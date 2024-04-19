package utils

import (
	"encoding/csv"
	"github.com/thiago1623/banck_transactions_api/models"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

// ParseCSV reads the CSV file and returns a list of transactions.
func ParseCSV(filePath string) ([]models.Transaction, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	var transactions []models.Transaction
	firstRow := true
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error al leer el archivo CSV:", err)
			continue
		}
		if firstRow && record[0] == "date" {
			firstRow = false
			continue
		}
		transaction, err := parseRecordToTransaction(record)
		if err != nil {
			log.Println("Error parsing CSV file log:", err)
			continue
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

// parseRecordToTransaction converts a CSV record to a transaction.
func parseRecordToTransaction(record []string) (models.Transaction, error) {
	var transaction models.Transaction
	if record[1] == "date" {
		return transaction, nil
	}
	date, err := time.Parse("1/2/2006", record[1]) // Formato: mes/día/año
	if err != nil {
		return transaction, err
	}
	amount, err := strconv.ParseFloat(record[2], 64)
	if err != nil {
		return transaction, err
	}

	transaction = models.Transaction{
		Date:        date,
		Transaction: amount,
	}

	return transaction, nil
}

// CreateSummaryCSV creates a CSV file with the data necessary to generate the summary.
func CreateSummaryCSV(transactions []models.Transaction) (string, error) {
	// Crear el archivo CSV
	filePath := "summary.csv"
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Crear un escritor CSV
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Escribir la cabecera del CSV
	header := []string{"Total balance", "Month", "Number of transactions", "Average debit", "Average credit"}
	if err := writer.Write(header); err != nil {
		return "", err
	}

	// Calcular el total balance, número de transacciones, promedio de débito y promedio de crédito
	totalBalance := 0.0
	transactionsByMonth := make(map[string][]models.Transaction)
	for _, transaction := range transactions {
		totalBalance += transaction.Transaction
		month := transaction.Date.Format("January")
		transactionsByMonth[month] = append(transactionsByMonth[month], transaction)
	}

	// Escribir los datos en el CSV
	isFirstRow := true
	for month, transactions := range transactionsByMonth {
		numTransactions := len(transactions)
		averageDebit, averageCredit := calculateAverages(transactions)
		row := []string{}
		if isFirstRow {
			row = append(row, strconv.FormatFloat(totalBalance, 'f', 2, 64))
			isFirstRow = false
		} else {
			row = append(row, "")
		}
		row = append(row, month, strconv.Itoa(numTransactions), strconv.FormatFloat(averageDebit, 'f', 2, 64), strconv.FormatFloat(averageCredit, 'f', 2, 64))
		if err := writer.Write(row); err != nil {
			return "", err
		}
	}

	return filePath, nil
}

// calculateAverages calculates the average debit and credit for a list of transactions.
func calculateAverages(transactions []models.Transaction) (float64, float64) {
	totalDebit := 0.0
	totalCredit := 0.0
	for _, transaction := range transactions {
		if transaction.Transaction < 0 {
			totalDebit += transaction.Transaction
		} else {
			totalCredit += transaction.Transaction
		}
	}
	numTransactions := len(transactions)
	averageDebit := totalDebit / float64(numTransactions)
	averageCredit := totalCredit / float64(numTransactions)
	return averageDebit, averageCredit
}
