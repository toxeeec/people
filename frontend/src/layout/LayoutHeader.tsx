import {
	Avatar,
	Header,
	Space,
	Drawer,
	Button,
	Group,
	Text,
} from "@mantine/core";
import { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import useAuth from "../hooks/useAuth";

export default function LayoutHeader() {
	const [opened, setOpened] = useState(false);
	const { clearAuth, auth } = useAuth();
	const location = useLocation();
	const [route, setRoute] = useState("");
	useEffect(() => {
		switch (location.pathname) {
			case "/home":
				setRoute("Home");
		}
	}, [location]);

	return (
		<>
			<Space h={60} />
			<Header height={60} fixed>
				<Group h={60} align="center">
					<Avatar radius="xl" ml={11} onClick={() => setOpened(true)} />
					<Text fz="xl" fw={700}>
						{route}
					</Text>
				</Group>
			</Header>
			<Drawer
				opened={opened}
				onClose={() => setOpened(false)}
				title="Account info"
				padding="md"
				size="md"
			>
				<Avatar size="lg" radius="xl" mb="xs" />
				<b>@{auth.user?.handle}</b>
				<Group mt="xs">
					<span>
						<b>{auth.user?.following}</b> Following
					</span>
					<span>
						<b>{auth.user?.followers}</b> Followers
					</span>
				</Group>
				<Button onClick={clearAuth} fullWidth radius="xl" mt="xl">
					Logout
				</Button>
			</Drawer>
		</>
	);
}
