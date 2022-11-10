import { Outlet } from "react-router";
import useAuth from "../hooks/useAuth";
import Footer from "./Footer";
import LayoutHeader from "./LayoutHeader";

export default function Layout() {
	const { auth } = useAuth();
	return (
		<>
			{auth.isAuthenticated ? <LayoutHeader /> : null}
			{auth.isAuthenticated ? null : <Footer />}
			<Outlet />
		</>
	);
}
