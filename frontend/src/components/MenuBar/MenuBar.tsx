import { useState } from "react";
import { Menu, RadioGroup } from "@headlessui/react";
import { CheckIcon } from "@heroicons/react/24/solid";

type Statuses = "Draft" | "Pending" | "Paid" | "";

function MenuBar() {
	let invoiceCount = 1;
	const [status, setStatus] = useState<Statuses>("");

	return (
		<div className="flex justify-between">
			<div>
				<h1 className="text-2xl font-bold text-gray-800 dark:text-white md:text-3xl">Invoices</h1>
				<p className="hidden text-sm text-gray-400 dark:text-gray-200 md:block">{`${
					invoiceCount == 0 ? "No invoices" : `There are ${invoiceCount} total invoices`
				}`}</p>
				<p className="text-sm text-gray-400 dark:text-gray-200 md:hidden">{`${
					invoiceCount == 0 ? "No invoices" : `${invoiceCount} invoices`
				}`}</p>
			</div>
			<div className="flex gap-10">
				<Menu as="div" className="relative self-center">
					<Menu.Button>
						<span className="z-0 hidden dark:text-white md:block">Filter by Status</span>
						<span className="text-xl dark:text-white md:hidden">Filter</span>
					</Menu.Button>
					<Menu.Items className="absolute left-0 right-0 ml-auto mt-4 h-32 w-48 translate-x-1/4 rounded-md bg-white shadow-lg">
						<RadioGroup
							value={status}
							onChange={setStatus}
							className="flex h-full w-full flex-col gap-2 p-6"
						>
							{["Draft", "Pending", "Paid"].map((status) => (
								<RadioGroup.Option value={status} key={status}>
									{({ checked }) => (
										<div className="group mx-auto flex gap-2">
											<div
												className={`h-4 w-4 rounded-sm group-hover:border-2 group-hover:border-purple-400 ${
													checked ? "border-2 border-purple-400 bg-purple-400" : "bg-gray-200"
												}`}
											>
												{checked && <CheckIcon className="text-white" />}
											</div>
											<span>{status}</span>
										</div>
									)}
								</RadioGroup.Option>
							))}
						</RadioGroup>
					</Menu.Items>
				</Menu>
				<button className="flex h-[48px] w-[90px] items-center gap-2 rounded-full bg-purple-400 p-2 hover:bg-purple-600 md:w-[150px] md:gap-4">
					<div className="flex h-8 w-8 items-center justify-center rounded-full bg-white">
						<PlusIcon />
					</div>
					<h3 className="hidden text-sm font-bold text-white md:block">New Invoice</h3>
					<h3 className="text-sm font-bold text-white md:hidden">New</h3>
				</button>
			</div>
		</div>
	);
}

function PlusIcon() {
	return (
		<svg width="11" height="11" xmlns="http://www.w3.org/2000/svg">
			<path
				d="M6.313 10.023v-3.71h3.71v-2.58h-3.71V.023h-2.58v3.71H.023v2.58h3.71v3.71z"
				fill="#7C5DFA"
				fillRule="nonzero"
			/>
		</svg>
	);
}

export { MenuBar };
