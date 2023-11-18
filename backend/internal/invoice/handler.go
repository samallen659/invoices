package invoice

import (
	"time"
)

type InvoiceRequest struct {
	ID            string        `json:"id"`
	PaymentDue    time.Time     `json:"paymentDue"`
	Description   string        `json:"description"`
	PaymentTerms  int           `json:"paymentTerms"`
	ClientName    string        `json:"clientName"`
	ClientEmail   string        `json:"clientEmail"`
	Status        InvoiceStatus `json:"status"`
	ClientAddress struct {
		Street   string `json:"street"`
		City     string `json:"city"`
		PostCode string `json:"postCode"`
		Country  string `json:"country"`
	} `json:"clientAddress"`
	SenderAddress struct {
		Street   string `json:"street"`
		City     string `json:"city"`
		PostCode string `json:"postCode"`
		Country  string `json:"country"`
	} `json:"senderAddress"`
	Items []struct {
		Name     string  `json:"name"`
		Quantity int     `json:"quantity"`
		Price    float64 `json:"price"`
	}
}

type Handler struct {
	svc Service
}

func NewHandler(svc Service) (*Handler, error) {
	return &Handler{svc: svc}, nil
}
