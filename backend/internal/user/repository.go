package user

import (
	"database/sql"
	"errors"
	"fmt"

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
	var user User
	err := p.conn.QueryRowx(`SELECT id, first_name, last_name, email FROM users WHERE id=$1`,
		id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	if err == sql.ErrNoRows {
		return nil, errors.New("No User found with supplied id")
	}
	if err != nil {
		return nil, errors.New("failed to run query against database")
	}

	return &user, nil
}

func (p *PostgresRepository) StoreUser(ctx context.Context, user User) error {
	_, err := p.conn.Queryx(`INSERT INTO users(id, first_name, last_name, email) VALUES ($1, $2, $3, $4)`,
		user.ID, user.FirstName, user.LastName, user.Email)
	if err != nil {
		return errors.New("failed to insert into users")
	}
	return nil
}

func (p *PostgresRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := p.conn.Queryx(`DELETE FROM users WHERE id=$1`, id)
	if err != nil {
		return errors.New("failed to delete from users")
	}
	return nil
}

func (p *PostgresRepository) UpdateUser(ctx context.Context, user User) error {
	_, err := p.conn.Queryx(`UPDATE users SET first_name=$2, last_name=$3, email=$4 WHERE id=$1`,
		user.ID, user.FirstName, user.LastName, user.Email)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("failed to update users")
	}
	return nil
}
