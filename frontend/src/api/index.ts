import { Invoice, InvoiceItem, InvoiceReq, Item, ItemReq } from "../types";

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
	const invReq = invoiceToInvoiceReq(i);
	const response = await fetch("/invoice", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
			"Access-Control-Allow-Origin": "http://localhost:8080/invoice",
		},
		body: JSON.stringify(invReq),
	});

	if (!response.ok) {
		const message = `An error has occured: ${response.status}`;
		throw new Error(message);
	}
}

function invoiceToInvoiceReq(i: Invoice): InvoiceReq {
	const invItems: InvoiceItem[] = i.Items.map((it: Item) => {
		return {
			Item: { Name: it.Name, Price: it.Price } as ItemReq,
			Quantity: it.Quantity,
			Total: it.Total,
		};
	});

	const iq: InvoiceReq = {
		ID: i.ID,
		PaymentDue: new Date(i.PaymentDue),
		CreatedAt: new Date(i.CreatedAt),
		Description: i.Description,
		PaymentTerms: i.PaymentTerms,
		Client: { ClientName: i.ClientName, ClientEmail: i.ClientEmail },
		Status: i.Status,
		ClientAddress: i.ClientAddress,
		SenderAddress: i.SenderAddress,
		InvoiceItems: invItems,
		Total: i.Total,
	};

	return iq;
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
