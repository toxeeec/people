import { useState } from "react";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Footer from "./components/Footer";
import AuthContext from "./context/AuthContext";
import useAuth from "./hooks/useAuth";
import Home from "./pages/Home";

export default function App() {
	const router = createBrowserRouter([{ path: "/", element: <Home /> }]);
	const { getAuth } = useAuth();
	const [authValues, setAuthValues] = useState(getAuth());
	return (
		<AuthContext.Provider value={{ authValues, setAuthValues }}>
			<RouterProvider router={router} />
			{authValues.isAuthenticated || <Footer />}
		</AuthContext.Provider>
	);
}
