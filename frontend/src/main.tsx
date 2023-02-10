import { MantineProvider } from "@mantine/core";
import { NotificationsProvider } from "@mantine/notifications";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import { AuthContextProvider } from "./context/AuthContext";
import { NotificationsContextProvider } from "./context/NotificationsContext";
import "./index.css";

const queryClient = new QueryClient({
	defaultOptions: {
		queries: {
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
			<NotificationsProvider position="bottom-center">
				<QueryClientProvider client={queryClient}>
					<AuthContextProvider>
						<NotificationsContextProvider>
							<App />
						</NotificationsContextProvider>
					</AuthContextProvider>
				</QueryClientProvider>
			</NotificationsProvider>
		</MantineProvider>
	</React.StrictMode>
);
