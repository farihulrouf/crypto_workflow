package wallet

import (
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`

	UserID uuid.UUID `gorm:"type:uuid;not null;index"`

	Asset string `gorm:"size:20;not null"`

	Balance float64 `gorm:"type:numeric(20,8);default:0"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Wallet) TableName() string {
	return "wallets"
}
