package invoice_test

import (
	"testing"

	"github.com/samallen659/invoices/backend/internal/invoice"
)

func TestNewAddress(t *testing.T) {
	street := "The Street"
	city := "Doncaster"
	postCode := "DN1 1AB"
	country := "United Kingdom"
	t.Run("Empty street", func(t *testing.T) {
		_, err := invoice.NewAddress("", city, postCode, country)
		if err == nil {
			t.Error("Expected error for emtpy street value")
		}
	})
	t.Run("Empty city", func(t *testing.T) {
		_, err := invoice.NewAddress(street, "", postCode, country)
		if err == nil {
			t.Error("Expected error for emtpy city value")
		}
	})
	t.Run("Empty postCode", func(t *testing.T) {
		_, err := invoice.NewAddress(street, city, "", country)
		if err == nil {
			t.Error("Expected error for emtpy postCode value")
		}
	})
	t.Run("Empty country", func(t *testing.T) {
		_, err := invoice.NewAddress(street, city, postCode, "")
		if err == nil {
			t.Error("Expected error for emtpy country value")
		}
	})
	t.Run("Returns a valid Address", func(t *testing.T) {
		add, err := invoice.NewAddress(street, city, postCode, country)
		if err != nil {
			t.Fatalf("Received error when none expected: %s", err.Error())
		}

		if add == nil {
			t.Fatal("Returned address is nil")
		}
		if add.Street != street {
			t.Fatalf("Street does not match provided value, received: %s expected: %s", add.Street, street)
		}
		if add.City != city {
			t.Fatalf("City does not match provided value, received: %s expected: %s", add.City, city)
		}
		if add.PostCode != postCode {
			t.Fatalf("PostCode does not match provided value, received: %s expected: %s", add.PostCode, postCode)
		}
		if add.Country != country {
			t.Fatalf("Country does not match provided value, received: %s expected: %s", add.Country, country)
		}
	})
}
