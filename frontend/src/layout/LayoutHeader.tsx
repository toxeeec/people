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
import { useLocation, useParams } from "react-router";
import AccountInfo from "../components/AccountInfo";
import AuthContext from "../context/AuthContext";
import UsersContext from "../context/UsersContext";

export default function LayoutHeader() {
	const [isOpened, setIsOpened] = useState(false);
	const { getAuth, clearAuth } = useContext(AuthContext)!;
	const usersCtx = useContext(UsersContext);
	const user = usersCtx!.users.get(getAuth().handle!)!;
	const params = useParams();
	const location = useLocation();
	const [route, setRoute] = useState("");
	useEffect(() => {
		if (location.pathname === "/home") {
			setRoute("Home");
		}
	}, [params, location]);

	return (
		<>
			<Space h={60} />
			<Header height={60} fixed>
				<Group h={60} align="center">
					<UnstyledButton onClick={() => setIsOpened(true)} ml={11}>
						<Avatar radius="xl" onClick={() => setIsOpened(true)} />
					</UnstyledButton>
					<Text fz="xl" fw={700}>
						{route}
					</Text>
				</Group>
			</Header>
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
		</>
	);
}
