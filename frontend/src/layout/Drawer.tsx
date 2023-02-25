import { Drawer as MantineDrawer, Button } from "@mantine/core";
import { type Dispatch, type SetStateAction, useContext } from "react";
import { useNavigate } from "react-router-dom";
import { UserInfo } from "@/components/user";
import { AuthContext } from "@/context/AuthContext";
import { type User } from "@/models";
import { postLogout } from "@/spec.gen";

type LayoutDrawerProps = {
	isOpened: boolean;
	setIsOpened: Dispatch<SetStateAction<boolean>>;
	user: User;
};

export function Drawer({ isOpened, setIsOpened, user }: LayoutDrawerProps) {
	const navigate = useNavigate();
	const { getAuth, clearAuth } = useContext(AuthContext);

	const handleLogout = () => {
		const { refreshToken } = getAuth();
		postLogout({ refreshToken });
		clearAuth();
	};

	return (
		<MantineDrawer
			opened={isOpened}
			onClose={() => setIsOpened(false)}
			title="Account info"
			padding="md"
			size="md"
			zIndex={9999}
		>
			<UserInfo user={user} />
			<Button onClick={() => navigate("/settings")} fullWidth radius="xl" mt="xl" variant="outline">
				Settings
			</Button>
			<Button onClick={handleLogout} fullWidth radius="xl" mt="xl">
				Logout
			</Button>
		</MantineDrawer>
	);
}
