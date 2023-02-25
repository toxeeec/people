import { useContext } from "react";
import { AuthContext } from "@/context/AuthContext";
import { usePostLogout } from "@/spec.gen";
import { DangerButton } from "@/components/buttons";

export function LogoutFromAll() {
	const { getAuth, clearAuth } = useContext(AuthContext);
	const { mutate } = usePostLogout({
		mutation: { retry: 1, onSuccess: clearAuth },
	});
	const handleLogout = () => {
		const { refreshToken } = getAuth();
		if (refreshToken) {
			mutate({ data: { refreshToken, logoutFromAll: true } });
		}
	};
	return (
		<DangerButton
			label="Note that other sessions will still be logged in for up to 15 minutes"
			onClick={handleLogout}
			text="Log out of all sessions"
		/>
	);
}
