package invoice_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/samallen659/invoices/backend/internal/invoice"
	"testing"
)

type InvoiceRepoStub struct{}

func (i *InvoiceRepoStub) GetInvoiceByID(ctx context.Context, id uuid.UUID) (*invoice.Invoice, error) {
	return nil, nil
}

func (i *InvoiceRepoStub) GetAllInvoices(ctx context.Context) ([]*invoice.Invoice, error) {
	return nil, nil
}

func (i *InvoiceRepoStub) StoreInvoice(ctx context.Context, inv invoice.Invoice) error {
	return nil
}

func TestNewService(t *testing.T) {
	irStub := InvoiceRepoStub{}
	svc, err := invoice.NewService(&irStub)
	if err != nil {
		t.Errorf("NewService returned an error: %s", err.Error())
	}
	if svc == nil {
		t.Error("NewService return a nil Service")
	}
}

func TestService(t *testing.T) {
	//TODO
}
