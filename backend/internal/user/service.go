package user

import (
	"context"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) (*Service, error) {
	return &Service{repo: repo}, nil
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
