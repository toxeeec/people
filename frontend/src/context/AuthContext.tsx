import { createContext, useState } from "react";
import { Tokens } from "../models";

export interface AuthValues {
	handle: string | null;
	accessToken: string | null;
	refreshToken: string | null;
}

export interface SetAuthProps {
	tokens?: Tokens;
	handle?: string;
}

interface AuthContextType {
	getAuth: () => AuthValues;
	setAuth: (props: SetAuthProps) => void;
	clearAuth: () => void;
	isAuthenticated: boolean;
}

export const AuthContext = createContext<AuthContextType>(null!);

interface AuthContextProviderProps {
	children: React.ReactNode;
}

export const AuthContextProvider = ({ children }: AuthContextProviderProps) => {
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

	return (
		<AuthContext.Provider
			value={{ getAuth, setAuth, clearAuth, isAuthenticated }}
		>
			{children}
		</AuthContext.Provider>
	);
};
