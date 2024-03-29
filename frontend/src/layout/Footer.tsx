import { Button, Divider, Group, Paper } from "@mantine/core";
import { useState } from "react";
import { Login } from "@/components/auth/Login";
import { Signup } from "@/components/auth/Signup";

export function Footer() {
	const [loginOpened, setLoginOpened] = useState(false);
	const [signupOpened, setSignupOpened] = useState(false);
	return (
		<>
			<Paper pb="md" bottom={0} pos="fixed" w="100%" style={{ zIndex: 1 }}>
				<Divider mb="md" />
				<Group position="center" grow px="md">
					<Button onClick={() => setLoginOpened(true)} variant="outline" fullWidth radius="xl">
						Log in
					</Button>
					<Button onClick={() => setSignupOpened(true)} fullWidth radius="xl">
						Sign up
					</Button>
				</Group>
			</Paper>
			<Login
				loginOpened={loginOpened}
				setLoginOpened={setLoginOpened}
				setSignupOpened={setSignupOpened}
			/>
			<Signup
				signupOpened={signupOpened}
				setSignupOpened={setSignupOpened}
				setLoginOpened={setLoginOpened}
			/>
		</>
	);
}
