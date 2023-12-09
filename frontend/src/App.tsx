import { useState } from "react";
import { SideBar } from "./components/SideBar/SideBar";
import { MenuBar } from "./components/MenuBar/MenuBar";
import { useQuery } from "react-query";
import { getAllInvoices } from "./api";
import { InvoiceList } from "./components/InvoiceList/InvoiceList";
import { useTransition } from "react-transition-state";

export type InvoiceIdx = number | null;

function App() {
	const [isDark, setIsDark] = useState(false);
	const [invoiceIdx, setInvoiceIdx] = useState<InvoiceIdx>(null);
	const [{ status, isMounted }, toggle] = useTransition({
		timeout: 500,
		mountOnEnter: true,
		unmountOnExit: true,
		preEnter: true,
	});

	const handleInvoiceIdxChange = (i: InvoiceIdx) => {
		setInvoiceIdx(i);
		toggle(!isMounted);
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
			<main className="flex min-h-screen flex-col bg-off-white dark:bg-indigo-600 lg:flex-row">
				<SideBar toggleIsDark={toggleIsDark} isDark={isDark} />
				<section className="mx-auto mt-24 w-full p-6 md:mt-36 md:max-w-[672px] lg:max-w-3xl lg:p-0 lg:pl-[104px] xl:max-w-4xl 2xl:max-w-6xl">
					{!isMounted && (
						<div>
							<MenuBar />
							<div className="mt-8 md:mt-14 lg:mt-16">
								{invoices && (
									<InvoiceList invoices={invoices} handleInvoiceIdxChange={handleInvoiceIdxChange} />
								)}
							</div>
						</div>
					)}
					{isMounted && (
						<div
							className={`tarnsition duration-500 ${
								status === "preEnter" || status === "exiting" ? "scale-75 transform opacity-0" : ""
							}`}
						>
							<div className="h-screen w-full bg-red-200"></div>
						</div>
					)}
				</section>
			</main>
		</div>
	);
}

export { App };
