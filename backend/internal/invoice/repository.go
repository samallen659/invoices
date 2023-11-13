package invoice

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository interface {
	GetInvoiceByID(ctx context.Context, id uuid.UUID) (*Invoice, error)
	GetAllInvoices(ctx context.Context) ([]*Invoice, error)
	StoreInvoice(ctx context.Context, invoice Invoice) error
}

type PostgresRepository struct {
	conn *sqlx.DB
}

func NewPostgresRespository(connectionURI string) (*PostgresRepository, error) {
	conn, err := sqlx.Open("postgres", connectionURI)
	if err != nil {
		return nil, fmt.Errorf("failed to open Postgres connection: %w", err)
	}
	return &PostgresRepository{conn: conn}, nil
}

func (pr *PostgresRepository) GetInvoiceByID(ctx context.Context, id uuid.UUID) (*Invoice, error) {
	//TODO
	return nil, nil
}

func (pr *PostgresRepository) GetAllInvoices(ctx context.Context) ([]*Invoice, error) {
	//TODO
	return nil, nil
}

func (pr *PostgresRepository) StoreInvoice(ctx context.Context, invoice Invoice) error {
	//TODO
	return nil
}
