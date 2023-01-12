import { Tabs, Button, Tooltip } from "@mantine/core";
import { useContext } from "react";
import { AuthContext } from "../context/AuthContext";
import { postLogout } from "../spec.gen";

const Settings = () => {
	const { getAuth, clearAuth } = useContext(AuthContext);
	const handleLogout = () => {
		const { refreshToken } = getAuth();
		if (refreshToken) {
			postLogout({ refreshToken, logoutFromAll: true });
		}
		clearAuth();
	};
	return (
		<Tabs defaultValue="account" orientation="vertical" h="calc(100% - 60px)">
			<Tabs.List w="30%">
				<Tabs.Tab value="account" p="md">
					Account
				</Tabs.Tab>
			</Tabs.List>

			<Tabs.Panel value="account" p="md">
				<Tooltip
					label="Note that other sessions will still be logged in for up to 15 minutes"
					zIndex={9999}
				>
					<Button
						fullWidth
						variant="subtle"
						c="red"
						styles={{ inner: { justifyContent: "start" } }}
						onClick={handleLogout}
					>
						Log out of all sessions
					</Button>
				</Tooltip>
			</Tabs.Panel>
		</Tabs>
	);
};

export default Settings;
