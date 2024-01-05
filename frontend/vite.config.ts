/// <reference types="vitest" />
import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";

// https://vitejs.dev/config/
export default defineConfig({
	server: {
		proxy: {
			"/invoice": "http://localhost:8080",
		},
	},
	plugins: [react()],
	test: {
		environment: "jsdom",
	},
	preview: {
		host: true,
		port: 5173,
	},
});
