import { Invoice } from "../types";

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

	const invoices: Invoice[] = await response.json();
	return [...invoices];
}

export { getAllInvoices };