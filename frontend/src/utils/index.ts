export function getShortID(id: string): string {
	let short = id.split("-")[0].toUpperCase();
	return `${short}...`;
}

export function getShortDate(date: Date): string {
	const d = new Date(date);
	return `Due ${d.getDate()} ${d.getMonth()} ${d.getFullYear()}`;
}
