package app

import (
	"crypto-flow/internal/auth"
	"crypto-flow/internal/ledger"
	"crypto-flow/internal/wallet"

	"gorm.io/gorm"
)

type Container struct {
	AuthHandler   *auth.Handler
	WalletHandler *wallet.Handler
	LedgerHandler *ledger.Handler
}

func NewContainer(
	db *gorm.DB,
) *Container {

	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	walletRepo := wallet.NewRepository(db)
	walletService := wallet.NewService(walletRepo)
	walletHandler := wallet.NewHandler(walletService)

	ledgerRepo := ledger.NewRepository(db)
	ledgerService := ledger.NewService(ledgerRepo)
	ledgerHandler := ledger.NewHandler(ledgerService)

	return &Container{
		AuthHandler:   authHandler,
		WalletHandler: walletHandler,
		LedgerHandler: ledgerHandler,
	}
}
