import { useState } from "react"
import { SideBar } from "./components/SideBar/SideBar"

function App() {
    const [isDark, setIsDark] = useState(false)

    const toggleIsDark = () => {
        setIsDark((prev) => !prev)
    }

    return (
        <>
            <div className={`bg-gray-800 ${isDark ? "dark" : "light"}`}>
                <main className="bg-off-white dark:bg-indigo-600 h-screen">
                    <SideBar toggleIsDark={toggleIsDark} isDark={isDark} />
                </main>
            </div >
        </>
    )
}

export { App }
