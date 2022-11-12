import { MantineProvider } from "@mantine/core";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import "./index.css";

const queryClient = new QueryClient({
	defaultOptions: { queries: { retry: false } },
});

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
	<React.StrictMode>
		<QueryClientProvider client={queryClient}>
			<MantineProvider
				withGlobalStyles
				withNormalizeCSS
				theme={{ colorScheme: "dark" }}
			>
				<App />
			</MantineProvider>
		</QueryClientProvider>
	</React.StrictMode>
);
