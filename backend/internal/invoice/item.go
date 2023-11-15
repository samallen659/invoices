package invoice

type Item struct {
	Name  string
	Price float32
}

type InvoiceItem struct {
	Quantity int
	Item     Item
	Total    float32
}
