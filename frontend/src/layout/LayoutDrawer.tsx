import { Drawer, Button } from "@mantine/core";
import { Dispatch, SetStateAction, useContext } from "react";
import { useNavigate } from "react-router-dom";
import { UserInfo } from "../components/UserInfo";
import { AuthContext } from "../context/AuthContext";
import { UsersContext } from "../context/UsersContext";
import { postLogout } from "../spec.gen";

interface LayoutDrawerProps {
	isOpened: boolean;
	setIsOpened: Dispatch<SetStateAction<boolean>>;
}

export const LayoutDrawer = ({ isOpened, setIsOpened }: LayoutDrawerProps) => {
	const navigate = useNavigate();
	const { getAuth, clearAuth } = useContext(AuthContext);
	const { users } = useContext(UsersContext);
	const user = users[getAuth().handle!];
	const handleLogout = () => {
		const { refreshToken } = getAuth();
		if (refreshToken) {
			postLogout({ refreshToken });
		}
		clearAuth();
	};
	return (
		<Drawer
			opened={isOpened}
			onClose={() => setIsOpened(false)}
			title="Account info"
			padding="md"
			size="md"
			zIndex={9999}
		>
			<UserInfo handle={user!.handle} />
			<Button
				fullWidth
				radius="xl"
				mt="xl"
				variant="outline"
				onClick={() => navigate("/settings")}
			>
				Settings
			</Button>
			<Button onClick={handleLogout} fullWidth radius="xl" mt="xl">
				Logout
			</Button>
		</Drawer>
	);
};
