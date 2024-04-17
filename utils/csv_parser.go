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
