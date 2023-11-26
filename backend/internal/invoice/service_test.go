package invoice_test

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/samallen659/invoices/backend/internal/invoice"
	"testing"
	"time"
)

type InvoiceRepoStub struct {
	invoices []*invoice.Invoice
}

func (i *InvoiceRepoStub) GetInvoiceByID(ctx context.Context, id uuid.UUID) (*invoice.Invoice, error) {
	for _, inv := range i.invoices {
		if inv.ID == id {
			return inv, nil
		}
	}
	return nil, errors.New("No matching invoice")
}

func (i *InvoiceRepoStub) GetAllInvoices(ctx context.Context) ([]*invoice.Invoice, error) {
	return i.invoices, nil
}

func (i *InvoiceRepoStub) StoreInvoice(ctx context.Context, inv invoice.Invoice) error {
	i.invoices = append(i.invoices, &inv)
	return nil
}

func (i *InvoiceRepoStub) UpdateInvoice(ctx context.Context, inv *invoice.Invoice) error {
	for j, invoices := range i.invoices {
		if invoices.ID == inv.ID {
			i.invoices[j] = inv
			return nil
		}
	}
	return errors.New("failed to update invoice in database")
}

func (i *InvoiceRepoStub) DeleteInvoice(ctx context.Context, id uuid.UUID) error {
	for j := range i.invoices {
		if i.invoices[j].ID == id {
			i.invoices = append(i.invoices[:j], i.invoices[j+1:]...)
			return nil
		}
	}
	return errors.New("invoice with provided id not found for deletion")
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
	irStub := InvoiceRepoStub{}
	svc, _ := invoice.NewService(&irStub)
	ctx := context.Background()

	t.Run("GetByID errors on non matching ID", func(t *testing.T) {
		id, _ := uuid.NewRandom()
		_, err := svc.GetByID(ctx, id)
		if err == nil {
			t.Error("Expected error for non match id")
		}
	})
	t.Run("GetByID returns matched invoice", func(t *testing.T) {
		id := addInvoice(t, &irStub)
		in, err := svc.GetByID(ctx, id)
		if err != nil {
			t.Fatalf("Received error when none expected: %s", err.Error())
		}

		if in == nil {
			t.Fatal("Received invoice is nil")
		}

		if in.ID != id {
			t.Errorf("ID of matched invoice not the same as provided ID, received: %s expected: %s", in.ID.String(), id.String())
		}

	})
	t.Run("GetAll returns full list of invoices", func(t *testing.T) {
		clearInvoices(t, &irStub)
		ids := []uuid.UUID{}
		for i := 0; i < 3; i++ {
			id := addInvoice(t, &irStub)
			ids = append(ids, id)
		}

		invoices, _ := svc.GetAll(ctx)
		if len(invoices) != len(ids) {
			t.Fatalf("Length of returned invoices does not match ids")
		}
		for i := range ids {
			if ids[i] != invoices[i].ID {
				t.Error("Returned invoices ID does not match")
			}
		}
	})
	t.Run("NewInvoice saves an invoiceRequests as an invoice to the repository", func(t *testing.T) {
		clearInvoices(t, &irStub)
		invRq := createInvoiceRequest(t)
		inv, err := svc.NewInvoice(ctx, invRq)

		if err != nil {
			t.Fatalf("Received error when none expected: %s", err.Error())
		}
		if inv == nil {
			t.Fatal("Returned invoice is nil")
		}

		checkInvoiceAgainstInvoiceRequest(t, inv, invRq)

	})
	t.Run("UpdateInvoice saves changes to an invoice to the repository", func(t *testing.T) {
		clearInvoices(t, &irStub)
		id := addInvoice(t, &irStub)

		invRq := createInvoiceRequest(t)
		inv, err := svc.UpdateInvoice(ctx, id, invRq)

		if err != nil {
			t.Fatalf("Received error when none expected: %s", err.Error())
		}
		if inv == nil {
			t.Fatal("Returned invoice is nil")
		}

		checkInvoiceAgainstInvoiceRequest(t, inv, invRq)
	})
	t.Run("DeleteInvoice deletes invoice from the repository", func(t *testing.T) {
		clearInvoices(t, &irStub)
		id := addInvoice(t, &irStub)

		err := svc.DeleteInvoice(ctx, id)
		if err != nil {
			t.Fatalf("Received error when none expected: %s", err.Error())
		}

		for _, inv := range irStub.invoices {
			if inv.ID == id {
				t.Fatalf("Invoice with id %s found in repository when expected it to have been deleted", id.String())
			}
		}
	})
}

func addInvoice(t testing.TB, irs *InvoiceRepoStub) uuid.UUID {
	t.Helper()

	in := invoice.NewInvoice()
	id := in.ID
	irs.invoices = append(irs.invoices, in)
	return id
}

