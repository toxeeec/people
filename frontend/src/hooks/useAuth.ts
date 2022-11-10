import { useLocalStorage } from "@mantine/hooks";
import { useContext } from "react";
import AuthContext, {
	AuthValues,
	defaultAuthValues,
} from "../context/AuthContext";
import { Tokens } from "../models";

export default function useAuth() {
	const ctx = useContext(AuthContext);
	const [accessToken, setAccessToken] = useLocalStorage({
		key: "accessToken",
		getInitialValueInEffect: false,
	});
	const [refreshToken, setRefreshToken] = useLocalStorage({
		key: "refreshToken",
		getInitialValueInEffect: false,
	});

	const setAuth = (tokens: Tokens) => {
		if (!ctx) {
			throw new Error("Context must be defined");
		}
		setAccessToken(tokens.accessToken);
		setRefreshToken(tokens.refreshToken);
		ctx.setAuthValues({
			...tokens,
			isAuthenticated: true,
		});
	};

	const getAuth = (): AuthValues => {
		const isAuthenticated = !!(accessToken && refreshToken);
		return { accessToken, refreshToken, isAuthenticated };
	};

	const clearAuth = () => {
		if (!ctx) {
			throw new Error("Context must be defined");
		}
		setAccessToken("");
		setRefreshToken("");
		ctx.setAuthValues({
			accessToken,
			refreshToken,
			isAuthenticated: false,
		});
	};

	return {
		auth: ctx?.authValues || defaultAuthValues,
		setAuth,
		getAuth,
		clearAuth,
	};
}
