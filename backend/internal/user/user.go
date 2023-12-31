package user

import (
	"errors"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	UserName  string    `json:"userName"`
}

func NewUser(id uuid.UUID, firstName string, lastName string, email string, userName string) (*User, error) {
	if firstName == "" {
		return nil, errors.New("firstName cannot be empty")
	}
	if lastName == "" {
		return nil, errors.New("lastName cannot be empty")
	}
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}
	if userName == "" {
		return nil, errors.New("userName cannot by empty")
	}

	return &User{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		UserName:  userName,
	}, nil
}
