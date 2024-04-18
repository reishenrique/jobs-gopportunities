package config

import (
	"os"

	"github.com/reishenrique/job-gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/openings.db"
	
	// Check if database file exists
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("Database file not found, creating...")
		// Create database file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}

		file.Close()
	}

	// Create database and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("Error connecting to the SQLite database: %v", err)
		return nil, err
	}

	// Migrate Schema
	err = db.AutoMigrate(&schemas.Opening{}, &schemas.Recruiters{})

	if err != nil {
		logger.Errorf("SQLite Automigration error: %v", err)
		return nil, err
	}

	// Return database
	return db, nil
}