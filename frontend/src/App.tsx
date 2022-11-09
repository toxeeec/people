import { useState } from "react";
import {
	createBrowserRouter,
	redirect,
	RouterProvider,
} from "react-router-dom";
import AuthContext from "./context/AuthContext";
import useAuth from "./hooks/useAuth";
import Layout from "./layout";
import Auth from "./pages/Auth";
import Home from "./pages/Home";

export default function App() {
	const { getAuth } = useAuth();
	const [authValues, setAuthValues] = useState(getAuth());
	const router = createBrowserRouter([
		{
			index: true,
			element: <Auth />,
			loader: () => {
				if (authValues.isAuthenticated) {
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
						if (!authValues.isAuthenticated) {
							return redirect("/");
						}
					},
				},
			],
		},
	]);
	return (
		<AuthContext.Provider value={{ authValues, setAuthValues }}>
			<RouterProvider router={router} />
		</AuthContext.Provider>
	);
}
