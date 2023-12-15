import { useState, useEffect, useRef } from "react";
import { CheckIcon } from "@heroicons/react/24/solid";
import { IconUpArrow, IconDownArrow, IconPlus } from "../Icons";

type Statuses = "Draft" | "Pending" | "Paid" | "";

type MenuBarProps = {
	setFormState: (v: "new" | "edit") => void;
	toggle: (t: boolean) => void;
};

function MenuBar({ setFormState, toggle }: MenuBarProps) {
	let invoiceCount = 1;
	const [status, setStatus] = useState<Statuses>("");
	const [filter, setFilter] = useState<string>("");
	const [showFilterMenu, setShowFilterMenu] = useState<boolean>(false);

	const handleClick = () => {
		setFormState("new");
		toggle(true);
	};

	const handleFilterMenuClick = (s: string) => {
		if (s === filter) {
			setFilter("");
			return;
		}
		setFilter(s);
	};

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
			<div className="flex gap-12">
				<FilterMenu
					filter={filter}
					showFilterMenu={showFilterMenu}
					setShowFilterMenu={setShowFilterMenu}
					handleFilterMenuClick={handleFilterMenuClick}
				/>
				<button
					onClick={handleClick}
					className="flex h-[48px] w-[90px] items-center gap-2 rounded-full bg-purple-400 p-2 hover:bg-purple-600 md:w-[150px] md:gap-4"
				>
					<div className="flex h-8 w-8 items-center justify-center rounded-full bg-white">
						<IconPlus />
					</div>
					<h3 className="hidden text-sm font-bold text-white md:block">New Invoice</h3>
					<h3 className="text-sm font-bold text-white md:hidden">New</h3>
				</button>
			</div>
		</div>
	);
}

type FilterMenuProps = {
	filter: string;
	showFilterMenu: boolean;
	setShowFilterMenu: (b: boolean) => void;
	handleFilterMenuClick: (s: string) => void;
};

function FilterMenu({ filter, showFilterMenu, setShowFilterMenu, handleFilterMenuClick }: FilterMenuProps) {
	const filterRef = useRef<HTMLDivElement>(null);

	useEffect(() => {
		const handler = (e: any) => {
			if (filterRef) {
				if (!filterRef.current?.contains(e.target)) {
					setShowFilterMenu(false);
				}
			}
		};

		document.addEventListener("mousedown", handler);
	});

	return (
		<div className="flex gap-10" ref={filterRef}>
			<div className="relative self-center">
				<button
					className="z-0 hidden items-center gap-2 dark:text-white md:flex"
					onClick={() => setShowFilterMenu(!showFilterMenu)}
				>
					Filter by Status
					{showFilterMenu ? (
						<div className="-mt-1">
							<IconUpArrow />
						</div>
					) : (
						<div className="-mt-1">
							<IconDownArrow />
						</div>
					)}
				</button>
				<button
					className="flex items-center gap-2 text-xl dark:text-white md:hidden"
					onClick={() => setShowFilterMenu(!showFilterMenu)}
				>
					Filter
					{showFilterMenu ? (
						<div className="-mt-1">
							<IconUpArrow />
						</div>
					) : (
						<div className="-mt-1">
							<IconDownArrow />
						</div>
					)}
				</button>
				{showFilterMenu && (
					<div className="absolute left-0 right-0 ml-auto mt-4 h-32 w-48 translate-x-1/4 rounded-md bg-white shadow-lg dark:bg-gray-600">
						<div className="flex h-full w-full flex-col items-start gap-2 p-6">
							{["Draft", "Pending", "Paid"].map((status) => (
								<div
									className="group flex w-full gap-2"
									onClick={() => handleFilterMenuClick(status)}
									key={status}
								>
									<div
										className={`h-4 w-4 rounded-sm group-hover:border-2 group-hover:border-purple-400 ${
											filter === status
												? "border-2 border-purple-400 bg-purple-400"
												: "bg-gray-200 dark:bg-indigo-800"
										}`}
									>
										{filter === status && <CheckIcon className="text-white" />}
									</div>
									<span className="dark:text-white">{status}</span>
								</div>
							))}
						</div>
					</div>
				)}
			</div>
		</div>
	);
}

export { MenuBar };
