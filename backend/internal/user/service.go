package user

import (
	"context"
	"fmt"
)

type Service struct {
	auth *CognitoAuthenticator
}

func NewService(auth *CognitoAuthenticator) (*Service, error) {
	return &Service{auth: auth}, nil
}

func (s *Service) SignUp(ctx context.Context, req SignUpRequest) error {
	out, err := s.auth.SignUp(ctx, req.Email, req.FirstName, req.LastName, req.Password)
	if err != nil {
		return err
	}

	fmt.Println(out)
	return nil
}
