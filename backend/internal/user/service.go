package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/samallen659/invoices/backend/internal/auth"
)

type Service struct {
	auth *auth.Authenticator
	repo Repository
}

func NewService(auth *auth.Authenticator, repo Repository) (*Service, error) {
	return &Service{auth: auth, repo: repo}, nil
}

func (s *Service) ValidateLocalUser(ctx context.Context, profile map[string]any) error {
	firstName := profile["given_name"].(string)
	lastName := profile["family_name"].(string)
	email := profile["email"].(string)
	userName := profile["cognito:username"].(string)
	idStr := profile["sub"].(string)
	id, err := uuid.Parse(idStr)
	if err != nil {
		return err
	}

	user, err := s.repo.GetUser(ctx, id)
	if err != nil {
		if err.Error() == "No User found with supplied id" {
			newUser, err := NewUser(id, firstName, lastName, email, userName)
			if err != nil {
				return err
			}
			err = s.repo.StoreUser(ctx, *newUser)
		} else {

			return err
		}
	}

	if user.FirstName != firstName || user.LastName != lastName || user.Email != email || user.UserName != userName {
		user.FirstName = firstName
		user.LastName = lastName
		user.Email = email
		user.UserName = userName
		err = s.repo.UpdateUser(ctx, *user)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) GetUser(ctx context.Context, id uuid.UUID) (*User, error) {
	user, err := s.repo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
