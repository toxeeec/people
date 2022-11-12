import { useContext } from "react";
import { Outlet } from "react-router";
import AuthContext from "../context/AuthContext";
import Footer from "./Footer";
import LayoutHeader from "./LayoutHeader";

export default function Layout() {
	const { auth } = useContext(AuthContext)!;
	return (
		<>
			{auth.isAuthenticated ? <LayoutHeader /> : null}
			{auth.isAuthenticated ? null : <Footer />}
			<Outlet />
		</>
	);
}
