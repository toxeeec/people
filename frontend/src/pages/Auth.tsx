import { Button, Center, Group } from "@mantine/core";
import { useState } from "react";
import Login from "../components/Login";
import Signup from "../components/Signup";

export default function Auth() {
	const [loginOpened, setLoginOpened] = useState(false);
	const [signupOpened, setSignupOpened] = useState(false);
	return (
		<Center w="100%" h="100%">
			<Group position="center">
				<Button
					onClick={() => setLoginOpened(true)}
					variant="outline"
					fullWidth
					radius="xl"
				>
					Log in
				</Button>
				<Button onClick={() => setSignupOpened(true)} fullWidth radius="xl">
					Sign up
				</Button>
			</Group>
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
		</Center>
	);
}
