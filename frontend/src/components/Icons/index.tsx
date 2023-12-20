export function IconDownArrow() {
	return (
		<svg width="11" height="7" xmlns="http://www.w3.org/2000/svg">
			<path d="M1 1l4.228 4.228L9.456 1" stroke="#7C5DFA" strokeWidth="2" fill="none" fillRule="evenodd" />
		</svg>
	);
}

export function IconUpArrow() {
	return (
		<svg width="11" height="7" xmlns="http://www.w3.org/2000/svg" className="rotate-180">
			<path d="M1 1l4.228 4.228L9.456 1" stroke="#7C5DFA" strokeWidth="2" fill="none" fillRule="evenodd" />
		</svg>
	);
}

export function IconPlus() {
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

export function IconLeftArrow() {
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

export function IconArrowRight() {
	return (
		<svg width="7" height="10" xmlns="http://www.w3.org/2000/svg">
			<path d="M1 1l4 4-4 4" stroke="#7C5DFA" strokeWidth="2" fill="none" fillRule="evenodd" />
		</svg>
	);
}

export function IconDelete() {
	return (
		<svg width="13" height="16" xmlns="http://www.w3.org/2000/svg">
			<path
				d="M11.583 3.556v10.666c0 .982-.795 1.778-1.777 1.778H2.694a1.777 1.777 0 01-1.777-1.778V3.556h10.666zM8.473 0l.888.889h3.111v1.778H.028V.889h3.11L4.029 0h4.444z"
				fill="#888EB0"
				fillRule="nonzero"
			/>
		</svg>
	);
}

export function IconSpinning() {
	return (
		<div role="status">
			<svg
				aria-hidden="true"
				className="fill-blue-600 h-8 w-8 animate-spin text-gray-200 dark:text-gray-600"
				viewBox="0 0 100 101"
				fill="none"
				xmlns="http://www.w3.org/2000/svg"
			>
				<path
					d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
					fill="currentColor"
				/>
				<path
					d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
					fill="currentFill"
				/>
			</svg>
			<span className="sr-only">Loading...</span>
		</div>
	);
}
