package invoice

import "github.com/Rhymond/go-money"

type InvoiceItem struct {
	Quantity int
	Item     Item
	Total    money.Money
}
