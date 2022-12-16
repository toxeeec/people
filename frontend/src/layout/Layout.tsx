import { useContext } from "react";
import { Outlet } from "react-router";
import { AuthContext } from "../context/AuthContext";
import { Footer } from "./Footer";
import { LayoutHeader } from "./LayoutHeader";

export const Layout = () => {
	const { isAuthenticated } = useContext(AuthContext);
	return (
		<>
			<LayoutHeader isAuthenticated={isAuthenticated} />
			{isAuthenticated ? null : <Footer />}
			<Outlet />
		</>
	);
};
