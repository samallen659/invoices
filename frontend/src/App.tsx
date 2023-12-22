import { useState } from "react";
import { SideBar } from "./components/SideBar/SideBar";
import { MenuBar } from "./components/MenuBar/MenuBar";
import { useQuery } from "react-query";
import { getAllInvoices } from "./api";
import { InvoiceList } from "./components/InvoiceList/InvoiceList";
import { useTransition } from "react-transition-state";
import { ViewInvoice } from "./components/ViewInvoice/ViewInvoice";
import { InvoiceForm } from "./components/InvoiceForm/InvoiceForm";
import { IconSpinning, IconIllustrationEmpty } from "./components/Icons";

export type InvoiceIdx = number;

function App() {
	const [isDark, setIsDark] = useState(false);
	const [filter, setFilter] = useState<string>("");

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
				<section className="mx-auto mt-24 w-full p-6 md:mt-36 md:max-w-[672px] lg:mt-20 lg:max-w-3xl lg:p-0 lg:pl-[104px] xl:max-w-4xl 2xl:max-w-6xl">
					{!isViewing && (
						<div>
							<MenuBar
								setFormState={setFormState}
								toggle={formToggle}
								setFilter={setFilter}
								filter={filter}
							/>
							{isLoading ? (
								<div className="absolute left-1/2 top-1/2 -translate-x-1/2  -translate-y-1/2 pt-20 lg:pl-[103px] lg:pt-0">
									<IconSpinning />
								</div>
							) : (
								<div className="mt-8 md:mt-14 lg:mt-16">
									{invoices && invoices?.length > 0 ? (
										<InvoiceList
											invoices={invoices}
											handleInvoiceIdxChange={handleInvoiceIdxChange}
											filter={filter}
										/>
									) : (
										<div className="absolute left-1/2 top-1/2 flex -translate-x-1/2 -translate-y-1/2 flex-col items-center justify-center gap-10 pt-20 lg:pl-[103px] lg:pt-0">
											<IconIllustrationEmpty />
											<div className="flex flex-col gap-6">
												<h2 className="text-2xl font-bold text-gray-800 dark:text-white">
													There is nothing here
												</h2>
												<p className="text-center text-sm text-gray-400">
													Create an invoice by clicking the
													<br />
													New <span className="hidden md:inline">invoice </span>
													button and get started
												</p>
											</div>
										</div>
									)}
								</div>
							)}
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
							{invoices && (
								<ViewInvoice
									invoice={invoices[invoiceIdx]}
									viewToggle={viewToggle}
									formToggle={formToggle}
									setFormState={setFormState}
								/>
							)}
						</div>
					)}
				</section>
				{isForm && (
					<>
						<div className="absolute left-0 top-0 h-full w-full bg-gray-800 opacity-30 "></div>
						<div
							className={`absolute h-full w-full transition duration-500 md:w-2/3 xl:w-1/2 ${
								formStatus === "preEnter" || formStatus === "exiting" ? "-translate-x-full" : ""
							}`}
						>
							<div className="h-full w-full bg-white dark:bg-indigo-600">
								<InvoiceForm
									state={formState}
									invoice={
										invoices != undefined && formState === "edit" ? invoices[invoiceIdx] : null
									}
									toggle={formToggle}
								/>
							</div>
						</div>
					</>
				)}
			</main>
		</div>
	);
}

export { App };
