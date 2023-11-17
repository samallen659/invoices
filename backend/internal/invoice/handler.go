package invoice

import (
	"time"
)

type InvoiceRequest struct {
	PaymentDue    time.Time `json:"paymentDue"`
	Description   string    `json:"description"`
	PaymentTerms  int       `json:"paymentTerms"`
	ClientName    string    `json:"clientName"`
	ClientEmail   string    `json:"clientEmail"`
	Status        string    `json:"status"`
	SenderAddress struct {
		Street   string `json:"street"`
		City     string `json:"city"`
		PostCode string `json:"postCode"`
		Country  string `json:"country"`
	} `json:"senderAddress"`
	ClientAddress struct {
		Street   string `json:"street"`
		City     string `json:"city"`
		PostCode string `json:"postCode"`
		Country  string `json:"country"`
	} `json:"clientAddress"`
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
