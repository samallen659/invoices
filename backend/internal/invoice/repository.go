package invoice

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository interface {
	GetInvoiceByID(ctx context.Context, id uuid.UUID) (*Invoice, error)
	GetAllInvoices(ctx context.Context) ([]*Invoice, error)
}

type PostgresRepository struct {
	conn *sqlx.DB
}

func NewPostgresRespository(connectionURI string) (*PostgresRepository, error) {
	//TODO
	return nil, nil
}

func (pr *PostgresRepository) GetInvoiceByID(ctx context.Context, id uuid.UUID) (*Invoice, error) {
	//TODO
	return nil, nil
}

func (pr *PostgresRepository) GetAllInvoices(ctx context.Context) ([]*Invoice, error) {
	//TODO
	return nil, nil
}
