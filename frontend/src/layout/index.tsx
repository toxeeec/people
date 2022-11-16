import { useContext } from "react";
import { Outlet } from "react-router";
import AuthContext from "../context/AuthContext";
import Footer from "./Footer";
import LayoutHeader from "./LayoutHeader";

export default function Layout() {
	const { isAuthenticated } = useContext(AuthContext)!;
	return (
		<>
			{isAuthenticated ? <LayoutHeader /> : null}
			{isAuthenticated ? null : <Footer />}
			<Outlet />
		</>
	);
}
