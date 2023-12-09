type InvoiceStatusBoxProps = {
	status: string;
};

function InvoiceStatusBox({ status }: InvoiceStatusBoxProps) {
	return (
		<div
			className={`h-10 w-[104px] rounded-md bg-opacity-10 font-bold ${getStatusColor(
				status,
			)} flex items-center justify-center capitalize`}
		>
			{status}
		</div>
	);
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

export { InvoiceStatusBox };
