package invoice

import "github.com/Rhymond/go-money"

type Invoice_item struct {
	Quantity int
	Item     Item
	Total    money.Money
}
