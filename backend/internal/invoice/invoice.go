package invoice

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type InvoiceStatus string

// The three available states of an Invoice
const (
	STATUS_PENDING InvoiceStatus = "pending"
	STATUS_DRAFT                 = "draft"
	STATUS_PAID                  = "paid"
)

// Invoice is the aggregate root for the Invoice domain
type Invoice struct {
	ID            uuid.UUID
	Description   string
	CreatedAt     time.Time
	PaymentDue    time.Time
	PaymentTerms  int
	Status        InvoiceStatus
	Client        Client
	ClientAddress Address
	SenderAddress Address
	InvoiceItems  []InvoiceItem
	Total         float64
}

// Returns a new NewInvoice
//
// ID, CreatedAt, Status and Total initialised
func NewInvoice() *Invoice {
	id := uuid.New()
	return &Invoice{
		ID:        id,
		CreatedAt: time.Now(),
		Status:    STATUS_DRAFT,
	}
}

// Sets Invoice description
func (i *Invoice) SetDescription(description string) error {
	if description == "" {
		return errors.New("description cannot be blank")
	}

	i.Description = description
	return nil
}

// Sets Invoice paymentTerms
func (i *Invoice) SetPaymentTerms(paymentTerms int) error {
	if paymentTerms < 1 || paymentTerms > 30 {
		return errors.New("payment terms cannot be less than 1 or greater than 30")
	}

	i.PaymentTerms = paymentTerms
	return nil
}

// Sets Invoice paymentDue
func (i *Invoice) SetPaymentDue(paymentDue time.Time) error {
	if paymentDue.Before(i.CreatedAt) {
		return errors.New("paymentDue cannot be before createdAt")
	}

	i.PaymentDue = paymentDue
	return nil
}

// Sets Invoice status
func (i *Invoice) SetStatus(status InvoiceStatus) error {
	if status != STATUS_PENDING && status != STATUS_PAID && status != STATUS_DRAFT {
		return errors.New("status can only be of value 'draft', 'paid' or 'pending'")
	}

	i.Status = status
	return nil
}

// Sets Invoice invoiceItems
func (i *Invoice) SetInvoiceItems(invoiceItems []InvoiceItem) error {
	if len(invoiceItems) == 0 {
		return errors.New("invoiceItems cannot be empty")
	}

	i.InvoiceItems = invoiceItems
	i.updateTotal()
	return nil
}

func (i *Invoice) SetClientAddress(addr Address) error {
	i.ClientAddress = addr

	return nil
}

func (i *Invoice) SetSenderAddress(addr Address) error {
	i.SenderAddress = addr

	return nil
}

func (i *Invoice) SetClient(client Client) error {
	i.Client = client

	return nil
}

func (i *Invoice) updateTotal() {
	newTotal := 0.0
	for _, invItem := range i.InvoiceItems {
		newTotal += invItem.Total
	}
	i.Total = newTotal
}
