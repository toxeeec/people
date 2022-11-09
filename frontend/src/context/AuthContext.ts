import { createContext, Dispatch, SetStateAction } from "react";

export interface AuthValues {
	isAuthenticated: boolean;
	accessToken: string;
	refreshToken: string;
}

interface AuthContextType {
	authValues: AuthValues;
	setAuthValues: Dispatch<SetStateAction<AuthValues>>;
}

export const defaultAuthValues: AuthValues = {
	isAuthenticated: false,
	accessToken: "",
	refreshToken: "",
};

export default createContext<AuthContextType | null>(null);
