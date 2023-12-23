package invoice

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Item struct {
	Name  string
	Price float64
}

type InvoiceItem struct {
	Quantity int
	Item     Item
	Total    float64
}

var MONEY_SCAN_ERROR = errors.New("Failed to scan into Money")

func NewInvoiceItem(name string, price float64, quantity int) (*InvoiceItem, error) {
	if name == "" {
		return nil, errors.New("name cannot be emtpy")
	}
	if price <= 0.0 {
		return nil, errors.New("price must be greater than 0")
	}
	if quantity <= 0 {
		return nil, errors.New("quantity must be greater than 0")
	}

	total := price * float64(quantity)
	invoiceItem := &InvoiceItem{
		Item: Item{
			Name:  name,
			Price: price,
		},
		Quantity: quantity,
		Total:    total,
	}
	return invoiceItem, nil
}

type Money struct {
	pounds int
	pence  int
}

// Implements sql/driver Valuer interface
func (m *Money) Value() (driver.Value, error) {
	return fmt.Sprintf("%d.%d", m.pounds, m.pence), nil
}

// Implements sql Scanner interface
func (m *Money) Scan(value any) error {
	if value == nil {
		m.pounds = 0
		m.pence = 0
		return nil
	}

	if sv, err := driver.String.ConvertValue(value); err == nil {
		if v, ok := sv.(string); ok {
			data := strings.Split(v, ".")
			if len(data) < 2 || len(data) > 2 {

				return MONEY_SCAN_ERROR
			}
			pounds, err := strconv.Atoi(data[0])
			if err != nil {
				return MONEY_SCAN_ERROR
			}
			pence, err := strconv.Atoi(data[1])
			if err != nil {
				return MONEY_SCAN_ERROR
			}
			m.pounds = pounds
			m.pence = pence
			return nil
		}
	}
	return MONEY_SCAN_ERROR
}
