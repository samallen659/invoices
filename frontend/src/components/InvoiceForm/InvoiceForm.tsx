import { Invoice } from "../../types";
import { useForm } from "react-hook-form";
import { Item } from "../../types";
import { IconLeftArrow, IconPlus } from "../Icons";

type InvoiceFormProps = {
	state: "new" | "edit";
	invoice: Invoice | null;
	toggle: (f: boolean) => void;
};

function InvoiceForm({ state, invoice, toggle }: InvoiceFormProps) {
	const inv = invoice
		? invoice
		: ({
				ID: "",
				Description: "",
				PaymentDue: new Date(),
				PaymentTerms: 0,
				ClientName: "",
				ClientEmail: "",
				Status: "draft",
				ClientAddress: { Street: "", City: "", PostCode: "", Country: "" },
				SenderAddress: { Street: "", City: "", PostCode: "", Country: "" },
				Items: [],
				Total: 0,
		  } as Invoice);
	const form = useForm({ defaultValues: inv });
	const { register, handleSubmit } = form;
	const items = form.watch("Items");

	const onSubmit = (data: Invoice) => {
		console.log("Form Submitted", data);
	};

	const handleAddItem = () => {
		let items = form.getValues("Items");
		let item = { Name: "", Quantity: 0, Price: 0, Total: 0 } as Item;
		items.push(item);
		form.setValue("Items", items);
	};

	return (
		<section className="relative mt-[80px] h-screen overflow-y-auto overflow-x-hidden overscroll-contain p-8 dark:bg-indigo-600 md:p-14 lg:ml-[103px] lg:mt-0">
			<button className="hidden w-24 items-center gap-5 lg:flex" onClick={() => toggle(false)}>
				<IconLeftArrow />
				<span className="mt-1 font-bold dark:text-white">Go Back</span>
			</button>
			<form onSubmit={handleSubmit(onSubmit)}>
				<h3 className="mt-10 font-bold text-purple-400">Bill From</h3>
				<div className="mt-6 grid grid-cols-2 gap-6">
					<div className="col-span-2 flex flex-col gap-2">
						<label htmlFor="senderStreet" className="form-label">
							Street Address
						</label>
						<input
							id="senderStreet"
							type="text"
							className="form-input"
							{...register("SenderAddress.Street")}
						/>
					</div>
					<div className="flex flex-col gap-2">
						<label htmlFor="senderCity" className="form-label">
							City
						</label>
						<input id="senderCity" type="text" className="form-input" {...register("SenderAddress.City")} />
					</div>
					<div className="flex flex-col gap-2">
						<label htmlFor="senderPostCode" className="form-label">
							Post Code
						</label>
						<input
							id="senderPostCode"
							type="text"
							className="form-input"
							{...register("SenderAddress.PostCode")}
						/>
					</div>
					<div className="col-span-2 flex flex-col gap-2">
						<label htmlFor="senderCountry" className="form-label">
							Country
						</label>
						<input
							id="senderCountry"
							type="text"
							className="form-input"
							{...register("SenderAddress.Country")}
						/>
					</div>
				</div>
				<h3 className="mt-10 font-bold text-purple-400">Bill To</h3>
				<div className="mt-6 grid grid-cols-2 gap-6">
					<div className="col-span-2 flex flex-col gap-2">
						<label htmlFor="clientName" className="form-label">
							Client's Name
						</label>
						<input id="clientName" type="text" className="form-input" {...register("ClientName")} />
					</div>
					<div className="col-span-2 flex flex-col gap-2">
						<label htmlFor="clientEmail" className="form-label">
							Client's Email
						</label>
						<input id="clientEmail" type="text" className="form-input" {...register("ClientEmail")} />
					</div>
					<div className="col-span-2 flex flex-col gap-2">
						<label htmlFor="clientStreet" className="form-label">
							Street Address
						</label>
						<input
							id="clientStreet"
							type="text"
							className="form-input"
							{...register("ClientAddress.Street")}
						/>
					</div>
					<div className="flex flex-col gap-2">
						<label htmlFor="clientCity" className="form-label">
							City
						</label>
						<input id="clientCity" type="text" className="form-input" {...register("ClientAddress.City")} />
					</div>
					<div className="flex flex-col gap-2">
						<label htmlFor="clientPostCode" className="form-label">
							Post Code
						</label>
						<input
							id="clientPostCode"
							type="text"
							className="form-input"
							{...register("ClientAddress.PostCode")}
						/>
					</div>
					<div className="col-span-2 flex flex-col gap-2">
						<label htmlFor="clientCountry" className="form-label">
							Country
						</label>
						<input
							id="clientCountry"
							type="text"
							className="form-input"
							{...register("ClientAddress.Country")}
						/>
					</div>
				</div>
				<div className="mt-2 grid gap-6">
					<div className="flex flex-col gap-2">
						<label htmlFor="paymentDue" className="form-label">
							Invoice Date
						</label>
						<input id="paymentDue" type="date" className="form-input" {...register("PaymentDue")} />
					</div>
					<div className="flex flex-col gap-2">
						<label htmlFor="paymentTerms" className="form-label">
							Payment Terms
						</label>
						<select id="paymentTerms" className="form-input" {...register("PaymentTerms")}>
							<option value="1">Net 1 Day</option>
							<option value="7">Net 7 Days</option>
							<option value="14">Net 14 Days</option>
							<option value="30">Net 30 Days</option>
						</select>
					</div>
				</div>
				<h3 className="my-6 text-lg font-bold text-[#777F98]">Item List</h3>
				<div className="flex flex-col gap-6">
					{items?.map((_, i: number) => (
						<div className="grid grid-cols-3 gap-6">
							<div className="col-span-3 flex flex-col gap-2">
								<label htmlFor={`itemName${i}`} className="form-label">
									Item Name
								</label>
								<input
									id={`itemName${i}`}
									type="text"
									className="form-input"
									{...register(`Items.${i}.Name`)}
								/>
							</div>
							<div className="flex flex-col gap-2">
								<label htmlFor={`itemQuantity${i}`} className="form-label">
									Qty.
								</label>
								<input
									id={`itemQuantity${i}`}
									type="text"
									className="form-input"
									{...register(`Items.${i}.Quantity`)}
								/>
							</div>
							<div className="flex flex-col gap-2">
								<label htmlFor={`itemPrice${i}`} className="form-label">
									Price
								</label>
								<input
									id={`itemPrice${i}`}
									type="text"
									className="form-input"
									{...register(`Items.${i}.Price`)}
								/>
							</div>
							<div className="flex flex-col gap-2">
								<label htmlFor={`itemTotal${i}`} className="form-label">
									Total
								</label>
								<input
									id={`itemTotal${i}`}
									type="text"
									className="form-input"
									{...register(`Items.${i}.Total`)}
								/>
							</div>
						</div>
					))}
					<button
						className="flex h-12 items-center justify-center gap-2 rounded-full bg-[#F9FAFE] text-[#979797] dark:bg-gray-600"
						onClick={handleAddItem}
					>
						<IconPlus /> <span className="-mb-1">Add New Item</span>
					</button>
				</div>
				<div className="flex h-[91px] w-full items-center justify-end gap-2 p-6">
					{state === "edit" ? (
						<>
							<CancelButton text={"Cancel"} onClick={() => console.log("Cancel edit")} />
							<SaveButton text={"Save Changes"} onClick={() => console.log("Save Changes edit")} />
						</>
					) : (
						<>
							<CancelButton text={"Discard"} onClick={() => console.log("Discard new")} />
							<button
								className="flex h-12 items-center justify-center rounded-full bg-[#373B53] p-4 text-gray-400 dark:text-gray-200"
								onClick={() => console.log("Save as Draft new")}
							>
								Save as Draft
							</button>
							<SaveButton text={"Save & Send"} onClick={() => console.log("Save & Send new")} />
						</>
					)}
				</div>
			</form>
		</section>
	);
}

type SaveButtonProps = {
	text: string;
	onClick: () => void;
};

function SaveButton({ text, onClick }: SaveButtonProps) {
	return (
		<button
			className="flex h-12 items-center justify-center rounded-full bg-purple-400 p-4 text-white"
			onClick={onClick}
		>
			{text}
		</button>
	);
}

type CancelButtonProps = {
	text: string;
	onClick: () => void;
};

function CancelButton({ text, onClick }: CancelButtonProps) {
	return (
		<button
			className="flex h-12 items-center justify-center rounded-full bg-[#F9FAFE] p-4 text-center text-indigo-200 dark:bg-gray-600 dark:text-gray-200"
			onClick={onClick}
		>
			{text}
		</button>
	);
}

export { InvoiceForm };
