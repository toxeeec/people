import { useState } from "react";
import { Tokens } from "../models";

export const useAuth = () => {
	const getAuth = () => {
		const accessToken = localStorage.getItem("accessToken");
		const refreshToken = localStorage.getItem("refreshToken");
		const handle = localStorage.getItem("handle");
		return { accessToken, refreshToken, handle };
	};

	const [isAuthenticated, setIsAuthenticated] = useState(!!getAuth().handle);

	const setAuth = ({
		tokens,
		handle,
	}: {
		tokens?: Tokens;
		handle?: string;
	}) => {
		if (tokens) {
			localStorage.setItem("accessToken", tokens.accessToken);
			localStorage.setItem("refreshToken", tokens.refreshToken);
		}
		if (handle) {
			localStorage.setItem("handle", handle);
		}
		setIsAuthenticated(true);
	};

	const clearAuth = () => {
		localStorage.removeItem("accessToken");
		localStorage.removeItem("refreshToken");
		localStorage.removeItem("handle");
		setIsAuthenticated(false);
	};

	return {
		getAuth,
		setAuth,
		clearAuth,
		isAuthenticated,
	};
};
