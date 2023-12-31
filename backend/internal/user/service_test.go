package user_test

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/samallen659/invoices/backend/internal/auth"
	u "github.com/samallen659/invoices/backend/internal/user"
	"testing"
)

type RepoStub struct {
	users []*u.User
}

func (rs *RepoStub) GetUser(ctx context.Context, id uuid.UUID) (*u.User, error) {
	for _, usr := range rs.users {
		if usr.ID.String() == id.String() {
			return usr, nil
		}
	}

	return nil, errors.New("No User found with supplied id")
}

func (rs *RepoStub) StoreUser(ctx context.Context, user u.User) error {
	rs.users = append(rs.users, &user)
	return nil
}

func (rs *RepoStub) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return nil
}

func (rs *RepoStub) UpdateUser(ctx context.Context, user u.User) error {
	for _, usr := range rs.users {
		if usr.ID.String() == user.ID.String() {
			usr.FirstName = user.FirstName
			usr.LastName = user.LastName
			usr.Email = user.Email
			usr.UserName = user.UserName
			return nil
		}
	}

	return errors.New("failed to update users")
}

func TestService(t *testing.T) {
	rs := &RepoStub{}
	a, _ := auth.NewAuthenticator()
	s, err := u.NewService(a, rs)
	if err != nil {
		t.Fatal("failed to initialise service")
	}
	id, err := uuid.NewRandom()
	if err != nil {
		t.Fatal("failed to create id for user")
	}
	firstName, lastName, email, userName := "testFirst", "testLast", "test@email.com", "testUser"
	user, err := u.NewUser(id, firstName, lastName, email, userName)
	if err != nil {
		t.Fatal("failed to create user")
	}

	rs.users = append(rs.users, user)

	t.Run("GetUser returns error when no user found", func(t *testing.T) {
		id, _ := uuid.NewRandom()
		ctx := context.Background()
		_, err := s.GetUser(ctx, id)
		if err == nil {
			t.Error("Expected error, none received")
		}
	})
	t.Run("GetUser returns user when passed valid id", func(t *testing.T) {
		ctx := context.Background()
		usr, err := s.GetUser(ctx, id)
		if err != nil {
			t.Fatalf("Error received when none expected: %s", err.Error())
		}
		if usr == nil {
			t.Fatal("User received is nil")
		}
		if usr.FirstName != firstName || usr.LastName != lastName || usr.Email != email || usr.UserName != userName {
			t.Error("Returned user details are incorrect")
		}
	})
	t.Run("ValidateLocalUser returns error on invalid profile", func(t *testing.T) {
		testParams := []struct {
			profile map[string]any
			err     string
		}{
			{
				profile: map[string]any{
					"given_name":  firstName,
					"family_name": lastName,
					"email":       email,
					"sub":         id.String(),
				},
				err: "Expected error for invalid cognito:username in profile",
			},
			{
				profile: map[string]any{
					"cognito:username": "testUser",
					"family_name":      lastName,
					"email":            email,
					"sub":              id.String(),
				},
				err: "Expected error for invalid given_name in profile",
			},
			{
				profile: map[string]any{
					"cognito:username": "testUser",
					"given_name":       firstName,
					"email":            email,
					"sub":              id.String(),
				},
				err: "Expected error for invalid family_name in profile",
			},
			{
				profile: map[string]any{
					"cognito:username": "testUser",
					"given_name":       firstName,
					"family_name":      lastName,
					"sub":              id.String(),
				},
				err: "Expected error for invalid email in profile",
			},
			{
				profile: map[string]any{
					"cognito:username": "testUser",
					"given_name":       firstName,
					"family_name":      lastName,
					"email":            email,
				},
				err: "Expected error for invalid sub in profile",
			},
			{
				profile: map[string]any{
					"cognito:username": "testUser",
					"given_name":       firstName,
					"family_name":      lastName,
					"email":            email,
					"sub":              "invalid id",
				},
				err: "Expected error for invalid sub in profile",
			},
		}

		for _, tt := range testParams {
			ctx := context.Background()
			err := s.ValidateLocalUser(ctx, tt.profile)
			if err == nil {
				t.Error(tt.err)
			}
		}
	})
	t.Run("ValidateLocalUser return nil error if user currently in repository", func(t *testing.T) {
		profile := map[string]any{
			"cognito:username": userName,
			"given_name":       firstName,
			"family_name":      lastName,
			"email":            email,
			"sub":              id.String(),
		}
		ctx := context.Background()
		err := s.ValidateLocalUser(ctx, profile)
		if err != nil {
			t.Errorf("Received error when none expected: %s", err.Error())
		}
	})
	t.Run("ValidateLocalUser Updates user if found but fields differ", func(t *testing.T) {
		userName, firstName, lastName, email := "diffUser", "diffFirst", "diffLast", "diff@email.com"
		profile := map[string]any{
			"cognito:username": userName,
			"given_name":       firstName,
			"family_name":      lastName,
			"email":            email,
			"sub":              id.String(),
		}
		ctx := context.Background()
		err := s.ValidateLocalUser(ctx, profile)
		if err != nil {
			t.Errorf("Received error when none expected: %s", err.Error())
		}

		usr, _ := rs.GetUser(ctx, id)

		if usr.FirstName != firstName || usr.LastName != lastName || usr.Email != email || usr.UserName != userName {
			t.Error("Returned user details have not been updated")
		}

	})
	t.Run("ValidateLocalUser stores user if not currently in repository", func(t *testing.T) {
		id, _ := uuid.NewRandom()
		userName, firstName, lastName, email := "newUser", "newFirst", "newLast", "new@email.com"
		profile := map[string]any{
			"cognito:username": userName,
			"given_name":       firstName,
			"family_name":      lastName,
			"email":            email,
			"sub":              id.String(),
		}
		ctx := context.Background()
		err := s.ValidateLocalUser(ctx, profile)
		if err != nil {
			t.Fatalf("Received error when none expected: %s", err.Error())
		}

		usr, _ := rs.GetUser(ctx, id)
		if usr == nil {
			t.Fatal("New User not stored by repo")
		}

		if usr.FirstName != firstName || usr.LastName != lastName || usr.Email != email || usr.UserName != userName {
			t.Error("New User store with incorrect fields")
		}
	})
}
