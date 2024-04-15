package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
	logger *Logger
)

func InitDatabase() error {
	var err error

	// Initialize SQLite Database
	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("Error Intializing SQLite: %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	// Initialize Logger
	logger = NewLogger(prefix)
	return logger
}