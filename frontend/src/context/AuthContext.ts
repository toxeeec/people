import { createContext } from "react";
import { Tokens, User } from "../models";

export interface AuthValues {
	user?: User;
	isAuthenticated: boolean;
	accessToken: string;
	refreshToken: string;
}

export interface SetAuthProps {
	tokens?: Tokens;
	user?: User;
}

interface AuthContextType {
	auth: AuthValues;
	setAuth: (_props: SetAuthProps) => void;
	clearAuth: () => void;
}

export default createContext<AuthContextType | null>(null);
