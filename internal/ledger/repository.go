package ledger

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(
	db *gorm.DB,
) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(
	ledger *Ledger,
) error {
	return r.db.Create(ledger).Error
}

func (r *Repository) FindByUserID(
	userID uuid.UUID,
) ([]Ledger, error) {

	var ledgers []Ledger

	err := r.db.
		Where("user_id = ?", userID).
		Order("created_at desc").
		Find(&ledgers).
		Error

	return ledgers, err
}
