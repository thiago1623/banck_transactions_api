package models

import (
	"gorm.io/gorm"
	"time"
)

// Transaction is a table of database
type Transaction struct {
	gorm.Model
	Date        time.Time
	Transaction float64
}

// TableName rename table
func (Transaction) TableName() string {
	return "transactions"
}
