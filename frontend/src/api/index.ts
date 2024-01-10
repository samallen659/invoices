import { Invoice, InvoiceItem, InvoiceRes, Item } from "../types";

const ENV = import.meta.env.VITE_ENVIRONMENT;
const apiBaseURL = ENV === "PROD" ? import.meta.env.VITE_API_BASE_URL : "";

async function getAllInvoices(): Promise<Invoice[]> {
	const response = await fetch(`${apiBaseURL}/invoice`, {
		method: "GET",
		headers: {
			"Content-Type": "application/json",
		},
	});

	if (!response.ok) {
		const message = `An error has occured: ${response.status}`;
		throw new Error(message);
	}

	const json = await response.json();
	const invReqs: InvoiceRes[] = json?.invoice;
	const invoices: Invoice[] = invReqs.map((iv: InvoiceRes) => invoiceResToInvoice(iv));
	return invoices;
}

async function deleteInvoice(id: string) {
	const response = await fetch(`${apiBaseURL}/invoice/${id}`, {
		method: "DELETE",
	});

	if (!response.ok) {
		const message = `An error has occured: ${response.status}`;
		throw new Error(message);
	}
}

async function newInvoice(i: Invoice) {
	const response = await fetch(`${apiBaseURL}/invoice`, {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify(i),
	});

	if (!response.ok) {
		const message = `An error has occured: ${response.status}`;
		throw new Error(message);
	}
}

async function editInvoice(i: Invoice) {
	const response = await fetch(`${apiBaseURL}/invoice/${i.ID}`, {
		method: "PUT",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify(i),
	});

	if (!response.ok) {
		const message = `An error has occured: ${response.status}`;
		throw new Error(message);
	}
}

function invoiceResToInvoice(iq: InvoiceRes): Invoice {
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
		PaymentDue: iq.PaymentDue.toString().split("T")[0], //required for loading into date field in form
		CreatedAt: iq.CreatedAt.toString(),
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

export { getAllInvoices, deleteInvoice, newInvoice, editInvoice };
