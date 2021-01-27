package config

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Driver for GORM
)

// DB Context
var DB *gorm.DB

// Database Initialization
func Database() (*gorm.DB, error) {
	driver := os.Getenv("DATABASE_DRIVER")
	database := os.Getenv("DATABASE_URL")

	var err error
	DB, err = gorm.Open(driver, database)

	if err != nil {
		return nil, err
	}

	return DB, nil
}
