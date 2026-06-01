package ledger

import (
	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func NewService(
	repo *Repository,
) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateLedger(
	userID uuid.UUID,
	req CreateLedgerRequest,
) error {

	ledger := &Ledger{
		UserID: userID,
		Asset:  req.Asset,
		Type:   req.Type,
		Amount: req.Amount,
	}

	return s.repo.Create(
		ledger,
	)
}

func (s *Service) GetLedgers(
	userID uuid.UUID,
) ([]Ledger, error) {

	return s.repo.FindByUserID(
		userID,
	)
}
