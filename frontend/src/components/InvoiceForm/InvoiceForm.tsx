import { Invoice } from "../../types";
import { useFormik } from "formik";
import { InvoiceItem } from "../../types";

type InvoiceFormProps = {
	state: "new" | "edit";
	invoice: Invoice | null;
	toggle: (f: boolean) => void;
};

function InvoiceForm({ state, invoice, toggle }: InvoiceFormProps) {
	const formik = useFormik({
		initialValues: {
			senderStreet: invoice ? invoice.SenderAddress.Street : "",
			senderCity: invoice ? invoice.SenderAddress.City : "",
			senderPostCode: invoice ? invoice.SenderAddress.PostCode : "",
			senderCountry: invoice ? invoice.SenderAddress.Country : "",
			clientName: invoice ? invoice.Client.ClientName : "",
			clientEmail: invoice ? invoice.Client.ClientEmail : "",
			clientStreet: invoice ? invoice.ClientAddress.Street : "",
			clientCity: invoice ? invoice.ClientAddress.City : "",
			clientPostCode: invoice ? invoice.ClientAddress.PostCode : "",
			clientCountry: invoice ? invoice.ClientAddress.Country : "",
			paymentDue: invoice ? invoice.PaymentDue : "",
			paymentTerms: invoice ? invoice.PaymentTerms : "",
			description: invoice ? invoice.Description : "",
			invoiceItems: invoice ? invoice.InvoiceItems : Array<InvoiceItem>,
		},
		onSubmit: (values) => {
			console.log(JSON.stringify(values));
		},
	});
	return (
		<section className="mt-[80px] p-8 md:p-14 lg:ml-[103px] lg:mt-0">
			<button className="flex w-24 items-center gap-5" onClick={() => toggle(false)}>
				<IconArrowLeft />
				<span className="mt-1 font-bold dark:text-white">Go Back</span>
			</button>
			<form onSubmit={formik.handleSubmit}>
				<div className="grid-row-3 grid-col-2 grid gap-6">
					<div className="col-span-2 flex flex-col gap-2">
						<label htmlFor="senderStreet">Street Address</label>
						<input
							id="senderStreet"
							name="senderStreet"
							type="text"
							onChange={formik.handleChange}
							value={formik.values.senderStreet}
						/>
					</div>
					<div className="flex flex-col gap-2">
						<label htmlFor="senderCity">City</label>
						<input
							id="senderCity"
							name="senderCity"
							type="text"
							onChange={formik.handleChange}
							value={formik.values.senderCity}
						/>
					</div>
					<div className="flex flex-col gap-2">
						<label htmlFor="senderPostCode">Post Code</label>
						<input
							id="senderPostCode"
							name="senderPostCode"
							type="text"
							onChange={formik.handleChange}
							value={formik.values.senderPostCode}
						/>
					</div>
					<div className="col-span-2 flex flex-col gap-2">
						<label htmlFor="senderCountry">Country</label>
						<input
							id="senderCountry"
							name="senderCountry"
							type="text"
							onChange={formik.handleChange}
							value={formik.values.senderCountry}
						/>
					</div>
				</div>
				<div></div>
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
