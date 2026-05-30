package database

import (
	"fmt"

	"crypto-flow/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Get("DB_HOST"),
		config.Get("DB_USER"),
		config.Get("DB_PASSWORD"),
		config.Get("DB_NAME"),
		config.Get("DB_PORT"),
		config.Get("DB_SSLMODE"),
	)

	return gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
}
