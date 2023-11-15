package invoicing

type InvoiceItem struct {
	Quantity int
	Item     Item
	Total    float32
}
