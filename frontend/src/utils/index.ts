export function getShortID(id: string): string {
	let short = id.split("-")[0].toUpperCase();
	return `${short}...`;
}

export function getShortDate(date: string): string {
	const d = new Date(date);
	const months: { [key: number]: string } = {
		0: "Jan",
		1: "Feb",
		2: "Mar",
		3: "Apr",
		4: "May",
		5: "Jun",
		6: "Jul",
		7: "Aug",
		8: "Sep",
		9: "Oct",
		10: "Nov",
		11: "Dec",
	};
	return `${d.getDate()} ${months[d.getMonth()]} ${d.getFullYear()}`;
}
