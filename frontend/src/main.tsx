import React from "react";
import ReactDOM from "react-dom/client";
import { AuthContextProvider } from "@/context/AuthContext";
import "@/index.css";
import App from "@/App";

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
	<React.StrictMode>
		<AuthContextProvider>
			<App />
		</AuthContextProvider>
	</React.StrictMode>
);
