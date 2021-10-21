package database

import (
	"fmt"

	"github.com/OE-OverEngineer/over-review-backend/configuration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection(dbConfig *configuration.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Bangkok",
		dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, err
}
