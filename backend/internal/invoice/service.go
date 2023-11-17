package invoice

import (
	"context"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) (*Service, error) {
	//TODO: checks
	return &Service{repo: repo}, nil
}

func (s *Service) GetAll(ctx context.Context) ([]*Invoice, error) {
	invoices, err := s.repo.GetAllInvoices(ctx)
	if err != nil {
		return nil, err
	}

	return invoices, nil
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*Invoice, error) {
	invoice, err := s.repo.GetInvoiceByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return invoice, nil
}

func (s *Service) NewInvoice(ctx context.Context, invoiceRequest InvoiceRequest) (*Invoice, error) {
	//TODO
	return nil, nil
}
