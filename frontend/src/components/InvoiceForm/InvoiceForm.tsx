import { Invoice } from "../../types";
import { useForm } from "react-hook-form";
import { Item } from "../../types";

type InvoiceFormProps = {
	state: "new" | "edit";
	invoice: Invoice | null;
	toggle: (f: boolean) => void;
};

function InvoiceForm({ state, invoice, toggle }: InvoiceFormProps) {
	const form = invoice ? useForm({ defaultValues: invoice }) : useForm<Invoice>();
	const { register, handleSubmit } = form;

	const onSubmit = (data: Invoice) => {
		console.log("Form Submitted", data);
	};

	return (
		<section className="mt-[80px] h-screen overflow-y-auto overscroll-contain p-8 md:p-14 lg:ml-[103px] lg:mt-0">
			<button className="flex w-24 items-center gap-5" onClick={() => toggle(false)}>
				<IconArrowLeft />
				<span className="mt-1 font-bold dark:text-white">Go Back</span>
			</button>
			<form onSubmit={handleSubmit(onSubmit)}>
				<div className="grid-row-3 grid-col-2 grid gap-6">
					<div className="col-span-2 flex flex-col gap-2">
						<label htmlFor="senderStreet">Street Address</label>
						<input id="senderStreet" type="text" {...register("SenderAddress.Street")} />
					</div>
					<div className="flex flex-col gap-2">
						<label htmlFor="senderCity">City</label>
						<input id="senderCity" type="text" {...register("SenderAddress.City")} />
					</div>
					<div className="flex flex-col gap-2">
						<label htmlFor="senderPostCode">Post Code</label>
						<input id="senderPostCode" type="text" {...register("SenderAddress.PostCode")} />
					</div>
					<div className="col-span-2 flex flex-col gap-2">
						<label htmlFor="senderCountry">Country</label>
						<input id="senderCountry" type="text" {...register("SenderAddress.Country")} />
					</div>
				</div>
				<div className="grid grid-cols-2 gap-6">
					<div className="col-span-2 flex flex-col gap-2">
						<label htmlFor="clientName">Client's Name</label>
						<input id="clientName" type="text" {...register("ClientName")} />
					</div>
					<div className="col-span-2 flex flex-col gap-2">
						<label htmlFor="clientEmail">Client's Email</label>
						<input id="clientEmail" type="text" {...register("ClientEmail")} />
					</div>
					<div className="col-span-2 flex flex-col gap-2">
						<label htmlFor="clientStreet">Street Address</label>
						<input id="clientStreet" type="text" {...register("ClientAddress.Street")} />
					</div>
					<div className="flex flex-col gap-2">
						<label htmlFor="clientCity">City</label>
						<input id="clientCity" type="text" {...register("ClientAddress.City")} />
					</div>
					<div className="flex flex-col gap-2">
						<label htmlFor="clientPostCode">Post Code</label>
						<input id="clientPostCode" type="text" {...register("ClientAddress.PostCode")} />
					</div>
					<div className="col-span-2 flex flex-col gap-2">
						<label htmlFor="clientCountry">Country</label>
						<input id="clientCountry" type="text" {...register("ClientAddress.Country")} />
					</div>
				</div>
				<div className="mt-2 grid gap-6">
					<div className="flex flex-col gap-2">
						<label htmlFor="paymentDue">Invoice Date</label>
						<input id="paymentDue" type="date" {...register("PaymentDue")} />
					</div>
					<div className="flex flex-col gap-2">
						<label htmlFor="paymentTerms">Payment Terms</label>
						<select id="paymentTerms" {...register("PaymentTerms")}>
							<option value="1">Net 1 Day</option>
							<option value="7">Net 7 Days</option>
							<option value="14">Net 14 Days</option>
							<option value="30">Net 30 Days</option>
						</select>
					</div>
				</div>
				<div className="mt-6 flex flex-col gap-6">
					<h3>Item List</h3>
					{invoice?.Items.map((item: Item, i: number) => (
						<div className="grid grid-cols-3 gap-6">
							<div className="col-span-3 flex flex-col gap-2">
								<label htmlFor={`itemName${i}`}>Item Name</label>
								<input id={`itemName${i}`} type="text" {...register(`Items.${i}.Name`)} />
							</div>
							<div className="flex flex-col gap-2">
								<label htmlFor={`itemQuantity${i}`}>Qty.</label>
								<input id={`itemQuantity${i}`} type="text" {...register(`Items.${i}.Quantity`)} />
							</div>
							<div className="flex flex-col gap-2">
								<label htmlFor={`itemPrice${i}`}>Price</label>
								<input id={`itemPrice${i}`} type="text" {...register(`Items.${i}.Price`)} />
							</div>
							<div className="flex flex-col gap-2">
								<label htmlFor={`itemTotal${i}`}>Total</label>
								<input id={`itemTotal${i}`} type="text" {...register(`Items.${i}.Total`)} />
							</div>
						</div>
					))}
				</div>
			</form>
		</section>
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

export { InvoiceForm };
