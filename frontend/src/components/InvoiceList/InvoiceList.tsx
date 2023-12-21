import { Invoice } from "../../types";
import { InvoiceIdx } from "../../App";
import { InvoiceStatusBox } from "../InvoiceStatusBox/InvoiceStatusBox";
import { getShortID, getShortDate } from "../../utils";
import { IconArrowRight } from "../Icons";

type InvoiceListProps = {
	invoices: Invoice[];
	handleInvoiceIdxChange: (i: InvoiceIdx) => void;
	filter: string;
};

function InvoiceList({ invoices, handleInvoiceIdxChange, filter }: InvoiceListProps) {
	return (
		<ul className="flex flex-col gap-4">
			{invoices.map((invoice, i) => {
				if (filter === "" || filter === invoice.Status) {
					return (
						<InvoiceListItem
							key={i}
							invoice={invoice}
							idx={i}
							handleInvoiceIdxChange={handleInvoiceIdxChange}
						/>
					);
				}
			})}
		</ul>
	);
}

type InvoiceListItemProps = {
	idx: number;
	invoice: Invoice;
	handleInvoiceIdxChange: (i: InvoiceIdx) => void;
};

function InvoiceListItem({ idx, invoice, handleInvoiceIdxChange }: InvoiceListItemProps) {
	return (
		<li
			className="grid h-36 cursor-pointer grid-cols-2 grid-rows-4 justify-evenly rounded-lg border-2 border-white border-opacity-0 bg-white 
            p-6 hover:border-purple-400 dark:bg-indigo-800 dark:text-white md:flex md:h-20 md:items-center"
			onClick={() => handleInvoiceIdxChange(idx)}
		>
			<p className="md:row-col-1 font-bold dark:text-white md:basis-1/5 ">
				<span className="text-indigo-200">#</span>
				{getShortID(invoice.ID)}
			</p>
			<p className="justify-self-end text-indigo-200 dark:text-white md:col-start-3 md:basis-1/5  md:justify-self-start">
				{invoice.ClientName}
			</p>
			<p className="row-start-3 -mt-1 text-indigo-200 dark:text-gray-200 md:col-start-2 md:mt-0 md:basis-1/5 ">
				{`Due ${getShortDate(invoice.PaymentDue)}`}
			</p>
			<p className="row-start-4 pt-1 font-bold md:col-start-4 md:basis-1/5 md:text-center xl:text-end">{`£${invoice.Total}`}</p>
			<div className="col-start-2 row-span-2 row-start-3 flex items-center gap-5 justify-self-end md:col-start-5 md:basis-1/5 md:justify-end md:justify-self-end">
				<InvoiceStatusBox status={invoice.Status} />
				<div className="hidden md:block">
					<IconArrowRight />
				</div>
			</div>
		</li>
	);
}

export { InvoiceList, InvoiceListItem };
