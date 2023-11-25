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
	UpdateInvoice(ctx context.Context, invoice Invoice) error
}

const (
	GET_ALL_INVOICE_QUERY = `
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
    `
	GET_INVOICE_BY_ID_QUERY = `
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
	GET_INVOICE_ITEMS = `
    SELECT 
    	item.name, item.price, invoice_item.quantity, invoice_item.total
    FROM 
	    invoice_item
    JOIN
	    item ON invoice_item.item_id=item.id
    WHERE
	    invoice_item.invoice_id=$1;
    `
)

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
	var invoice Invoice
	err := pr.conn.QueryRowx(GET_INVOICE_BY_ID_QUERY, id).Scan(
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
		return nil, errors.New("failed to run query against database")
	}

	err = pr.getInvoiceItems(&invoice)
	if err != nil {
		return nil, err
	}

	return &invoice, nil
}

func (pr *PostgresRepository) GetAllInvoices(ctx context.Context) ([]*Invoice, error) {
	rows, err := pr.conn.Queryx(GET_ALL_INVOICE_QUERY)
	if err != nil {
		return nil, errors.New("failed to run query against database")
	}
	defer rows.Close()

	var invoices []*Invoice
	for rows.Next() {
		var invoice Invoice

		err := rows.Scan(
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
		if err != nil {
			return nil, errors.New("failed to run query against database")
		}

		err = pr.getInvoiceItems(&invoice)
		if err != nil {
			return nil, err
		}

		invoices = append(invoices, &invoice)

	}

	return invoices, nil
}

func (pr *PostgresRepository) getInvoiceItems(invoice *Invoice) error {
	rows, err := pr.conn.Queryx(GET_INVOICE_ITEMS, invoice.ID)
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

func (pr *PostgresRepository) StoreInvoice(ctx context.Context, invoice Invoice) error {
	tx, err := pr.conn.Beginx()
	if err != nil {
		return errors.New("failed to start sql transaction")
	}
	defer tx.Rollback()

	var clientID int
	err = tx.QueryRowx("INSERT INTO client(client_name, client_email) VALUES ($1, $2) RETURNING id",
		invoice.Client.ClientName, invoice.Client.ClientEmail).Scan(&clientID)
	if err != nil {
		return errors.New("failed on INSERT INTO client")
	}

	var senderAddressID int
	err = tx.QueryRowx("INSERT INTO address(street, city, post_code, country) VALUES ($1, $2, $3, $4) RETURNING id",
		invoice.SenderAddress.Street, invoice.SenderAddress.City, invoice.SenderAddress.PostCode, invoice.SenderAddress.Country).Scan(&senderAddressID)
	if err != nil {
		return errors.New("failed on INSERT INTO address for SenderAddress")
	}

	var clientAddressID int
	err = tx.QueryRowx("INSERT INTO address(street, city, post_code, country) VALUES ($1, $2, $3, $4) RETURNING id",
		invoice.ClientAddress.Street, invoice.ClientAddress.City, invoice.ClientAddress.PostCode, invoice.ClientAddress.Country).Scan(&clientAddressID)
	if err != nil {
		return errors.New("failed on INSERT INTO address for ClientAddress")
	}

	_, err = tx.Exec(`INSERT INTO invoice(id, created_at, payment_due, description, client_id, payment_terms, status, total, sender_address_id, client_address_id)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`, invoice.ID, invoice.CreatedAt, invoice.PaymentDue, invoice.Description, clientID, invoice.PaymentTerms,
		invoice.Status, invoice.Total, senderAddressID, clientAddressID)
	if err != nil {
		return errors.New("failed on INSERT INTO invoice")
	}

	err = pr.insertInvoiceItems(tx, invoice.ID, invoice.InvoiceItems)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return errors.New("failed Commiting transaction to database")
	}

	return nil
}

func (pr *PostgresRepository) UpdateInvoice(ctx context.Context, invoice Invoice) error {
	tx, err := pr.conn.Beginx()
	if err != nil {
		return errors.New("failed to start sql transaction")
	}
	defer tx.Rollback()

	var clientID string
	var senderAddressID string
	var clientAddressID string
	err = tx.QueryRowx(`UPDATE invoice SET payment_due=$1, description=$2, payment_terms=$3, status=$4, total=$5 
        WHERE id=$6 RETURNING client_id, sender_address_id, client_address_id`, invoice.PaymentDue, invoice.Description,
		invoice.PaymentTerms, invoice.Status, invoice.Total, invoice.ID).Scan(&clientID, &senderAddressID, &clientAddressID)
	if err != nil {
		return errors.New("failed on UPDATE to invoice")
	}

	_, err = tx.Exec(`UPDATE client SET client_name=$1, client_email=$2 WHERE client_id=$3`, invoice.Client.ClientName, invoice.Client.ClientEmail, clientID)
	if err != nil {
		return errors.New("failed on UPDATE to client")
	}

	_, err = tx.Exec(`UPDATE address SET street=$1, city=$2, post_code=$3, country=$4 WHERE address_id=$5`, invoice.SenderAddress.Street,
		invoice.SenderAddress.City, invoice.SenderAddress.PostCode, invoice.SenderAddress.Country, senderAddressID)
	if err != nil {
		return errors.New("failed on UPDATE to senderAddress")
	}

	_, err = tx.Exec(`UPDATE address SET street=$1, city=$2, post_code=$3, country=$4 WHERE address_id=$5`, invoice.ClientAddress.Street,
		invoice.ClientAddress.City, invoice.ClientAddress.PostCode, invoice.ClientAddress.Country, clientAddressID)
	if err != nil {
		return errors.New("failed on UPDATE to clientAddress")
	}

	rows, err := tx.Queryx(`DELETE FROM invoice_item WHERE invoice_id=$1 RETURNING item_id`, invoice.ID)
	if err != nil {
		return errors.New("failed on DELETE to old invoiceItems")
	}
	defer rows.Close()

	var itemIDs []int
	for rows.Next() {
		var itemID int

		err := rows.Scan(&itemID)
		if err != nil {
			return errors.New("failed scanning itemID")
		}

		itemIDs = append(itemIDs, itemID)
	}

	for _, id := range itemIDs {
		_, err = tx.Exec(`DELETE FROM item WHERE id=$1`, id)
		if err != nil {
			return errors.New("failed on DELETE to old item")
		}
	}

	err = pr.insertInvoiceItems(tx, invoice.ID, invoice.InvoiceItems)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return errors.New("failed Commiting transaction to database")
	}

	return nil
}

func (pr *PostgresRepository) insertInvoiceItems(tx *sqlx.Tx, invoiceID uuid.UUID, invoiceItems []InvoiceItem) error {
	for _, iItem := range invoiceItems {
		var itemID int
		err := tx.QueryRowx("INSERT INTO item(name, price) VALUES ($1, $2) RETURNING id", iItem.Item.Name, iItem.Item.Price).Scan(&itemID)
		if err != nil {
			return errors.New("failed on INSERT INTO item")
		}

		_, err = tx.Exec("INSERT INTO invoice_item(invoice_id, item_id, quantity, total) VALUES ($1, $2, $3, $4)",
			invoiceID, itemID, iItem.Quantity, iItem.Total)
		if err != nil {
			return errors.New("failed on INSERT INTO invoice_item")
		}
	}

	return nil
}
