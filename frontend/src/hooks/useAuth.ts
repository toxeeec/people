import { useContext, useState } from "react";
import UsersContext from "../context/UsersContext";
import { Tokens } from "../models";
import { useGetUsersHandle } from "../spec.gen";

export default function useAuth() {
	const getAuth = () => {
		const accessToken = localStorage.getItem("accessToken");
		const refreshToken = localStorage.getItem("refreshToken");
		const handle = localStorage.getItem("handle");
		return { accessToken, refreshToken, handle };
	};

	const [isAuthenticated, setIsAuthenticated] = useState(!!getAuth().handle);
	const usersCtx = useContext(UsersContext);

	const { data: user } = useGetUsersHandle(getAuth().handle!, {
		query: { enabled: isAuthenticated },
	});

	if (user) {
		usersCtx?.setUser(user!.handle!, user!);
	}

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
}
