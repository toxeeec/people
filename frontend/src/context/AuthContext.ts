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
	setAuth: (props: SetAuthProps) => void;
	clearAuth: () => void;
	isAuthenticated: boolean;
}

export const AuthContext = createContext<AuthContextType>(null!);
