import { Invoice, InvoiceItem, InvoiceReq, InvoiceRes, Item } from "../types";

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
	const invReqs: InvoiceRes[] = json?.invoice;
	const invoices: Invoice[] = invReqs.map((iv: InvoiceRes) => invoiceResToInvoice(iv));
	return invoices;
}

async function deleteInvoice(id: string) {
	const response = await fetch(`/invoice/${id}`, {
		method: "DELETE",
		headers: {
			"Access-Control-Allow-Origin": "http://localhost:8080/invoice",
		},
	});

	console.log(response);

	if (!response.ok) {
		const message = `An error has occured: ${response.status}`;
		throw new Error(message);
	}
}

async function newInvoice(i: Invoice) {
	const response = await fetch("/invoice", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
			"Access-Control-Allow-Origin": "http://localhost:8080/invoice",
		},
		body: JSON.stringify(i),
	});

	if (!response.ok) {
		const message = `An error has occured: ${response.status}`;
		console.log(message);
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

export { getAllInvoices, deleteInvoice, newInvoice };
