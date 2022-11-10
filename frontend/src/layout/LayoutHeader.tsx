import { Avatar, Header, Space, Flex, Drawer, Button } from "@mantine/core";
import { useState } from "react";
import useAuth from "../hooks/useAuth";

export default function LayoutHeader() {
	const [opened, setOpened] = useState(false);
	const { clearAuth } = useAuth();
	return (
		<>
			<Space h={60} />
			<Header height={60} fixed>
				<Flex px={0} h={60} align="center">
					<Avatar radius="xl" ml={11} onClick={() => setOpened(true)} />
				</Flex>
			</Header>
			<Drawer
				opened={opened}
				onClose={() => setOpened(false)}
				title="Account info"
				padding="md"
				size="md"
			>
				<Avatar size="lg" radius="xl" />
				<Button onClick={clearAuth} fullWidth radius="xl" mt="xl">
					Logout
				</Button>
			</Drawer>
		</>
	);
}
