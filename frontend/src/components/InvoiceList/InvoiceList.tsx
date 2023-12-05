import { Invoice } from "../../types";

type InvoiceListProps = {
	invoices: Invoice[];
};

function InvoiceList({ invoices }: InvoiceListProps) {
	return (
		<ul>
			{invoices.map((invoice, i) => (
				<InvoiceListItem key={i} invoice={invoice} />
			))}
		</ul>
	);
}

type InvoiceListItemProps = {
	invoice: Invoice;
};

function InvoiceListItem({ invoice }: InvoiceListItemProps) {
	return <li>{invoice.Description}</li>;
}

export { InvoiceList, InvoiceListItem };
