import { ProfilePictureModal } from "@/components/images";
import { CenterLoader } from "@/components/utils";
import { AuthContext } from "@/context/AuthContext";
import { Footer } from "@/layout/Footer";
import { Header } from "@/layout/Header";
import { useGetUsersHandle } from "@/spec.gen";
import { useQueryClient } from "@tanstack/react-query";
import { useContext } from "react";
import { Outlet } from "react-router";

export function Layout() {
	const queryClient = useQueryClient();
	const { getAuth, isAuthenticated, isNewAccount, setIsNewAccount } = useContext(AuthContext);
	const { data: user, isLoading } = useGetUsersHandle(getAuth().handle, {
		query: { enabled: isAuthenticated },
	});
	return (
		<>
			{isAuthenticated && isLoading ? (
				<CenterLoader />
			) : (
				user && (
					<>
						<Header user={user} />
						<ProfilePictureModal
							user={user}
							opened={isNewAccount}
							handleChange={() => queryClient.resetQueries()}
							handleClose={() => setIsNewAccount(false)}
						/>
					</>
				)
			)}
			<Outlet />
			{!isAuthenticated && <Footer />}
		</>
	);
}
