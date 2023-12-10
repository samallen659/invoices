import { Invoice } from "../../types";

type InvoiceFormProps = {
	state: "new" | "edit";
	invoice: Invoice | null;
};

function InvoiceForm({ state, invoice }: InvoiceFormProps) {
	return <div></div>;
}

export { InvoiceForm };
