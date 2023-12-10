import { useState } from "react";
import { SideBar } from "./components/SideBar/SideBar";
import { MenuBar } from "./components/MenuBar/MenuBar";
import { useQuery } from "react-query";
import { getAllInvoices } from "./api";
import { InvoiceList } from "./components/InvoiceList/InvoiceList";
import { useTransition } from "react-transition-state";
import { ViewInvoice, ViewInvoiceBar } from "./components/ViewInvoice/ViewInvoice";

export type InvoiceIdx = number;

function App() {
	const [isDark, setIsDark] = useState(false);

	const [invoiceIdx, setInvoiceIdx] = useState<InvoiceIdx>(0);
	const [{ status: viewStatus, isMounted: isViewing }, viewToggle] = useTransition({
		timeout: 500,
		mountOnEnter: true,
		unmountOnExit: true,
		preEnter: true,
	});

	const [formState, setFormState] = useState<"new" | "edit">("new");
	const [{ status: formStatus, isMounted: isForm }, formToggle] = useTransition({
		timeout: 500,
		mountOnEnter: true,
		unmountOnExit: true,
		preEnter: true,
	});

	const handleInvoiceIdxChange = (i: InvoiceIdx) => {
		setInvoiceIdx(i);
		viewToggle(!isViewing);
	};

	const toggleIsDark = () => {
		setIsDark((prev) => !prev);
	};

	const { data: invoices, isLoading } = useQuery({
		queryKey: ["invoices"],
		queryFn: getAllInvoices,
	});

	return (
		<div className={`bg-gray-800 ${isDark ? "dark" : "light"}`}>
			<main className="relative flex min-h-screen flex-col bg-off-white dark:bg-indigo-600 lg:flex-row">
				<SideBar toggleIsDark={toggleIsDark} isDark={isDark} />
				<section className="mx-auto mt-24 w-full p-6 md:mt-36 md:max-w-[672px] lg:max-w-3xl lg:p-0 lg:pl-[104px] xl:max-w-4xl 2xl:max-w-6xl">
					{!isViewing && (
						<div>
							<MenuBar setFormState={setFormState} toggle={formToggle} />
							<div className="mt-8 md:mt-14 lg:mt-16">
								{invoices && (
									<InvoiceList invoices={invoices} handleInvoiceIdxChange={handleInvoiceIdxChange} />
								)}
							</div>
						</div>
					)}
					{isViewing && (
						<div
							className={`transition duration-500 ${
								viewStatus === "preEnter" || viewStatus === "exiting"
									? "scale-75 transform opacity-0"
									: ""
							}`}
						>
							{invoices && <ViewInvoice invoice={invoices[invoiceIdx]} toggle={viewToggle} />}
						</div>
					)}
				</section>
				{/* ViewInvoiceBar fixed to bottom for small screens */}
				{isViewing && (
					<div className="mt-24 md:hidden">
						{invoices && <ViewInvoiceBar status={invoices[invoiceIdx].Status} />}
					</div>
				)}
				{isForm && (
					<>
						<div className="left-0 top-0 hidden h-full w-full bg-gray-800 opacity-30 md:absolute md:block"></div>
						<div
							className={`absolute h-full w-1/2 bg-white transition duration-500 ${
								formStatus === "preEnter" || formStatus === "exiting" ? "-translate-x-1/2" : ""
							}`}
						>
							<InvoiceForm state={formState} invoice={invoices[invoiceIdx]} />
						</div>
					</>
				)}
			</main>
		</div>
	);
}

export { App };
