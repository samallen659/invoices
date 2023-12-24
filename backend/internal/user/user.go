package user

import (
	"errors"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
}

func NewUser(firstName string, lastName string, email string) (*User, error) {
	if firstName == "" {
		return nil, errors.New("firstName cannot be empty")
	}
	if lastName == "" {
		return nil, errors.New("firstName cannot be empty")
	}
	if email == "" {
		return nil, errors.New("firstName cannot be empty")
	}

	return &User{
		ID:        uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}, nil
}
