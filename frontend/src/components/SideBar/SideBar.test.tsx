import { describe, test, expect } from "vitest";
import { render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { SideBar } from "./SideBar";

describe("<SideBar />", () => {
	let isDark = false;
	let toggleIsDarkStubCalls: number;
	const toggleIsDarkStub = () => {
		toggleIsDarkStubCalls++;
	};

	test("SideBar mounts properly", () => {
		const wrapper = render(<SideBar isDark={isDark} toggleIsDark={toggleIsDarkStub} />);
		expect(wrapper).toBeTruthy();
		wrapper.unmount();
	});
	test("Clicking dark mode button calls toggleIsDark function", async () => {
		render(<SideBar isDark={isDark} toggleIsDark={toggleIsDarkStub} />);
		const user = userEvent.setup();

		toggleIsDarkStubCalls = 0;

		const button = screen.getByRole("button");
		console.log(button);
		await user.click(button);

		expect(toggleIsDarkStubCalls).toBe(1);
	});
});
