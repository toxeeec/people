import { MantineProvider } from "@mantine/core";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import UsersContext, { initialUsersContext } from "./context/UsersContext";
import "./index.css";

const queryClient = new QueryClient({
	defaultOptions: {
		queries: {
			refetchOnWindowFocus: false,
			refetchOnMount: false,
			refetchOnReconnect: false,
			retry: 1,
		},
	},
});

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
	<React.StrictMode>
		<MantineProvider
			withGlobalStyles
			withNormalizeCSS
			theme={{ colorScheme: "dark" }}
		>
			<QueryClientProvider client={queryClient}>
				<UsersContext.Provider value={initialUsersContext}>
					<App />
				</UsersContext.Provider>
			</QueryClientProvider>
		</MantineProvider>
	</React.StrictMode>
);
