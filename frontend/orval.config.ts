import { defineConfig } from "orval";

export default defineConfig({
	spec: {
		output: {
			target: "src/spec.gen.ts",
			schemas: "src/models",
			client: "react-query",
		},
		input: {
			target: "./openapi.yaml",
		},
	},
});
