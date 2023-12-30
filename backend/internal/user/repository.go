package user

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang.org/x/net/context"
)

type Repository interface {
	GetUser(ctx context.Context, id uuid.UUID) (*User, error)
	StoreUser(ctx context.Context, user User) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
	UpdateUser(ctx context.Context, user User) error
}

type PostgresRepository struct {
	conn *sqlx.DB
}

func NewPostgresRepository(conn *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{conn: conn}
}

func (p *PostgresRepository) GetUser(ctx context.Context, id uuid.UUID) (*User, error) {
	return nil, nil
}

func (p *PostgresRepository) StoreUser(ctx context.Context, user User) error {
	return nil
}

func (p *PostgresRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return nil
}

func (p *PostgresRepository) UpdateUser(ctx context.Context, user User) error {
	return nil
}
