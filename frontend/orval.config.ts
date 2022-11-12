import { defineConfig } from "orval";

export default defineConfig({
	spec: {
		output: {
			target: "src/spec.gen.ts",
			schemas: "src/models",
			client: "react-query",
			override: {
				mutator: {
					path: "./src/custom-instance.ts",
					name: "customInstance",
				},
			},
		},
		input: {
			target: "./openapi.json",
		},
	},
});
