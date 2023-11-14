package invoice

import (
	"context"
	"database/sql"
	"errors"
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

const GET_INVOICE_BY_ID_QUERY = `
SELECT 
    invoice.id, invoice.created_at, invoice.payment_due, invoice.description, invoice.payment_terms, invoice.status, invoice.total,
	client.client_name, client.client_email,
    sub1.street AS sender_street,
	sub1.city AS sender_city,
	sub1.post_code AS sender_post_code,
	sub1.country AS sender_country,
    sub2.street AS client_street,
	sub2.city AS client_city, 
	sub2.post_code AS client_post_code,
	sub2.country AS client_county
FROM 
    invoice
JOIN
	client ON invoice.client_id=client.id
JOIN 
    (SELECT id, street, city, post_code, country FROM address) sub1 ON sub1.id = invoice.sender_address_id
JOIN 
    (SELECT id, street, city, post_code, country FROM address) sub2 ON sub2.id = invoice.client_address_id
WHERE
	invoice.id=$1;
`

const GET_INVOICE_ITEMS = `
SELECT 
	item.name, item.price, invoice_item.quantity, invoice_item.total
FROM 
	invoice_item
JOIN
	item ON invoice_item.item_id=item.id
WHERE
	invoice_item.invoice_id=$1;
`

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

func (pr *PostgresRepository) GetInvoiceByID(ctx context.Context, id string) (*Invoice, error) {
	var invoice Invoice
	err := pr.conn.QueryRow(GET_INVOICE_BY_ID_QUERY, id).Scan(
		&invoice.ID,
		&invoice.CreatedAt,
		&invoice.PaymentDue,
		&invoice.Description,
		&invoice.PaymentTerms,
		&invoice.Status,
		&invoice.Total,
		&invoice.Client.ClientName,
		&invoice.Client.ClientEmail,
		&invoice.SenderAddress.Street,
		&invoice.SenderAddress.City,
		&invoice.SenderAddress.PostCode,
		&invoice.SenderAddress.Country,
		&invoice.ClientAddress.Street,
		&invoice.ClientAddress.City,
		&invoice.ClientAddress.PostCode,
		&invoice.ClientAddress.Country,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("no invoice found with the given ID: %s", id)
	}
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("failed to run query against database")
	}

	err = pr.getInvoiceItems(&invoice)
	if err != nil {
		return nil, err
	}

	return &invoice, nil
}

func (pr *PostgresRepository) getInvoiceItems(invoice *Invoice) error {
	rows, err := pr.conn.Query(GET_INVOICE_ITEMS, invoice.ID)
	if err != nil {
		fmt.Println(err)
		return errors.New("failed to run query to get invoice items")
	}

	for rows.Next() {
		var invoiceItem InvoiceItem

		err := rows.Scan(
			&invoiceItem.Item.Name,
			&invoiceItem.Item.Price,
			&invoiceItem.Quantity,
			&invoiceItem.Total,
		)

		if err != nil {
			return fmt.Errorf("failed scanning rows into invoiceItem: %w", err)
		}

		invoice.InvoiceItems = append(invoice.InvoiceItems, invoiceItem)
	}

	return nil
}

func (pr *PostgresRepository) GetAllInvoices(ctx context.Context) ([]*Invoice, error) {
	//TODO
	return nil, nil
}

func (pr *PostgresRepository) StoreInvoice(ctx context.Context, invoice Invoice) error {
	//TODO
	return nil
}
