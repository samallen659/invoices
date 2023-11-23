/** @type {import('tailwindcss').Config} */
export default {
    content: [
        "./index.html",
        "./src/**/*.{js,ts,jsx,tsx}",
    ],
    darkMode: 'class',
    theme: {
        extend: {},
        fontFamily: {
            sans: ["League Spartan", "sans-serif"],
        },
        colors: {
            purple: {
                400: "#7C5DFA",
                600: "#9277FF",
            },
            gray: {
                200: "#DFE3FA",
                400: "#888EB0",
                600: "#252945",
                800: "#1E2139",
            },
            indigo: {
                400: "7E88C3",
                800: "#141625",
            },
            red: {
                200: "#9277FF",
                400: "#EC5757",
            },
            "off-white": "#F8F8FB",
            white: "#FFFFFF",
        }
    },
    plugins: [],
}
