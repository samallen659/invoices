package user

import (
	"context"
	// "fmt"
)

type Service struct {
	auth *CognitoAuthenticator
}

func NewService(auth *CognitoAuthenticator) (*Service, error) {
	return &Service{auth: auth}, nil
}

func (s *Service) GetLoginURL() string {
	return s.auth.GetLoginURL()
}

func (s *Service) GetAccessToken(ctx context.Context, authCode string) (string, error) {
	token, err := s.auth.GetAccessToken(ctx, authCode)
	if err != nil {
		return "", err
	}
	return token, nil
}
