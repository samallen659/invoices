package invoice

import (
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
