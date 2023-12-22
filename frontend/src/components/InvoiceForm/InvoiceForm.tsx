import { Invoice } from "../../types";
import { useForm, useFieldArray, FieldErrors } from "react-hook-form";
import { IconDelete, IconLeftArrow, IconPlus, IconSpinning } from "../Icons";
import { ChangeEvent, useEffect, useState } from "react";
import { editInvoice, newInvoice } from "../../api";
import { useMutation, useQueryClient } from "react-query";
import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup";

type InvoiceFormProps = {
	state: "new" | "edit";
	invoice: Invoice | null;
	toggle: (f: boolean) => void;
};

const invoiceSchema = yup.object({
	ID: yup.string(),
	Description: yup.string().required("Description is required"),
	PaymentDue: yup.string().required("Invoice Due is required"),
	CreatedAt: yup.string(),
	PaymentTerms: yup.number().min(1).required("PaymentTerms is required"),
	ClientName: yup.string().required("Client Name is required"),
	ClientEmail: yup.string().email("Email format is not valid").required("Client Email is required"),
	ClientAddress: yup.object({
		Street: yup.string().required("Client Street is required"),
		City: yup.string().required("Client City is required"),
		PostCode: yup.string().required("Client PostCode is required"),
		Country: yup.string().required("Client Country is required"),
	}),
	SenderAddress: yup.object({
		Street: yup.string().required("Sender Street is required"),
		City: yup.string().required("Sender City is required"),
		PostCode: yup.string().required("Sender PostCode is required"),
		Country: yup.string().required("Sender Country is required"),
	}),
	Items: yup
		.array()
		.of(
			yup.object({
				Name: yup.string().required("Item Name is required"),
				Price: yup.number().min(0.01).required("Item Price is required"),
				Quantity: yup.number().min(1).required("Item Quantity is required"),
				Total: yup.number(),
			}),
		)
		.min(1)
		.required("Items are required"),
	Total: yup.number(),
});

