import { Invoice } from "../../types";

type InvoiceListProps = {
	invoices: Invoice[];
};

function InvoiceList({ invoices }: InvoiceListProps) {
	return (
		<ul className="flex flex-col gap-4">
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
	return (
		<li className="grid h-36 grid-cols-2 grid-rows-4 rounded-lg border-2 border-white border-opacity-0 bg-white p-6 hover:border-purple-400 dark:bg-indigo-800 dark:text-white md:h-20">
			<p className="font-bold dark:text-white">
				<span className="text-indigo-200">#</span>
				{getShortID(invoice.ID)}
			</p>
			<p className="text-indigo-200 dark:text-white">{invoice.Client.ClientName}</p>
			<p className="row-start-3 -mt-1 text-indigo-200 dark:text-gray-200">{getShortDate(invoice.PaymentDue)}</p>
			<p className="row-start-4 pt-1 font-bold">{`£${invoice.Total}`}</p>
			<div
				className={`col-start-2 row-span-2 row-start-3 rounded-md bg-opacity-10 font-bold ${getStatusColor(
					invoice.Status,
				)} flex items-center justify-center capitalize`}
			>
				{invoice.Status}
			</div>
		</li>
	);
}

function getShortID(id: string): string {
	let short = id.split("-")[0].toUpperCase();
	return `${short}...`;
}

function getShortDate(date: Date): string {
	const d = new Date(date);
	return `Due ${d.getDate()} ${d.getMonth()} ${d.getFullYear()}`;
}

function getStatusColor(status: string): string {
	let color: string;
	switch (status) {
		case "pending":
			color = "bg-[#FF8F00] text-[#FF8F00]";
			break;
		case "paid":
			color = "bg-[#33D69F] text-[#33D69F]";
			break;
		default:
			color = "bg-[#979797] text-[#979797]";
			break;
	}

	return color;
}

export { InvoiceList, InvoiceListItem };
