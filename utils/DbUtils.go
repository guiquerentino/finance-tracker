package utils

import (
	"financial-tracker/structs"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open("db/finance.db"), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("error during database connection")
	}

	err = db.AutoMigrate(&structs.Transaction{})

	if err != nil {
		return nil, fmt.Errorf("error during database migration")
	}

	return db, nil

}
