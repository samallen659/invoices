package user

import (
	"context"
	"errors"
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
	var (
		idStr     string
		firstName string
		lastName  string
		email     string
		userName  string
	)
	var ok bool

	if firstName, ok = profile["given_name"].(string); !ok {
		return errors.New("Unable to validate given_name in profile")
	}
	if lastName, ok = profile["family_name"].(string); !ok {
		return errors.New("Unable to validate family_name in profile")
	}
	if email, ok = profile["email"].(string); !ok {
		return errors.New("Unable to validate email in profile")
	}
	if userName, ok = profile["cognito:username"].(string); !ok {
		return errors.New("Unable to validate cognito:username in profile")
	}
	if idStr, ok = profile["sub"].(string); !ok {
		return errors.New("Unable to validate sub in profile")
	}

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
			if err != nil {
				return err
			}
			return nil
		} else {
			return err
		}
	}

	if user.FirstName != firstName || user.LastName != lastName || user.Email != email || user.UserName != userName {
		updatedUser, err := NewUser(id, firstName, lastName, email, userName)
		err = s.repo.UpdateUser(ctx, *updatedUser)
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

func (s *Service) DeleteUser(ctx context.Context, id uuid.UUID) error {
	err := s.repo.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
