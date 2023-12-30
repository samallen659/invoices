package user

import (
	"github.com/samallen659/invoices/backend/internal/auth"
)

type Service struct {
	auth *auth.Authenticator
	repo Repository
}

func NewService(auth *auth.Authenticator, repo Repository) (*Service, error) {
	return &Service{auth: auth, repo: repo}, nil
}
