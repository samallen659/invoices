package invoice_test

import (
	"github.com/samallen659/invoices/backend/internal/invoice"
	"testing"
)

func TestInvoice(t *testing.T) {
	t.Run("NewInvoice returns an Invoice with UUID set ID", func(t *testing.T) {
		in := invoice.NewInvoice()
		if in.ID.String() == "00000000-0000-0000-0000-000000000000" {
			t.Errorf("ID on invoice not set to valid UUID: %s", in.ID.String())
		}
	})
}
