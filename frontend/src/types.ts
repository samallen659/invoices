type InvoiceStatus = "pending" | "paid" | "draft";

type Address = {
	Street: string;
	City: string;
	PostCode: string;
	Country: string;
};

type Item = {
	Name: string;
	Quantity: number;
	Price: number;
	Total: number;
};

type ItemReq = {
	Name: string;
	Price: number;
};

type InvoiceItem = {
	Item: ItemReq;
	Quantity: number;
	Total: number;
};

type Client = {
	ClientName: string;
	ClientEmail: string;
};

type Invoice = {
	ID: string;
	PaymentDue: Date;
	Description: string;
	PaymentTerms: number;
	ClientName: string;
	ClientEmail: string;
	Status: InvoiceStatus;
	ClientAddress: Address;
	SenderAddress: Address;
	Items: Item[];
	Total: number;
};

type InvoiceReq = {
	ID: string;
	CreatedAt: Date;
	PaymentDue: Date;
	Description: string;
	PaymentTerms: number;
	Client: Client;
	Status: InvoiceStatus;
	ClientAddress: Address;
	SenderAddress: Address;
	InvoiceItems: InvoiceItem[];
	Total: number;
};

export type { InvoiceStatus, Address, Item, InvoiceItem, Client, Invoice, InvoiceReq };
