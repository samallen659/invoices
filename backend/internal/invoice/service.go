package invoice

import (
	"context"
	"fmt"
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
	inv := NewInvoice()

	if err := s.setInvoice(inv, invoiceRequest); err != nil {
		return nil, err
	}

	if err := s.repo.StoreInvoice(ctx, *inv); err != nil {
		return nil, fmt.Errorf("failed creating invoice: %w", err)
	}

	return inv, nil
}

func (s *Service) UpdateInvoice(ctx context.Context, id uuid.UUID, invoiceRequest InvoiceRequest) (*Invoice, error) {
	inv, err := s.repo.GetInvoiceByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := s.setInvoice(inv, invoiceRequest); err != nil {
		return nil, err
	}

	if err := s.repo.UpdateInvoice(ctx, inv); err != nil {
		return nil, fmt.Errorf("failed updating invoice: %w", err)
	}

	return inv, nil
}

func (s *Service) DeleteInvoice(ctx context.Context, id uuid.UUID) error {
	err := s.repo.DeleteInvoice(ctx, id)
	if err != nil {
		return fmt.Errorf("failed deleting invoice: %w", err)
	}

	return nil
}

func (s *Service) setInvoice(inv *Invoice, invoiceRequest InvoiceRequest) error {
	if err := inv.SetDescription(invoiceRequest.Description); err != nil {
		return fmt.Errorf("failed creating invoice: %w", err)
	}
	if err := inv.SetPaymentDue(invoiceRequest.PaymentDue); err != nil {
		return fmt.Errorf("failed creating invoice: %w", err)
	}
	if err := inv.SetPaymentTerms(invoiceRequest.PaymentTerms); err != nil {
		return fmt.Errorf("failed creating invoice: %w", err)
	}

	if err := inv.SetStatus(invoiceRequest.Status); err != nil {
		return fmt.Errorf("failed creating invoice: %w", err)
	}

	client, err := NewClient(invoiceRequest.ClientName, invoiceRequest.ClientEmail)
	if err != nil {
		return fmt.Errorf("failed creating invoice: %w", err)
	}
	if err := inv.SetClient(*client); err != nil {
		return fmt.Errorf("failed creating invoice: %w", err)
	}

	clientAddress, err := NewAddress(invoiceRequest.ClientAddress.Street,
		invoiceRequest.ClientAddress.City, invoiceRequest.ClientAddress.PostCode, invoiceRequest.ClientAddress.Country)
	if err != nil {
		return fmt.Errorf("failed creating invoice: %w", err)
	}
	if err := inv.SetClientAddress(*clientAddress); err != nil {
		return fmt.Errorf("failed creating invoice: %w", err)
	}

	senderAddress, err := NewAddress(invoiceRequest.SenderAddress.Street,
		invoiceRequest.SenderAddress.City, invoiceRequest.SenderAddress.PostCode, invoiceRequest.SenderAddress.Country)
	if err != nil {
		return fmt.Errorf("failed creating invoice: %w", err)
	}
	if err := inv.SetSenderAddress(*senderAddress); err != nil {
		return fmt.Errorf("failed creating invoice: %w", err)
	}

	invoiceItems := []InvoiceItem{}
	for _, item := range invoiceRequest.Items {
		invItem, err := NewInvoiceItem(item.Name, item.Price, item.Quantity)
		if err != nil {
			return fmt.Errorf("failed creating invoice: %w", err)
		}
		invoiceItems = append(invoiceItems, *invItem)
	}

	if err := inv.SetInvoiceItems(invoiceItems); err != nil {
		return fmt.Errorf("failed creating invoice: %w", err)
	}

	return nil
}
