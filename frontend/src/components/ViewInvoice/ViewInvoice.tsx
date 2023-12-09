import { InvoiceIdx } from "../../App";
import { Invoice } from "../../types";
import { InvoiceStatusBox } from "../InvoiceStatusBox/InvoiceStatusBox";

type ViewInvoiceProps = {
	invoice: Invoice;
	handleInvoiceIdxChange: (i: InvoiceIdx) => void;
};

function ViewInvoice({ invoice, handleInvoiceIdxChange }: ViewInvoiceProps) {
	return (
		<div className="flex flex-col gap-6">
			<button className="flex w-24 items-center gap-5" onClick={() => handleInvoiceIdxChange(0)}>
				<IconArrowLeft />
				<span className="mt-1 font-bold">Go Back</span>
			</button>
			<ViewInvoiceBar status={invoice.Status} />
			<ViewInvoiceDetails invoice={invoice} />
		</div>
	);
}

type ViewInvoiceBarProps = {
	status: string;
};
function ViewInvoiceBar({ status }: ViewInvoiceBarProps) {
	return (
		<div className="absolute bottom-0 left-0 flex h-24 w-screen items-center justify-center bg-white px-8 dark:bg-indigo-800 md:relative md:h-[88px] md:w-full md:justify-between md:rounded-md">
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
	return <div className="rounded-md bg-white dark:bg-indigo-800"></div>;
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
