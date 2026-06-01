package auth

import (
	"crypto-flow/internal/config"
	"crypto-flow/internal/pkg/jwt"
	"errors"

	"golang.org/x/crypto/bcrypt"
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

func (s *Service) Register(
	req RegisterRequest,
) error {

	existing, _ := s.repo.FindByEmail(
		req.Email,
	)

	if existing != nil {
		return errors.New(
			"email already exists",
		)
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user := &User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hash),
	}

	return s.repo.Create(user)
}

func (s *Service) Login(
	req LoginRequest,
) (string, error) {

	user, err := s.repo.FindByEmail(
		req.Email,
	)

	if err != nil {
		return "", errors.New(
			"invalid credentials",
		)
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	)

	if err != nil {
		return "", errors.New(
			"invalid credentials",
		)
	}

	token, err := jwt.GenerateToken(
		user.ID.String(),
		config.Get("JWT_SECRET"),
	)

	if err != nil {
		return "", err
	}

	return token, nil
}