func clearInvoices(t testing.TB, irs *InvoiceRepoStub) {
	t.Helper()

	irs.invoices = []*invoice.Invoice{}
}

func createInvoiceRequest(t testing.TB) invoice.InvoiceRequest {
	t.Helper()

	paymentDue := time.Now().Add(24 * time.Hour)
	invRq := invoice.InvoiceRequest{
		PaymentDue:   paymentDue,
		Description:  "testing",
		PaymentTerms: 2,
		ClientName:   "Sam Allen",
		ClientEmail:  "sam@email.com",
		Status:       "pending",
		ClientAddress: struct {
			Street   string `json:"street"`
			City     string `json:"city"`
			PostCode string `json:"postCode"`
			Country  string `json:"country"`
		}{
			Street:   "15 Eastfield Road",
			City:     "Doncaster",
			PostCode: "DN9 1JQ",
			Country:  "United Kingdom",
		},
		SenderAddress: struct {
			Street   string `json:"street"`
			City     string `json:"city"`
			PostCode string `json:"postCode"`
			Country  string `json:"country"`
		}{
			Street:   "77 High Street",
			City:     "Doncaster",
			PostCode: "DN9 1JS",
			Country:  "United Kingdom",
		},
		Items: []struct {
			Name     string  `json:"name"`
			Quantity int     `json:"quantity"`
			Price    float64 `json:"price"`
		}{{
			Name:     "Web Design",
			Quantity: 1,
			Price:    1499.99,
		}},
	}

	return invRq
}

func checkInvoiceAgainstInvoiceRequest(t testing.TB, inv *invoice.Invoice, invRq invoice.InvoiceRequest) {

	testParams := []struct {
		reqParam any
		invParam any
		error    string
	}{{
		reqParam: invRq.Description,
		invParam: inv.Description,
		error:    "Descriptions do not match",
	}, {
		reqParam: invRq.PaymentTerms,
		invParam: inv.PaymentTerms,
		error:    "PaymentTerms do not match",
	}, {
		reqParam: invRq.ClientName,
		invParam: inv.Client.ClientName,
		error:    "ClientNames do not match",
	}, {
		reqParam: invRq.ClientEmail,
		invParam: inv.Client.ClientEmail,
		error:    "ClientEmails do not match",
	}, {
		reqParam: invRq.Status,
		invParam: inv.Status,
		error:    "Statuses do not match",
	}, {
		reqParam: invRq.PaymentDue,
		invParam: inv.PaymentDue,
		error:    "PaymentDue do not match",
	}, {
		reqParam: invRq.ClientAddress.Street,
		invParam: inv.ClientAddress.Street,
		error:    "ClientAddress Street do not match",
	}, {
		reqParam: invRq.ClientAddress.City,
		invParam: inv.ClientAddress.City,
		error:    "ClientAddress City do not match",
	}, {
		reqParam: invRq.ClientAddress.PostCode,
		invParam: inv.ClientAddress.PostCode,
		error:    "ClientAddress PostCode do not match",
	}, {
		reqParam: invRq.ClientAddress.Country,
		invParam: inv.ClientAddress.Country,
		error:    "ClientAddress Country do not match",
	}, {
		reqParam: invRq.SenderAddress.Street,
		invParam: inv.SenderAddress.Street,
		error:    "SenderAddress Street do not match",
	}, {
		reqParam: invRq.SenderAddress.City,
		invParam: inv.SenderAddress.City,
		error:    "SenderAddress City do not match",
	}, {
		reqParam: invRq.SenderAddress.PostCode,
		invParam: inv.SenderAddress.PostCode,
		error:    "SenderAddress PostCode do not match",
	}, {
		reqParam: invRq.SenderAddress.Country,
		invParam: inv.SenderAddress.Country,
		error:    "SenderAddress Country do not match",
	},
	}

	for _, tp := range testParams {
		if tp.reqParam != tp.invParam {
			t.Errorf(tp.error)
		}
	}

	if len(invRq.Items) != len(inv.InvoiceItems) {
		t.Errorf("Incorrect number of InvoiceItems on Invoice, received: %d expected: %d", len(inv.InvoiceItems), len(invRq.Items))
	}
	for i := range invRq.Items {
		if invRq.Items[i].Name != inv.InvoiceItems[i].Item.Name {
			t.Errorf("Items names do not match")
		}
		if invRq.Items[i].Price != inv.InvoiceItems[i].Item.Price {
			t.Errorf("Items Price do not match")
		}
		if invRq.Items[i].Quantity != inv.InvoiceItems[i].Quantity {
			t.Errorf("Items quantity do not match")
		}
	}

}
