package invoice_test

import (
	"github.com/samallen659/invoices/backend/internal/invoice"
	"testing"
)

func TestInvoiceItemFactory(t *testing.T) {
	t.Run("Pass empty name string", func(t *testing.T) {
		_, err := invoice.InvoiceItemFactory("", 1.0, 1)
		if err == nil {
			t.Error("Expected error for emtpy name string")
		}
	})
	t.Run("Pass less than 0.0 price", func(t *testing.T) {
		_, err := invoice.InvoiceItemFactory("test", -1.0, 1)
		if err == nil {
			t.Error("Expected error for less than 0.0 price")
		}
	})
	t.Run("Pass less than 1 quantity", func(t *testing.T) {
		_, err := invoice.InvoiceItemFactory("test", 1.0, 0)
		if err == nil {
			t.Error("Expected error for less than 1 quantity")
		}
	})

	t.Run("Returns an InvoiceItem with nested Item", func(t *testing.T) {
		name, price, quantity := "Item", 25.99, 3
		it, err := invoice.InvoiceItemFactory(name, price, quantity)
		if err != nil {
			t.Errorf("Recieved an error: %s", err.Error())
		}

		if it == nil {
			t.Error("Received nil InvoiceItem")
		}

		if it.Item.Name != name {
			t.Errorf("Incorrect name, expected: %s recieved: %s", name, it.Item.Name)
		}

		if it.Item.Price != price {
			t.Errorf("Incorrect price, expected: %f received: %f", price, it.Item.Price)
		}

		if it.Quantity != quantity {
			t.Errorf("Incorrect quantity, expected: %d received: %d", quantity, it.Quantity)
		}
	})
}
