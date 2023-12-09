export function getShortID(id: string): string {
	let short = id.split("-")[0].toUpperCase();
	return `${short}...`;
}

export function getShortDate(date: Date): string {
	const d = new Date(date);
	const months = {
		1: "Jan",
		2: "Feb",
		3: "Mar",
		4: "Apr",
		5: "May",
		6: "Jun",
		7: "Jul",
		8: "Aug",
		9: "Sep",
		10: "Oct",
		11: "Nov",
		12: "Dec",
	};
	return `${d.getDate()} ${months[d.getMonth()]} ${d.getFullYear()}`;
}
