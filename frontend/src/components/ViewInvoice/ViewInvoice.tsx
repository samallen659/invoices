import { InvoiceIdx } from "../../App";
import { Invoice, InvoiceItem } from "../../types";
import { getShortDate, getShortID } from "../../utils";
import { InvoiceStatusBox } from "../InvoiceStatusBox/InvoiceStatusBox";

type ViewInvoiceProps = {
	invoice: Invoice;
    toggle: (f: boolean) => void;
};

function ViewInvoice({ invoice, toggle }: ViewInvoiceProps) {
	return (
		<div className=" flex flex-col gap-6 overflow-auto">
			<button className="flex w-24 items-center gap-5" onClick={() => toggle(false)}>
				<IconArrowLeft />
				<span className="mt-1 font-bold dark:text-white">Go Back</span>
			</button>
			<div className="hidden md:block">
				<ViewInvoiceBar status={invoice.Status} />
			</div>
			<ViewInvoiceDetails invoice={invoice} />
		</div>
	);
}

export type ViewInvoiceBarProps = {
	status: string;
};
export function ViewInvoiceBar({ status }: ViewInvoiceBarProps) {
	return (
		<div className="fixed bottom-0 left-0 flex h-24 w-screen items-center justify-center bg-white px-8 dark:bg-indigo-800 md:relative md:h-[88px] md:w-full md:justify-between md:rounded-md">
			<div className="hidden items-center gap-5 md:flex">
				<p className="text-sm text-[#858BB2]">Status</p>
				<InvoiceStatusBox status={status} />
			</div>
			<div className="flex justify-center gap-2">
				<button className="h-12 w-[73px] rounded-full bg-[#F9FAFE] text-indigo-200 dark:bg-gray-600 dark:text-gray-200">
					Edit
				</button>
				<button className="h-12 w-[89px] rounded-full bg-red-400 text-white">Delete</button>
				<button className="h-12 w-[149px] rounded-full bg-purple-400 text-white md:w-[131px]">
					Mark as Paid
				</button>
			</div>
		</div>
	);
}

type ViewInvoiceDetailsProps = {
	invoice: Invoice;
};

function ViewInvoiceDetails({ invoice }: ViewInvoiceDetailsProps) {
	return (
		<div className="rounded-md bg-white p-6 dark:bg-indigo-800">
			<div className="flex flex-col gap-7 md:flex-row md:justify-between">
				<div>
					<h2 className="font-bold dark:text-white">
						<span className="text-indigo-200">#</span>
						{/*TODO: Way to see full ID*/}
						{getShortID(invoice.ID)}
					</h2>
				</div>
				<p className="text-sm text-indigo-200 dark:text-gray-200 md:text-end">
					{invoice.ClientAddress.Street}
					<br />
					{invoice.ClientAddress.City}
					<br />
					{invoice.ClientAddress.PostCode}
					<br />
					{invoice.ClientAddress.Country}
				</p>
			</div>
			<div className="grid grid-cols-2 grid-rows-3 gap-2 md:grid-cols-3 md:grid-rows-2">
				<div>
					<p className="pb-2 pt-4 text-sm text-indigo-200 dark:text-gray-200">Invoice Date</p>
					<p className="dark:text-white">{getShortDate(invoice.CreatedAt)}</p>
				</div>
				<div className="row-start-2">
					<p className="pb-2 pt-4 text-sm text-indigo-200 dark:text-gray-200">PaymentDue</p>
					<p className="dark:text-white">{getShortDate(invoice.PaymentDue)}</p>
				</div>
				<div className="row-start-3 md:col-start-3 md:row-start-1">
					<p className="pb-2 pt-4 text-sm text-indigo-200 dark:text-gray-200">Sent to</p>
					<p className="dark:text-white">{invoice.Client.ClientEmail}</p>
				</div>
				<div className="col-start-2 row-span-2">
					<p className="pb-2 pt-4 text-sm text-indigo-200 dark:text-gray-200">Bill To</p>
					<p className="dark:text-white">{invoice.Client.ClientName}</p>
					<p className="text-sm text-indigo-200 dark:text-gray-200">
						{invoice.SenderAddress.Street}
						<br />
						{invoice.SenderAddress.City}
						<br />
						{invoice.SenderAddress.PostCode}
						<br />
						{invoice.SenderAddress.Country}
					</p>
				</div>
			</div>
			<div className="mt-8 rounded-md bg-off-white dark:bg-gray-600">
				<div className="md:grid-col-5 grid gap-4 p-6">
					<h4 className="col-span-2 col-start-1 text-indigo-200 dark:text-gray-200">Item Name</h4>
					<h4 className="col-start-3 text-indigo-200 dark:text-gray-200">QTY.</h4>
					<h4 className="col-start-4 text-indigo-200 dark:text-gray-200">Price</h4>
					<h4 className="col-start-5 text-indigo-200 dark:text-gray-200">Total</h4>
					{invoice.InvoiceItems.map((invItem: InvoiceItem) => (
						<>
							<p className="col-span-2 dark:text-white">{invItem.Item.Name}</p>
							<p className="dark:text-white">{invItem.Quantity}</p>
							<p className="dark:text-white">{`£${invItem.Item.Price}`}</p>
							<p className="dark:text-white">{`£${invItem.Total}`}</p>
						</>
					))}
				</div>
				<div className="flex items-center justify-between rounded-b-md bg-[#373B53] p-6 dark:bg-gray-800">
					<p className="hidden text-sm text-white md:block">Amount Due</p>
					<p className=" text-sm text-white md:hidden">Grand Total</p>
					<h3 className="text-2xl font-bold text-white">{`£${invoice.Total}`}</h3>
				</div>
			</div>
		</div>
	);
}

function IconArrowLeft() {
	return (
		<svg width="7" height="10" xmlns="http://www.w3.org/2000/svg">
			<path
				d="M6.342.886L2.114 5.114l4.228 4.228"
				stroke="#9277FF"
				strokeWidth="2"
				fill="none"
				fillRule="evenodd"
			/>
		</svg>
	);
}
export { ViewInvoice };
