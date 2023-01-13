import { Tabs } from "@mantine/core";
import { DeleteAccount } from "../components/settings/DeleteAccount";
import { LogoutFromAll } from "../components/settings/LogoutFromAll";

const Settings = () => {
	return (
		<Tabs defaultValue="account" orientation="vertical" h="calc(100% - 60px)">
			<Tabs.List w="30%">
				<Tabs.Tab value="account" p="md">
					Account
				</Tabs.Tab>
			</Tabs.List>

			<Tabs.Panel value="account" p="md">
				<LogoutFromAll />
				<DeleteAccount />
			</Tabs.Panel>
		</Tabs>
	);
};

export default Settings;
