package invoice

import "errors"

type Item struct {
	Name  string
	Price float32
}

type InvoiceItem struct {
	Quantity int
	Item     Item
	Total    float32
}

func InvoiceItemFactory(name string, price float32, quantity int) (*InvoiceItem, error) {
	if name == "" {
		return nil, errors.New("name cannot be emtpy")
	}
	if price <= 0.0 {
		return nil, errors.New("price must be greater than 0")
	}
	if quantity <= 0 {
		return nil, errors.New("quantity must be greater than 0")
	}

	//TODO
	return nil, nil
}
