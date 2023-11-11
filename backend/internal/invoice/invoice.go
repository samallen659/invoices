package invoice

import (
	"errors"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type InvoiceStatus string

const (
	STATUS_PENDING InvoiceStatus = "pending"
	STATUS_DRAFT                 = "draft"
	STATUS_PAID                  = "paid"
)

type Invoice struct {
	ID           uuid.UUID
	Description  string
	CreatedAt    time.Time
	PaymentDue   time.Time
	PaymentTerms int
	Status       InvoiceStatus
	Total        money.Money
}

func NewInvoice() *Invoice {
	id := uuid.New()
	return &Invoice{
		ID:        id,
		CreatedAt: time.Now(),
		Status:    STATUS_DRAFT,
	}
}

func (i *Invoice) SetDescription(description string) error {
	if description == "" {
		return errors.New("description cannot be blank")
	}

	i.Description = description
	return nil
}

func (i *Invoice) SetPaymentTerms(paymentTerms int) error {
	if paymentTerms < 1 || paymentTerms > 30 {
		return errors.New("payment terms cannot be less than 1 or greater than 30")
	}

	i.PaymentTerms = paymentTerms
	return nil
}

func (i *Invoice) SetPaymentDue(paymentDue time.Time) error {
	if paymentDue.Before(i.CreatedAt) {
		return errors.New("paymentDue cannot be before createdAt")
	}

	i.PaymentDue = paymentDue
	return nil
}

func (i *Invoice) SetStatus(status InvoiceStatus) {
	i.Status = status
}
