import { defineConfig, loadEnv } from "vite";
import { resolve } from "path";
import react from "@vitejs/plugin-react";

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => {
	const env = loadEnv(mode, process.cwd(), "");
	return {
		plugins: [react()],
		build: {
			rollupOptions: {
				input: {
					main: resolve(__dirname, "index.html"),
					swagger: resolve(__dirname, "swagger.html"),
				},
				output: {
					manualChunks(id) {
						if (id.includes("node_modules")) {
							return id
								.toString()
								.split("node_modules/")[1]
								.split("/")[0]
								.toString();
						}
					},
				},
			},
		},
		define: {
			BACKEND_PORT: env.BACKEND_PORT || "8000",
		},
		server: {
			port: parseInt(env.FRONTEND_PORT) || 80,
		},
	};
});
