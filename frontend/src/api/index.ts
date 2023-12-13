import { Invoice, InvoiceItem, InvoiceReq, Item } from "../types";

async function getAllInvoices(): Promise<Invoice[]> {
	const response = await fetch("/invoice", {
		method: "GET",
		headers: {
			"Content-Type": "application/json",
			"Access-Control-Allow-Origin": "http://localhost:8080/invoice",
		},
	});

	if (!response.ok) {
		const message = `An error has occured: ${response.status}`;
		throw new Error(message);
	}

	const json = await response.json();
	const invReqs: InvoiceReq[] = json?.invoice;
	const invoices: Invoice[] = invReqs.map((iv: InvoiceReq) => invoiceReqToInvoice(iv));
	return invoices;
}

function invoiceReqToInvoice(iq: InvoiceReq): Invoice {
	const items: Item[] = iq.InvoiceItems.map((iv: InvoiceItem) => {
		return {
			Name: iv.Item.Name,
			Quantity: iv.Quantity,
			Price: iv.Item.Price,
			Total: iv.Total,
		} as Item;
	});
	const i: Invoice = {
		ID: iq.ID,
		CreatedAt: iq.CreatedAt,
		PaymentDue: iq.PaymentDue,
		Description: iq.Description,
		PaymentTerms: iq.PaymentTerms,
		ClientName: iq.Client.ClientName,
		ClientEmail: iq.Client.ClientEmail,
		Status: iq.Status,
		ClientAddress: {
			Street: iq.ClientAddress.Street,
			PostCode: iq.ClientAddress.PostCode,
			City: iq.ClientAddress.City,
			Country: iq.ClientAddress.Country,
		},
		SenderAddress: {
			Street: iq.SenderAddress.Street,
			PostCode: iq.SenderAddress.PostCode,
			City: iq.SenderAddress.City,
			Country: iq.SenderAddress.Country,
		},
		Items: items,
		Total: iq.Total,
	};

	return i;
}

export { getAllInvoices };
