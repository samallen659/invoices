package user

import (
	"github.com/samallen659/invoices/backend/internal/auth"
)

type Service struct {
	auth *auth.Authenticator
}

func NewService(auth *auth.Authenticator) (*Service, error) {
	return &Service{auth: auth}, nil
}
