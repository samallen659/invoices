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

func (s *Service) LoginUser(ctx context.Context, loginReq LoginRequest) (bool, error) {
	var user *User
	var err error
	if loginReq.UserName != "" {
		user, err = s.repo.GetUserWithUsername(ctx, loginReq.UserName)
		if err != nil {
			return false, err
		}
	} else {
		user, err = s.repo.GetUserWithEmail(ctx, loginReq.Email)
		if err != nil {
			return false, err
		}
	}

	pwdMatch, err := VerifyPassword(loginReq.Password, user.Password)
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
	var user *User
	var err error
	if userName != "" {
		user, err = s.repo.GetUserWithUsername(ctx, userName)
		if err != nil {
			return nil, err
		}
	} else {
		user, err = s.repo.GetUserWithEmail(ctx, email)
		if err != nil {
			return nil, err
		}
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
