import Logo from "../../assets/logo.svg";
import SunLogo from "../../assets/icon-sun.svg";
import MoonLogo from "../../assets/icon-moon.svg";

type SideBarProps = {
	toggleIsDark: () => void;
	isDark: boolean;
};

function SideBar({ toggleIsDark, isDark }: SideBarProps) {
	return (
		<div className="fixed z-50 flex h-20 w-screen flex-row gap-5 bg-indigo-800 lg:h-screen lg:w-[103px] lg:flex-col lg:rounded-r-[20px]">
			<div className="relative h-20 w-20 flex-none rounded-r-[20px] bg-purple-400 lg:h-[103px] lg:w-[103px]">
				<img
					className="absolute bottom-0 left-0 right-0 top-0 z-20 m-auto h-6 w-6 lg:h-8 lg:w-8"
					src={Logo}
					alt="Invoices"
				/>
				<div className="absolute bottom-0 z-0 h-10 w-20 rounded-br-[20px] rounded-tl-[20px] bg-purple-600 lg:h-[51px] lg:w-[103px]"></div>
			</div>
			<div className="flex flex-auto justify-end lg:mx-auto lg:flex-col">
				<button onClick={toggleIsDark}>
					<img src={`${isDark ? SunLogo : MoonLogo}`} alt={`${isDark ? "Light Mode" : "Dark Mode"}`} />
				</button>
			</div>
			<div className="h-20 w-20 flex-none border-l-2 border-indigo-400 lg:h-[103px] lg:w-[103px] lg:rounded-br-[20px] lg:border-l-0 lg:border-t-2"></div>
		</div>
	);
}

export { SideBar };