function InvoiceForm({ state, invoice, toggle }: InvoiceFormProps) {
	const inv = invoice
		? invoice
		: ({
				ID: "",
				Description: "",
				PaymentDue: "",
				CreatedAt: "",
				PaymentTerms: 0,
				ClientName: "",
				ClientEmail: "",
				Status: "draft",
				ClientAddress: { Street: "", City: "", PostCode: "", Country: "" },
				SenderAddress: { Street: "", City: "", PostCode: "", Country: "" },
				Items: [],
				Total: 0,
		  } as Invoice);
	const form = useForm({ defaultValues: inv, resolver: yupResolver(invoiceSchema) });
	const { register, handleSubmit, control, getValues, setValue, formState } = form;
	const { errors } = formState;
	const { fields, append, remove } = useFieldArray({
		name: "Items",
		control,
	});

	const queryClient = useQueryClient();
	const newMutation = useMutation(newInvoice, {
		onSuccess: () => {
			queryClient.invalidateQueries("invoices");
			toggle(false);
		},
	});
	const editMutation = useMutation(editInvoice, {
		onSuccess: () => {
			queryClient.invalidateQueries("invoices");
			toggle(false);
		},
	});

	const onSaveSubmit = (data: any) => {
		setTotal(data);
		data.Status = "pending";
		data.PaymentDue = `${data.PaymentDue}T00:00:00Z`;
		newMutation.mutate(data);
	};

	const onEditSubmit = (data: any) => {
		setTotal(data);
		data.PaymentDue = `${data.PaymentDue}T00:00:00Z`;
		editMutation.mutate(data);
	};

	const onDraftSubmit = (data: any) => {
		setTotal(data);
		data.Status = "draft";
		data.PaymentDue = `${data.PaymentDue}T00:00:00Z`;
		newMutation.mutate(data);
	};

	const setTotal = (i: Invoice) => {
		let total = 0;
		for (let item of i.Items) {
			item.Total = item.Quantity * item.Price;
			total += item.Total;
		}
		i.Total = total;
	};

	const handleItemPriceChange = (e: ChangeEvent<HTMLInputElement>, i: number) => {
		const item = getValues(`Items.${i}`);
		const total = item.Quantity * e.target.value;
		setValue(`Items.${i}.Total`, total);
	};

	const handleItemQuantityChange = (e: ChangeEvent<HTMLInputElement>, i: number) => {
		const item = getValues(`Items.${i}`);
		const total = e.target.value * item.Price;
		setValue(`Items.${i}.Total`, total);
	};

	const handleCancel = (event: any) => {
		event.preventDefault();
		toggle(false);
	};

	useEffect(() => {
		console.log(errors);
	}, [errors]);

	return (
		<section className="relative  mt-[80px] h-screen overflow-y-auto overflow-x-hidden overscroll-contain p-8 md:p-14 lg:ml-[103px] lg:mt-0">
			{!newMutation.isLoading ? (
				<>
					<button className="hidden w-24 items-center gap-5 lg:flex" onClick={handleCancel}>
						<IconLeftArrow />
						<span className="mt-1 font-bold dark:text-white">Go Back</span>
					</button>
					<form noValidate className="mb-14 md:mb-8 lg:mb-0">
						<h3 className="mt-10 font-bold text-purple-400">Bill From</h3>
						<div className="mt-6 grid grid-cols-2 gap-6">
							<div className="col-span-2 flex flex-col gap-2">
								<label htmlFor="senderStreet" className="text-sm text-indigo-200 dark:text-gray-200">
									Street Address
								</label>
								<input
									id="senderStreet"
									type="text"
									className={`rounded-md border  p-4 text-gray-800  dark:bg-indigo-800 dark:text-white ${
										errors.SenderAddress?.Street
											? "border-red-400"
											: "border-gray-200 dark:border-gray-600"
									}`}
									{...register("SenderAddress.Street")}
								/>
							</div>
							<div className="flex flex-col gap-2">
								<label htmlFor="senderCity" className="text-sm text-indigo-200 dark:text-gray-200">
									City
								</label>
								<input
									id="senderCity"
									type="text"
									className={`rounded-md border p-4 text-gray-800 dark:bg-indigo-800 dark:text-white ${
										errors.SenderAddress?.City
											? "border-red-400"
											: "border-gray-200 dark:border-gray-600"
									}`}
									{...register("SenderAddress.City")}
								/>
							</div>
							<div className="flex flex-col gap-2">
								<label htmlFor="senderPostCode" className="text-sm text-indigo-200 dark:text-gray-200">
									Post Code
								</label>
								<input
									id="senderPostCode"
									type="text"
									className={`rounded-md border  p-4 text-gray-800  dark:bg-indigo-800 dark:text-white ${
										errors.SenderAddress?.PostCode
											? "border-red-400"
											: "border-gray-200 dark:border-gray-600"
									}`}
									{...register("SenderAddress.PostCode")}
								/>
							</div>
							<div className="col-span-2 flex flex-col gap-2">
								<label htmlFor="senderCountry" className="text-sm text-indigo-200 dark:text-gray-200">
									Country
								</label>
								<input
									id="senderCountry"
									type="text"
									className={`rounded-md border  p-4 text-gray-800 dark:bg-indigo-800 dark:text-white ${
										errors.SenderAddress?.Country
											? "border-red-400"
											: "border-gray-200 dark:border-gray-600"
									}`}
									{...register("SenderAddress.Country")}
								/>
							</div>
						</div>
						<h3 className="mt-10 font-bold text-purple-400">Bill To</h3>
						<div className="mt-6 grid grid-cols-2 gap-6 md:grid-cols-3">
							<div className="col-span-2 flex flex-col gap-2 md:col-span-3">
								<label htmlFor="clientName" className="text-sm text-indigo-200 dark:text-gray-200">
									Client's Name
								</label>
								<input
									id="clientName"
									type="text"
									className={`rounded-md border  p-4 text-gray-800 dark:bg-indigo-800 dark:text-white ${
										errors.ClientName ? "border-red-400" : "border-gray-200 dark:border-gray-600"
									}`}
									{...register("ClientName")}
								/>
							</div>
							<div className="col-span-2 flex flex-col gap-2 md:col-span-3">
								<label htmlFor="clientEmail" className="text-sm text-indigo-200 dark:text-gray-200">
									Client's Email
								</label>
								<input
									id="clientEmail"
									type="text"
									className={`rounded-md border p-4 text-gray-800 dark:bg-indigo-800 dark:text-white ${
										errors.ClientEmail ? "border-red-400" : "border-gray-200 dark:border-gray-600"
									}`}
									{...register("ClientEmail")}
								/>
							</div>
							<div className="col-span-2 flex flex-col gap-2 md:col-span-3">
								<label htmlFor="clientStreet" className="text-sm text-indigo-200 dark:text-gray-200">
									Street Address
								</label>
								<input
									id="clientStreet"
									type="text"
									className={`rounded-md border  p-4 text-gray-800 dark:bg-indigo-800 dark:text-white ${
										errors.ClientAddress?.Street
											? "border-red-400"
											: "border-gray-200 dark:border-gray-600"
									}`}
									{...register("ClientAddress.Street")}
								/>
							</div>
							<div className="flex flex-col gap-2 md:col-span-1">
								<label htmlFor="clientCity" className="text-sm text-indigo-200 dark:text-gray-200">
									City
								</label>
								<input
									id="clientCity"
									type="text"
									className={`rounded-md border p-4 text-gray-800 dark:bg-indigo-800 dark:text-white ${
										errors.ClientAddress?.City
											? "border-red-400"
											: "border-gray-200 dark:border-gray-600"
									}`}
									{...register("ClientAddress.City")}
								/>
							</div>
							<div className="flex flex-col gap-2">
								<label htmlFor="clientPostCode" className="text-sm text-indigo-200 dark:text-gray-200">
									Post Code
								</label>
								<input
									id="clientPostCode"
									type="text"
									className={`rounded-md border p-4 text-gray-800 dark:bg-indigo-800 dark:text-white ${
										errors.ClientAddress?.PostCode
											? "border-red-400"
											: "border-gray-200 dark:border-gray-600"
									}`}
									{...register("ClientAddress.PostCode")}
								/>
							</div>
							<div className="col-span-2 flex flex-col gap-2 md:col-span-1">
								<label htmlFor="clientCountry" className="text-sm text-indigo-200 dark:text-gray-200">
									Country
								</label>
								<input
									id="clientCountry"
									type="text"
									className={`rounded-md border  p-4 text-gray-800 dark:bg-indigo-800 dark:text-white ${
										errors.ClientAddress?.Country
											? "border-red-400"
											: "border-gray-200 dark:border-gray-600"
									}`}
									{...register("ClientAddress.Country")}
								/>
							</div>
						</div>
						<div className="mt-2 grid gap-6">
							<div className="flex flex-col gap-2">
								<label htmlFor="paymentDue" className="text-sm text-indigo-200 dark:text-gray-200">
									Invoice Date
								</label>
								<input
									id="paymentDue"
									type="date"
									className={`rounded-md border p-4 text-gray-800 dark:bg-indigo-800 dark:text-white ${
										errors.PaymentDue ? "border-red-400" : "border-gray-200 dark:border-gray-600"
									}`}
									{...register("PaymentDue")}
								/>
							</div>
							<div className="flex flex-col gap-2">
								<label htmlFor="paymentTerms" className="text-sm text-indigo-200 dark:text-gray-200">
									Payment Terms
								</label>
								<select
									id="paymentTerms"
									className={`rounded-md border p-4 text-gray-800 dark:bg-indigo-800 dark:text-white ${
										errors.PaymentTerms ? "border-red-400" : "border-gray-200 dark:border-gray-600"
									}`}
									{...register("PaymentTerms", { valueAsNumber: true })}
								>
									<option value="1">Net 1 Day</option>
									<option value="7">Net 7 Days</option>
									<option value="14">Net 14 Days</option>
									<option value="30">Net 30 Days</option>
								</select>
							</div>
							<div className="flex flex-col gap-2">
								<label htmlFor="description" className="text-sm text-indigo-200 dark:text-gray-200">
									Project Description
								</label>
								<input
									id="description"
									type="text"
									className={`rounded-md border p-4 text-gray-800 dark:bg-indigo-800 dark:text-white ${
										errors.Description ? "border-red-400" : "border-gray-200 dark:border-gray-600"
									}`}
									{...register("Description")}
								/>
							</div>
						</div>
						<h3 className="my-6 text-lg font-bold text-[#777F98]">Item List</h3>
						<div className="flex flex-col gap-6">
							{fields?.map((_, i: number) => (
								<div className="grid grid-cols-6 justify-center gap-2 md:gap-6" key={i}>
									<div className="col-span-6 flex flex-col gap-2">
										<label
											htmlFor={`itemName${i}`}
											className="text-sm text-indigo-200 dark:text-gray-200"
										>
											Item Name
										</label>
										<input
											id={`itemName${i}`}
											type="text"
											className={`rounded-md border p-4 text-gray-800 dark:bg-indigo-800 dark:text-white ${
												errors.Items?.[i]?.Name
													? "border-red-400"
													: "border-gray-200 dark:border-gray-600"
											}`}
											{...register(`Items.${i}.Name`)}
										/>
									</div>
									<div className="flex flex-col gap-2">
										<label
											htmlFor={`itemQuantity${i}`}
											className="text-sm text-indigo-200 dark:text-gray-200"
										>
											Qty.
										</label>
										<input
											id={`itemQuantity${i}`}
											type="number"
											className={`rounded-md border p-4 text-gray-800 dark:bg-indigo-800 dark:text-white ${
												errors.Items?.[i]?.Quantity
													? "border-red-400"
													: "border-gray-200 dark:border-gray-600"
											}`}
											{...register(`Items.${i}.Quantity`, { valueAsNumber: true })}
											onChange={(e) => handleItemQuantityChange(e, i)}
										/>
									</div>
									<div className="col-span-2 flex flex-col gap-2">
										<label
											htmlFor={`itemPrice${i}`}
											className="text-sm text-indigo-200 dark:text-gray-200"
										>
											Price
										</label>
										<input
											id={`itemPrice${i}`}
											type="number"
											className={`rounded-md border p-4 text-gray-800 dark:bg-indigo-800 dark:text-white ${
												errors.Items?.[i]?.Price
													? "border-red-400"
													: "border-gray-200 dark:border-gray-600"
											}`}
											{...register(`Items.${i}.Price`, { valueAsNumber: true })}
											onChange={(e) => handleItemPriceChange(e, i)}
										/>
									</div>
									<div className="col-span-2 flex flex-col gap-2">
										<label
											htmlFor={`itemTotal${i}`}
											className="text-sm text-indigo-200 dark:text-gray-200"
										>
											Total
										</label>
										<input
											id={`itemTotal${i}`}
											type="text"
											className="rounded-md p-4 text-gray-400 dark:text-gray-200"
											{...register(`Items.${i}.Total`, {
												disabled: true,
											})}
										/>
									</div>
									<button type="button" onClick={() => remove(i)} className="mt-6 justify-self-end">
										<IconDelete />
									</button>
								</div>
							))}
							<button
								className="flex h-12 items-center justify-center gap-2 rounded-full bg-[#F9FAFE] text-[#979797] dark:bg-gray-600"
								onClick={() => append({ Name: "", Quantity: 0, Price: 0, Total: 0 })}
								type="button"
							>
								<IconPlus /> <span className="-mb-1">Add New Item</span>
							</button>
						</div>
						<div className="flex flex-col">
							{Object.keys(errors).length !== 0 && <FormErrors errors={errors} />}
						</div>
						<div className="flex h-[91px] w-full items-center justify-end gap-2">
							{state === "edit" ? (
								<>
									<CancelButton text={"Cancel"} onClick={handleCancel} />
									<SaveButton text={"Save Changes"} onClick={handleSubmit(onEditSubmit)} />
								</>
							) : (
								<>
									<CancelButton text={"Discard"} onClick={handleCancel} />
									<button
										className="flex h-12 items-center justify-center rounded-full bg-[#373B53] p-4 text-gray-400 dark:text-gray-200"
										onClick={handleSubmit(onDraftSubmit)}
									>
										Save as Draft
									</button>
									<SaveButton text={"Save & Send"} onClick={handleSubmit(onSaveSubmit)} />
								</>
							)}
						</div>
					</form>
				</>
			) : (
				<div className="flex h-full w-full items-center justify-center">
					<IconSpinning />
				</div>
			)}
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
			type="button"
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

type FormErrorsProps = {
	errors: any;
};

function FormErrors({ errors }: FormErrorsProps) {
	const keys = Object.keys(errors);
	const [errorVals, setErrorVals] = useState<string[]>([]);

	useEffect(() => {
		console.log("effect");
		if (keys.includes("Items") && keys.length >= 2) {
			setErrorVals(["All fields must be added", "An item must be added"]);
		} else if (!keys.includes("Items") && keys.length >= 1) {
			setErrorVals(["All fields must be added"]);
		} else if (keys.includes("Items") && keys.length === 1) {
			setErrorVals(["An item must be added"]);
		} else if (keys.length === 0) {
			setErrorVals([]);
		}
	}, [errors]);

	return (
		<div className="mt-6 flex flex-col">
			{errorVals.map((val: string, i: number) => (
				<p key={i} className="text-sm text-red-400">{`- ${val}`}</p>
			))}
		</div>
	);
}

export { InvoiceForm };
