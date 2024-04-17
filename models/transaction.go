package models

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	gorm.Model
	Date        time.Time
	Transaction float64
}

func (Transaction) TableName() string {
	return "transactions"
}
