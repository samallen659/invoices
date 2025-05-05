package user

import (
	"context"
	"errors"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) (*Service, error) {
	return &Service{repo: repo}, nil
}

func (s *Service) LoginUser(ctx context.Context, email string, userName string, password string) (bool, error) {
	user, err := s.repo.GetUser(ctx, email, userName)
	if err != nil {
		return false, err
	}

	pwdMatch, err := VerifyPassword(password, user.Password)
	if err != nil {
		return false, err
	}

	if !pwdMatch {
		return false, errors.New("Failed to match password")
	}

	return true, nil
}

func (s *Service) SignUpUser(ctx context.Context, signUpReq SignUpRequest) (bool, error) {
	user, err := NewUser(signUpReq.UserName, signUpReq.FirstName, signUpReq.LastName, signUpReq.Email, signUpReq.Password)
	if err != nil {
		return false, err
	}

	err = s.repo.StoreUser(ctx, *user)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *Service) GetUser(ctx context.Context, email string, userName string) (*User, error) {
	user, err := s.repo.GetUser(ctx, email, userName)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) DeleteUser(ctx context.Context, userName string) error {
	err := s.repo.DeleteUser(ctx, userName)
	if err != nil {
		return err
	}

	return nil
}
