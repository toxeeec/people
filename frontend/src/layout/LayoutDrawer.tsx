import { Drawer, Button } from "@mantine/core";
import { Dispatch, SetStateAction, useContext } from "react";
import { AccountInfo } from "../components/AccountInfo";
import { AuthContext } from "../context/AuthContext";
import { UsersContext } from "../context/UsersContext";

interface LayoutDrawerProps {
	isOpened: boolean;
	setIsOpened: Dispatch<SetStateAction<boolean>>;
}

export const LayoutDrawer = ({ isOpened, setIsOpened }: LayoutDrawerProps) => {
	const { getAuth, clearAuth } = useContext(AuthContext);
	const { users } = useContext(UsersContext);
	const user = users[getAuth().handle!];
	return (
		<Drawer
			opened={isOpened}
			onClose={() => setIsOpened(false)}
			title="Account info"
			padding="md"
			size="md"
		>
			<AccountInfo handle={user!.handle} />
			<Button onClick={clearAuth} fullWidth radius="xl" mt="xl">
				Logout
			</Button>
		</Drawer>
	);
};
