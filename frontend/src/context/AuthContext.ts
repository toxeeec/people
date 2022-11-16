import { createContext } from "react";
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
	setAuth: (_props: SetAuthProps) => void;
	clearAuth: () => void;
	isAuthenticated: boolean;
}

export default createContext<AuthContextType | null>(null);
