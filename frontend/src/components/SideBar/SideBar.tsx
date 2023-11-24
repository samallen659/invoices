import Logo from "../../assets/logo.svg"

type SideBarProps = {
    toggleIsDark: () => void;
    isDark: boolean;
}

function SideBar({ toggleIsDark, isDark }: SideBarProps) {
    return <div className="flex flex-row bg-indigo-800 lg:flex-col h-20 w-screen lg:h-screen lg:w-[103px] lg:rounded-r-[20px] gap-5">
        <div className="flex-none w-20 h-20 lg:w-[103px] lg:h-[103px] bg-purple-400 rounded-r-[20px] relative">
            <img className="absolute z-20 w-6 h-6 lg:w-8 lg:h-8 m-auto left-0 right-0 top-0 bottom-0" src={Logo} alt="Invoices" />
            <div className="absolute z-0 bottom-0 bg-purple-600 w-20 h-10 lg:w-[103px] lg:h-[51px] rounded-tl-[20px] rounded-br-[20px]"></div>
        </div>
        <div className="flex-auto bg-red-200"></div>
        <div className="flex-none bg-red-200 w-20 h-20 lg:w-[103px] lg:h-[103px] lg:rounded-br-[20px]"></div>
    </div>
}

export { SideBar }
