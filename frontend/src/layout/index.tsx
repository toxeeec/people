import { Outlet } from "react-router";
import useAuth from "../hooks/useAuth";
import Footer from "./Footer";

export default function Layout() {
	const { auth } = useAuth();
	return (
		<>
			{auth.isAuthenticated ? null : <Footer />}
			<Outlet />
		</>
	);
}
