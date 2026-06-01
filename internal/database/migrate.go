package database

import (
	"crypto-flow/internal/auth"
	"crypto-flow/internal/wallet"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&auth.User{},
		&wallet.Wallet{},
	)
}
