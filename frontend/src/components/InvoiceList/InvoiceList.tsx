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
		<li
			className="grid h-36 cursor-pointer grid-cols-2 grid-rows-4 justify-evenly rounded-lg border-2 border-white border-opacity-0 bg-white 
            p-6 hover:border-purple-400 dark:bg-indigo-800 dark:text-white md:flex md:h-20 md:items-center"
		>
			<p className="md:row-col-1 font-bold dark:text-white md:basis-1/5 ">
				<span className="text-indigo-200">#</span>
				{getShortID(invoice.ID)}
			</p>
			<p className="justify-self-end text-indigo-200 dark:text-white md:col-start-3 md:basis-1/5  md:justify-self-start">
				{invoice.Client.ClientName}
			</p>
			<p className="row-start-3 -mt-1 text-indigo-200 dark:text-gray-200 md:col-start-2 md:mt-0 md:basis-1/5 ">
				{getShortDate(invoice.PaymentDue)}
			</p>
			<p className="row-start-4 pt-1 font-bold md:col-start-4 md:basis-1/5 md:text-center xl:text-end">{`£${invoice.Total}`}</p>
			<div className="col-start-2 row-span-2 row-start-3 flex items-center gap-5 justify-self-end md:col-start-5 md:basis-1/5 md:justify-end md:justify-self-end">
				<div
					className={`h-10 w-[104px] rounded-md bg-opacity-10 font-bold ${getStatusColor(
						invoice.Status,
					)} flex items-center justify-center capitalize`}
				>
					{invoice.Status}
				</div>
				<div className="hidden md:block">
					<IconArrowRight />
				</div>
			</div>
		</li>
	);
}

function IconArrowRight() {
	return (
		<svg width="7" height="10" xmlns="http://www.w3.org/2000/svg">
			<path d="M1 1l4 4-4 4" stroke="#7C5DFA" strokeWidth="2" fill="none" fillRule="evenodd" />
		</svg>
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
