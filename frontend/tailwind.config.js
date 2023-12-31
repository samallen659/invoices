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
                800: "#0C0E16",
            },
            indigo: {
                200: "#7E88C3",
                400: "#373B53",
                600: "#141625",
                800: "#1E2139",
            },
            red: {
                200: "#FF9797",
                400: "#EC5757",
            },
            "off-white": "#F8F8FB",
            white: "#FFFFFF",
        },
    },
    plugins: [],
}
