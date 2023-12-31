package user_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/samallen659/invoices/backend/internal/user"
)

func TestUser(t *testing.T) {
	t.Run("Fails on empty FirstName", func(t *testing.T) {
		id, _ := uuid.NewRandom()
		_, err := user.NewUser(id, "", "testLast", "test@email.com", "testUesr")
		if err == nil {
			t.Error("Expected error for empty firstName")
		}
	})
	t.Run("Fails on empty LastName", func(t *testing.T) {
		id, _ := uuid.NewRandom()
		_, err := user.NewUser(id, "testFirst", "", "test@email.com", "testUesr")
		if err == nil {
			t.Error("Expected error for empty lastName")
		}
	})
	t.Run("Fails on empty Email", func(t *testing.T) {
		id, _ := uuid.NewRandom()
		_, err := user.NewUser(id, "testFirst", "testLast", "", "testUesr")
		if err == nil {
			t.Error("Expected error for empty email")
		}
	})
	t.Run("Fails on empty UserName", func(t *testing.T) {
		id, _ := uuid.NewRandom()
		_, err := user.NewUser(id, "testFirst", "testLast", "test@email.com", "")
		if err == nil {
			t.Error("Expected error for empty userName")
		}
	})
	t.Run("Returns User with correct fields", func(t *testing.T) {
		id, _ := uuid.NewRandom()
		firstName, lastName, email, userName := "testFirst", "testLast", "test@email.com", "testUser"
		user, err := user.NewUser(id, firstName, lastName, email, userName)
		if err != nil {
			t.Fatalf("Received error when none expected: %s", err.Error())
		}

		if user.FirstName != firstName || user.LastName != lastName ||
			user.Email != email || user.UserName != userName {
			t.Error("Return user fields do not match provided values")
		}
	})
}
