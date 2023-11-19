package invoice_test

import (
	"github.com/samallen659/invoices/backend/internal/invoice"
	"testing"
)

func TestNewInvoice(t *testing.T) {
	t.Run("NewInvoice returns an Invoice with UUID set ID, CreatedAt set and 0 Money for total", func(t *testing.T) {
		in := invoice.NewInvoice()

		if in.ID.String() == "00000000-0000-0000-0000-000000000000" {
			t.Errorf("ID on invoice not set to valid UUID: %s", in.ID.String())
		}

		if in.CreatedAt.String() == "0001-01-01 00:00:00 +0000 UTC" {
			t.Errorf("CreatedAt on invoice not set to valid time: %s", in.CreatedAt.String())
		}
	})
}

func TestInvoice(t *testing.T) {
	t.Run("SetDescription errors on empty value", func(t *testing.T) {
		in := invoice.NewInvoice()
		err := in.SetDescription("")
		if err == nil {
			t.Error("Expected error for empty description value")
		}
	})
	t.Run("SetDescription sets description on invoice to provided value", func(t *testing.T) {
		description := "test description"
		in := invoice.NewInvoice()
		in.SetDescription(description)
		if in.Description != description {
			t.Errorf("Description does not match provided value, received: %s expected: %s", in.Description, description)
		}
	})
	t.Run("SetPaymentTerms errors on paymentTerm < 1", func(t *testing.T) {
		in := invoice.NewInvoice()
		err := in.SetPaymentTerms(0)
		if err == nil {
			t.Error("Expected error for paymentTerm less than 1")
		}
	})
	t.Run("SetPaymentTerms errors on paymentTerm > 30", func(t *testing.T) {
		in := invoice.NewInvoice()
		err := in.SetPaymentTerms(31)
		if err == nil {
			t.Error("Expected error for paymentTerm greater than 30")
		}
	})
	t.Run("SetPaymentTerms sets paymentTerms on invoice to provided value", func(t *testing.T) {
		paymentTerms := 7
		in := invoice.NewInvoice()
		in.SetPaymentTerms(paymentTerms)
		if in.PaymentTerms != paymentTerms {
			t.Errorf("PaymentTerms does not match provided value, received: %d expected: %d", in.PaymentTerms, paymentTerms)
		}
	})
	t.Run("SetStatus errors on status not matching pending|paid|draft", func(t *testing.T) {
		in := invoice.NewInvoice()
		err := in.SetStatus("something")
		if err == nil {
			t.Error("Expected error for status not matching pending|paid|draft")
		}
	})
	t.Run("SetStatus sets status on invoice to provided value", func(t *testing.T) {
		statuses := []invoice.InvoiceStatus{"pending", "paid", "draft"}
		in := invoice.NewInvoice()
		for _, status := range statuses {
			in.SetStatus(status)
			if in.Status != status {
				t.Errorf("Status does not match provided value, received: %s expected: %s", in.Status, status)
			}
		}
	})
	t.Run("SetInvoiceItems errors on emtpy invoiceItems", func(t *testing.T) {
		invoiceItems := []invoice.InvoiceItem{}
		in := invoice.NewInvoice()
		err := in.SetInvoiceItems(invoiceItems)
		if err == nil {
			t.Error("Expected error for empty invoiceItems")
		}
	})
	t.Run("SetInvoiceItems sets invoiceItems on invoice to provided value", func(t *testing.T) {
		invoiceItems := []invoice.InvoiceItem{}
		invoiceItem, _ := invoice.NewInvoiceItem("test", 9.9, 1)
		invoiceItems = append(invoiceItems, *invoiceItem)
		in := invoice.NewInvoice()
		in.SetInvoiceItems(invoiceItems)
		if in.InvoiceItems[0] != *invoiceItem {
			t.Errorf("InvoiceItems does not match provided value")
		}
	})
	t.Run("SetClientAddress sets ClientAddress on invoice to provided value", func(t *testing.T) {
		clientAddress, _ := invoice.NewAddress("Street", "City", "Postcode", "Country")
		in := invoice.NewInvoice()
		in.SetClientAddress(*clientAddress)
		if in.ClientAddress != *clientAddress {
			t.Error("ClientAddress does not match provided value")
		}
	})
	t.Run("SetSenderAddress sets SenderAddress on invoice to provided value", func(t *testing.T) {
		senderAddress, _ := invoice.NewAddress("Street", "City", "Postcode", "Country")
		in := invoice.NewInvoice()
		in.SetSenderAddress(*senderAddress)
		if in.SenderAddress != *senderAddress {
			t.Error("SenderAddress does not match provided value")
		}
	})
}
