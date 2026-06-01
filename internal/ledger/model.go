package ledger

import (
	"time"

	"github.com/google/uuid"
)

type Ledger struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`

	UserID uuid.UUID `gorm:"type:uuid;not null;index"`

	Asset string `gorm:"size:20;not null"`

	Type string `gorm:"size:30;not null"`

	Amount float64 `gorm:"type:numeric(20,8);not null"`

	CreatedAt time.Time
}

func (Ledger) TableName() string {
	return "ledgers"
}
