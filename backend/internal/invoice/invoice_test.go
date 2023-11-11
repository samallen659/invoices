package invoice_test

import (
	"github.com/samallen659/invoices/backend/internal/invoice"
	"testing"
)

func TestInvoice(t *testing.T) {
	t.Run("NewInvoice returns an Invoice with UUID set ID, CreatedAt set and 0 Money for total", func(t *testing.T) {
		in := invoice.NewInvoice()

		if in.ID.String() == "00000000-0000-0000-0000-000000000000" {
			t.Errorf("ID on invoice not set to valid UUID: %s", in.ID.String())
		}

		if in.CreatedAt.String() == "0001-01-01 00:00:00 +0000 UTC" {
			t.Errorf("CreatedAt on invoice not set to valid time: %s", in.CreatedAt.String())
		}

		if in.Total.Currency().Code != "GBP" {
			t.Errorf("Total not initialised as GBP")
		}
	})
}
