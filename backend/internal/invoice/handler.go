package invoice

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
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

type ErrorResponse struct {
	Error string `json:"error"`
}

type InvoiceResponse struct {
	Invoice []*Invoice `json:"invoice"`
}

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) (*Handler, error) {
	return &Handler{svc: svc}, nil
}

func (h *Handler) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	uid, err := uuid.Parse(id)
	if err != nil {
		h.writeJson(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	inv, err := h.svc.GetByID(r.Context(), uid)
	if err != nil {
		h.writeJson(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	h.writeJson(w, http.StatusOK, InvoiceResponse{Invoice: []*Invoice{inv}})
	return
}

func (h *Handler) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	inv, err := h.svc.GetAll(r.Context())
	if err != nil {
		h.writeJson(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	h.writeJson(w, http.StatusOK, InvoiceResponse{Invoice: inv})
	return
}

func (h *Handler) HandleStore(w http.ResponseWriter, r *http.Request) {
	var ir InvoiceRequest
	err := json.NewDecoder(r.Body).Decode(&ir)
	if err != nil {
		h.writeJson(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	fmt.Println(ir)
}

func (h *Handler) writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
