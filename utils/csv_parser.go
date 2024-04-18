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

// ParseCSV lee el archivo CSV y devuelve una lista de transacciones.
func ParseCSV(filePath string) ([]models.Transaction, error) {
	// Abrir el archivo CSV
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Crear un lector CSV
	reader := csv.NewReader(file)

	// Variables para almacenar las transacciones procesadas
	var transactions []models.Transaction

	// Leer el archivo CSV línea por línea
	firstRow := true
	for {
		// Leer una fila del archivo CSV
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error al leer el archivo CSV:", err)
			continue
		}

		// Ignorar la primera fila si contiene la cadena "date"
		if firstRow && record[0] == "date" {
			firstRow = false
			continue
		}

		// Convertir los campos del registro en una transacción
		transaction, err := parseRecordToTransaction(record)
		if err != nil {
			log.Println("Error al analizar el registro del archivo CSV:", err)
			continue
		}

		// Agregar la transacción a la lista
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

// parseRecordToTransaction convierte un registro CSV en una transacción.
func parseRecordToTransaction(record []string) (models.Transaction, error) {
	var transaction models.Transaction

	// Verificar si la fecha es igual a "date" y omitirla
	if record[1] == "date" {
		return transaction, nil
	}

	// Parsear la fecha
	date, err := time.Parse("1/2/2006", record[1]) // Formato: mes/día/año
	if err != nil {
		return transaction, err
	}

	// Parsear la cantidad de la transacción
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

// CreateSummaryCSV crea un archivo CSV con los datos necesarios para generar el resumen.
// Retorna el path del archivo CSV creado.
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

// calculateAverages calcula el promedio de débito y crédito para una lista de transacciones.
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
