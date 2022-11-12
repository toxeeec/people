import { useLocalStorage } from "@mantine/hooks";
import { useMemo } from "react";
import { AuthValues } from "../context/AuthContext";
import { Tokens, User } from "../models";

export default function useAuth() {
	const [accessToken, setAccessToken] = useLocalStorage({
		key: "accessToken",
		defaultValue: "",
		getInitialValueInEffect: false,
	});
	const [refreshToken, setRefreshToken] = useLocalStorage({
		key: "refreshToken",
		defaultValue: "",
		getInitialValueInEffect: false,
	});
	const [userString, setUserString] = useLocalStorage({
		key: "user",
		defaultValue: "",
		getInitialValueInEffect: false,
	});

	const auth: AuthValues = useMemo(() => {
		const isAuthenticated = !!(accessToken && refreshToken);
		const user = userString ? JSON.parse(userString) : undefined;
		return { accessToken, refreshToken, isAuthenticated, user };
	}, [accessToken, refreshToken, userString]);

	const setAuth = ({ tokens, user }: { tokens?: Tokens; user?: User }) => {
		if (tokens) {
			setAccessToken(tokens.accessToken);
			setRefreshToken(tokens.refreshToken);
		}
		if (user) {
			setUserString(JSON.stringify(user));
		}
	};

	const clearAuth = () => {
		setAccessToken("");
		setRefreshToken("");
		setUserString("");
	};

	return {
		auth,
		setAuth,
		clearAuth,
	};
}
