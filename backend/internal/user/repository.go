package user

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang.org/x/net/context"
)

type Repository interface {
	GetUser(ctx context.Context, email string, userName string) (*User, error)
	StoreUser(ctx context.Context, user User) error
	DeleteUser(ctx context.Context, userName string) error
	UpdateUser(ctx context.Context, user User) error
}

type PostgresRepository struct {
	conn *sqlx.DB
}

func NewPostgresRepository(conn *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{conn: conn}
}

func (p *PostgresRepository) GetUser(ctx context.Context, email string, userName string) (*User, error) {
	var user User
	var err error
	if userName != "" {
		err = p.conn.QueryRowx(`SELECT user_name, first_name, last_name, email, password FROM users WHERE user_name=$1`,
			userName).Scan(&user.UserName, &user.FirstName, &user.LastName, &user.Email, &user.Password)
		if err == sql.ErrNoRows {
			return nil, errors.New("No User found with supplied username")
		}
	} else {
		err = p.conn.QueryRowx(`SELECT user_name, first_name, last_name, email, password FROM users WHERE email=$1`,
			email).Scan(&user.UserName, &user.FirstName, &user.LastName, &user.Email, &user.Password)
		if err == sql.ErrNoRows {
			return nil, errors.New("No User found with supplied email")
		}
	}
	if err != nil {
		return nil, errors.New("failed to run query against database")
	}

	return &user, nil
}

func (p *PostgresRepository) StoreUser(ctx context.Context, user User) error {
	_, err := p.conn.Queryx(`INSERT INTO users(user_name, first_name, last_name, email, password) VALUES ($1, $2, $3, $4, $5)`,
		user.UserName, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return errors.New("failed to insert into users")
	}
	return nil
}

func (p *PostgresRepository) DeleteUser(ctx context.Context, userName string) error {
	_, err := p.conn.Queryx(`DELETE FROM users WHERE user_name=$1`, userName)
	if err != nil {
		return errors.New("failed to delete from users")
	}
	return nil
}

func (p *PostgresRepository) UpdateUser(ctx context.Context, user User) error {
	_, err := p.conn.Queryx(`UPDATE users SET first_name=$2, last_name=$3, email=$4, password=$6 WHERE user_name=$1`,
		user.UserName, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("failed to update users")
	}
	return nil
}
