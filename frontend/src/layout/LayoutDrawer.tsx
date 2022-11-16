import { Drawer, Button } from "@mantine/core";
import { Dispatch, SetStateAction, useContext } from "react";
import AccountInfo from "../components/AccountInfo";
import AuthContext from "../context/AuthContext";
import UsersContext from "../context/UsersContext";

interface LayoutDrawerProps {
	isOpened: boolean;
	setIsOpened: Dispatch<SetStateAction<boolean>>;
}

export default function LayoutDrawer({
	isOpened,
	setIsOpened,
}: LayoutDrawerProps) {
	const { getAuth, clearAuth } = useContext(AuthContext)!;
	const usersCtx = useContext(UsersContext);
	const user = usersCtx!.users.get(getAuth().handle!)!;
	return (
		<Drawer
			opened={isOpened}
			onClose={() => setIsOpened(false)}
			title="Account info"
			padding="md"
			size="md"
		>
			<AccountInfo user={user} />
			<Button onClick={clearAuth} fullWidth radius="xl" mt="xl">
				Logout
			</Button>
		</Drawer>
	);
}
