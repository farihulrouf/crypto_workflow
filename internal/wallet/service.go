package wallet

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

func (s *Service) CreateWallet(
	userID string,
	asset string,
) error {

	wallet := &Wallet{
		ID:      uuid.New(),
		UserID:  uuid.MustParse(userID),
		Asset:   asset,
		Balance: 0,
	}

	return s.repo.Create(wallet)
}

func (s *Service) GetWallets(
	userID string,
) ([]Wallet, error) {

	return s.repo.FindAllByUser(userID)
}
