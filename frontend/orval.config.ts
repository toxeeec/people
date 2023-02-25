import { defineConfig } from "orval";

export default defineConfig({
	spec: {
		output: {
			target: "src/spec.gen.ts",
			schemas: "src/models",
			client: "react-query",
			prettier: true,
			override: {
				query: {
					options: {
						refetchOnWindowFocus: false,
						refetchOnMount: false,
						refetchOnReconnect: false,
						retry: 1,
					},
				},
				mutator: {
					path: "./src/custom-instance.ts",
					name: "customInstance",
				},
			},
		},
		input: {
			target: "./openapi.json",
		},
		hooks: {
			afterAllFilesWrite: "eslint src/models/*.ts src/spec.gen.ts --fix",
		},
	},
});
