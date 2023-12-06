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
			<main className="flex h-screen flex-col bg-off-white dark:bg-indigo-600 lg:flex-row">
				<SideBar toggleIsDark={toggleIsDark} isDark={isDark} />
				<section className="mx-auto mt-16 w-full max-w-[327px]  md:max-w-[672px] lg:max-w-2xl xl:max-w-4xl 2xl:max-w-6xl">
					<MenuBar />
					{invoices && <InvoiceList invoices={invoices} />}
				</section>
			</main>
		</div>
	);
}

export { App };
