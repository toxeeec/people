import { useEffect } from "react";
import {
	createBrowserRouter,
	redirect,
	RouterProvider,
} from "react-router-dom";
import AuthContext from "./context/AuthContext";
import UsersContext, { initialUsersContext } from "./context/UsersContext";
import {
	AXIOS_INSTANCE,
	createRequestInterceptor,
	createResponseInterceptor,
} from "./custom-instance";
import useAuth from "./hooks/useAuth";
import Layout from "./layout";
import Auth from "./pages/Auth";
import Home from "./pages/Home";

export default function App() {
	const { auth, setAuth, clearAuth } = useAuth();

	useEffect(() => {
		const requestInterceptor = createRequestInterceptor(auth.accessToken);
		const responseInterceptor = createResponseInterceptor(
			auth.refreshToken,
			setAuth,
			clearAuth
		);

		return () => {
			AXIOS_INSTANCE.interceptors.request.eject(requestInterceptor);
			AXIOS_INSTANCE.interceptors.response.eject(responseInterceptor);
		};
	}, [clearAuth, auth, setAuth]);
	const router = createBrowserRouter([
		{
			index: true,
			element: <Auth />,
			loader: () => {
				if (auth.isAuthenticated) {
					return redirect("/home");
				}
			},
		},
		{
			element: <Layout />,
			children: [
				{
					path: "/home",
					element: <Home />,
					loader: () => {
						if (!auth.isAuthenticated) {
							return redirect("/");
						}
					},
				},
			],
		},
	]);
	return (
		<UsersContext.Provider value={initialUsersContext}>
			<AuthContext.Provider value={{ auth, setAuth, clearAuth }}>
				<RouterProvider router={router} />
			</AuthContext.Provider>
		</UsersContext.Provider>
	);
}
