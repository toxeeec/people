import { useQueryClient } from "@tanstack/react-query";
import { useContext } from "react";
import { Outlet } from "react-router";
import { CenterLoader } from "../components/CenterLoader";
import { ProfilePictureModal } from "../components/images/ProfilePictureModal";
import { AuthContext } from "../context/AuthContext";
import { useGetUsersHandle } from "../spec.gen";
import { Footer } from "./Footer";
import { LayoutHeader } from "./LayoutHeader";

export const Layout = () => {
	const queryClient = useQueryClient();
	const { getAuth, isAuthenticated, isNewAccount, setIsNewAccount } =
		useContext(AuthContext);
	const { data: user, isLoading } = useGetUsersHandle(getAuth().handle!, {
		query: { enabled: isAuthenticated },
	});
	return isLoading || !user ? (
		<CenterLoader />
	) : (
		<>
			<LayoutHeader user={user} isAuthenticated={isAuthenticated} />
			{isAuthenticated ? null : <Footer />}
			<ProfilePictureModal
				user={user}
				opened={isNewAccount}
				onClose={() => {
					queryClient.resetQueries();
					setIsNewAccount(false);
				}}
			/>
			<Outlet />
		</>
	);
};
