type SideBarProps = {
    toggleIsDark: () => void;
    isDark: boolean;
}

function SideBar({ toggleIsDark, isDark }: SideBarProps) {
    return <div className="flex flex-row bg-indigo-800 lg:flex-col h-20 w-screen lg:h-screen lg:w-[103px] lg:rounded-r-[20px]">

    </div>
}

export { SideBar }
