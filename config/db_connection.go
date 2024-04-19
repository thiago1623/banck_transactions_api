package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB global variable of database
var DB *gorm.DB

// ConnectDB this function config the database with postgres
func ConnectDB() error {
	cfg := LoadSettings()
	dbSection := cfg.Section("Database")
	var (
		DBHost       = dbSection.Key("DBHost").String()
		userName     = dbSection.Key("Username").String()
		Password     = dbSection.Key("Password").String()
		DatabaseName = dbSection.Key("DatabaseName").String()
		DBPort, _    = dbSection.Key("DBPort").Int()
	)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC", DBHost, userName, Password, DatabaseName, DBPort)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
