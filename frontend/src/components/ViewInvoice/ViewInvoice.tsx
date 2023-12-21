import { Fragment } from "react";
import { Invoice, Item } from "../../types";
import { getShortDate, getShortID } from "../../utils";
import { InvoiceStatusBox } from "../InvoiceStatusBox/InvoiceStatusBox";
import { useTransition } from "react-transition-state";
import { IconLeftArrow, IconSpinning } from "../Icons";
import { useMutation, useQueryClient } from "react-query";
import { deleteInvoice, editInvoice } from "../../api";

type ViewInvoiceProps = {
	invoice: Invoice;
	viewToggle: (f: boolean) => void;
	formToggle: (f: boolean) => void;
	setFormState: (v: "new" | "edit") => void;
};

function ViewInvoice({ invoice, viewToggle, formToggle, setFormState }: ViewInvoiceProps) {
	const handleEditClick = () => {
		setFormState("edit");
		formToggle(true);
		viewToggle(false);
	};

	const [{ status: deleteStatus, isMounted: isDeleting }, deleteToggle] = useTransition({
		timeout: 500,
		mountOnEnter: true,
		unmountOnExit: true,
		preEnter: true,
	});

	const queryClient = useQueryClient();
	const deleteMutation = useMutation(deleteInvoice, {
		onSuccess: () => {
			viewToggle(false);
			setTimeout(() => queryClient.invalidateQueries("invoices"), 750);
		},
	});
	const paidMutation = useMutation(editInvoice, {
		onSuccess: () => {
			queryClient.invalidateQueries("invoices");
		},
	});

	const handleMarkAsPaid = () => {
		const i = { ...invoice };
		i.Status = "paid";
		i.PaymentDue = `${i.PaymentDue}T00:00:00Z`;
		paidMutation.mutate(i);
	};

	return (
		<div className=" mb-20 flex flex-col gap-6 overflow-auto md:mb-6">
			<button className="flex w-24 items-center gap-5" onClick={() => viewToggle(false)}>
				<IconLeftArrow />
				<span className="mt-1 font-bold dark:text-white">Go Back</span>
			</button>
			<ViewInvoiceBar
				status={invoice.Status}
				editClick={handleEditClick}
				deleteToggle={deleteToggle}
				handleMarkAsPaid={handleMarkAsPaid}
			/>
			<ViewInvoiceDetails invoice={invoice} />
			{isDeleting && (
				<>
					<div
						className={`absolute left-0 top-0 z-40 h-full w-full overscroll-none bg-gray-800 opacity-30 transition duration-500`}
					></div>
					<div
						className={`absolute left-1/2 top-1/2 z-50 flex h-[220px] w-[327px] -translate-x-1/2 -translate-y-1/2 flex-col gap-3 rounded-lg bg-white p-8 transition ${
							deleteStatus === "preEnter" || deleteStatus === "exiting"
								? "scale-75 transform opacity-0"
								: ""
						}`}
					>
						{!deleteMutation.isLoading ? (
							<>
								<h2 className="text-2xl font-bold text-gray-800">Confirm Deletion</h2>
								<p className="text-sm text-gray-400">
									Are you sure you want to delete invoice {`#${getShortID(invoice.ID)}`}? This action
									cannot be undone.
								</p>
								<div className="flex justify-end gap-2">
									<button
										className="h-[48px] w-[91px] rounded-full bg-[#F9FAFE] text-indigo-200"
										onClick={() => deleteToggle(false)}
									>
										Cancel
									</button>
									<button
										className="h-[48px] w-[89px] rounded-full bg-red-400 text-white"
										onClick={() => {
											deleteMutation.mutate(invoice.ID);
										}}
									>
										Delete
									</button>
								</div>
							</>
						) : (
							<div>
								<IconSpinning />
							</div>
						)}
					</div>
				</>
			)}
		</div>
	);
}

export type ViewInvoiceBarProps = {
	status: string;
	editClick: () => void;
	deleteToggle: (b: boolean) => void;
	handleMarkAsPaid: () => void;
};
export function ViewInvoiceBar({ status, editClick, deleteToggle, handleMarkAsPaid }: ViewInvoiceBarProps) {
	return (
		<div className="fixed bottom-0 left-0 flex h-24 w-screen items-center justify-center bg-white px-8 dark:bg-indigo-800 md:relative md:h-[88px] md:w-full md:justify-between md:rounded-md">
			<div className="hidden items-center gap-5 md:flex">
				<p className="text-sm text-[#858BB2]">Status</p>
				<InvoiceStatusBox status={status} />
			</div>
			<div className="flex justify-center gap-2">
				<button
					onClick={editClick}
					className="h-12 w-[73px] rounded-full bg-[#F9FAFE] text-indigo-200 dark:bg-gray-600 dark:text-gray-200"
				>
					Edit
				</button>
				<button className="h-12 w-[89px] rounded-full bg-red-400 text-white" onClick={() => deleteToggle(true)}>
					Delete
				</button>
				<button
					className={`${
						status === "paid" && "hidden"
					} h-12 w-[149px] rounded-full bg-purple-400 text-white md:w-[131px]`}
					onClick={handleMarkAsPaid}
				>
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
					<p className="dark:text-white">{invoice.ClientEmail}</p>
				</div>
				<div className="col-start-2 row-span-2">
					<p className="pb-2 pt-4 text-sm text-indigo-200 dark:text-gray-200">Bill To</p>
					<p className="dark:text-white">{invoice.ClientName}</p>
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
					{invoice.Items.map((item: Item, i: number) => (
						<Fragment key={i}>
							<p className="col-span-2 dark:text-white">{item.Name}</p>
							<p className="dark:text-white">{item.Quantity}</p>
							<p className="dark:text-white">{`£${item.Price}`}</p>
							<p className="dark:text-white">{`£${item.Total}`}</p>
						</Fragment>
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

export { ViewInvoice };
