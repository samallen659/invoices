package invoice_test

import (
	"testing"

	"github.com/samallen659/invoices/backend/internal/invoice"
)

func TestNewClient(t *testing.T) {
	t.Run("Empty name value", func(t *testing.T) {
		_, err := invoice.NewClient("", "email")
		if err == nil {
			t.Error("Expected error for emtpy name value")
		}
	})
	t.Run("Empty email value", func(t *testing.T) {
		_, err := invoice.NewClient("name", "")
		if err == nil {
			t.Error("Expected error for emtpy email value")
		}
	})
	t.Run("Returns valid Client", func(t *testing.T) {
		name, email := "name", "email@email.com"
		client, err := invoice.NewClient(name, email)
		if err != nil {
			t.Errorf("Recieved error when none expected: %s", err.Error())
		}

		if client == nil {
			t.Fatal("Returned client is nil")
		}

		if client.ClientName != name {
			t.Errorf("Name does not match provieded value, received: %s expected: %s",
				client.ClientName, name)
		}
		if client.ClientEmail != email {
			t.Errorf("Email does not match provided value, received: %s expected: %s", client.ClientEmail, email)
		}
	})
}
