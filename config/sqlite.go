package config

import (
	"os"

	"github.com/notOliveira/apigolang/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	// Check if the database file exists
	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		
		// Create the directory if it doesn't exist
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		// Create the database file
		file , err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}

		// Close the file after creating it
		file.Close()
	}


	logger := GetLogger("sqlite")
	// Create DB
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("Failed to connect to SQLite database: %v", err)
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errorf("Failed to migrate SQLite database: %v", err)
		return nil, err
	}

	logger.Info("SQLite database initialized successfully")

	return db, err
}