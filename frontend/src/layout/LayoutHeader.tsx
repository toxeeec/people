import {
	Avatar,
	Header,
	Space,
	Drawer,
	Button,
	Group,
	Text,
	UnstyledButton,
} from "@mantine/core";
import { useContext, useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import AccountInfo from "../components/AccountInfo";
import UsersContext from "../context/UsersContext";
import useAuth from "../hooks/useAuth";

export default function LayoutHeader() {
	const [opened, setOpened] = useState(false);
	const { clearAuth, auth } = useAuth();
	const usersCtx = useContext(UsersContext);
	usersCtx?.setUser(auth.user!.handle, auth.user!);

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
					<UnstyledButton onClick={() => setOpened(true)} ml={11}>
						<Avatar radius="xl" />
					</UnstyledButton>
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
				<AccountInfo handle={auth.user!.handle} />
				<Button onClick={clearAuth} fullWidth radius="xl" mt="xl">
					Logout
				</Button>
			</Drawer>
		</>
	);
}
