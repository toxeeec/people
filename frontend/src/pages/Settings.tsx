import { DeleteAccount, LogoutFromAll } from "@/components/settings";
import { RouteContext } from "@/context/RouteContext";
import { HEADER_HEIGHT } from "@/layout/Header";
import { Tabs } from "@mantine/core";
import { useContext, useEffect } from "react";

function Settings() {
	const { setRouteName } = useContext(RouteContext);
	useEffect(() => {
		setRouteName("Settings");
	}, [setRouteName]);

	return (
		<Tabs defaultValue="account" orientation="vertical" h={`calc(100% - ${HEADER_HEIGHT}px)`}>
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
}

export default Settings;
