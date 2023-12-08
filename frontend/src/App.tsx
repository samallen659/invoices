import { useState } from "react";
import { SideBar } from "./components/SideBar/SideBar";
import { MenuBar } from "./components/MenuBar/MenuBar";
import { useQuery } from "react-query";
import { getAllInvoices } from "./api";
import { InvoiceList } from "./components/InvoiceList/InvoiceList";

function App() {
	const [isDark, setIsDark] = useState(false);

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
				<section className="mx-auto mt-24 w-full p-6 md:mt-36 md:max-w-[672px] md:p-0 lg:max-w-2xl xl:max-w-4xl 2xl:max-w-6xl">
					<MenuBar />
					<div className="mt-8 md:mt-14 lg:mt-16">{invoices && <InvoiceList invoices={invoices} />}</div>
				</section>
			</main>
		</div>
	);
}

export { App };
