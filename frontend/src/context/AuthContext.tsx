import { createContext, type Dispatch, type SetStateAction, useState } from "react";
import { type Tokens } from "@/models";

export type AuthValues = {
	handle: string;
	accessToken: string;
	refreshToken: string;
};

export type SetAuthProps = {
	tokens?: Tokens;
	handle?: string;
};

type AuthContextType = {
	getAuth: () => AuthValues;
	setAuth: (props: SetAuthProps) => void;
	clearAuth: () => void;
	isAuthenticated: boolean;
	isNewAccount: boolean;
	setIsNewAccount: Dispatch<SetStateAction<boolean>>;
};

export const AuthContext = createContext<AuthContextType>({} as AuthContextType);

export function AuthContextProvider({ children }: { children: React.ReactNode }) {
	const getAuth = () => {
		const accessToken = localStorage.getItem("accessToken") ?? "";
		const refreshToken = localStorage.getItem("refreshToken") ?? "";
		const handle = localStorage.getItem("handle") ?? "";
		return { accessToken, refreshToken, handle };
	};

	const [isAuthenticated, setIsAuthenticated] = useState(!!getAuth().handle);

	const setAuth = ({ tokens, handle }: SetAuthProps) => {
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

	const [isNewAccount, setIsNewAccount] = useState(false);

	return (
		<AuthContext.Provider
			value={{
				getAuth,
				setAuth,
				clearAuth,
				isAuthenticated,
				isNewAccount,
				setIsNewAccount,
			}}
		>
			{children}
		</AuthContext.Provider>
	);
}
