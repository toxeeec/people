import {
	Anchor,
	Button,
	LoadingOverlay,
	Modal,
	PasswordInput,
	TextInput,
	Text,
} from "@mantine/core";
import { useForm } from "@mantine/form";
import { Dispatch, SetStateAction, useContext } from "react";
import AuthContext from "../context/AuthContext";
import { AuthUser } from "../models";
import { usePostRegister } from "../spec.gen";

interface SignupProps {
	signupOpened: boolean;
	setSignupOpened: Dispatch<SetStateAction<boolean>>;
	setLoginOpened: Dispatch<SetStateAction<boolean>>;
}

export default function Signup({
	signupOpened,
	setSignupOpened,
	setLoginOpened,
}: SignupProps) {
	const form = useForm<AuthUser>({
		initialValues: { handle: "", password: "" },
		validate: {
			handle: (value) => (value.length < 5 ? "Invalid Handle" : null),
			password: (value) => (value.length < 12 ? "Invalid Password" : null),
		},
	});
	const { setAuth } = useContext(AuthContext)!;

	const { mutate, isLoading } = usePostRegister();
	function handleSubmit(values: AuthUser) {
		const data = { data: values };
		mutate(data, {
			onSuccess(data) {
				setAuth(data);
				setSignupOpened(false);
			},
			onError(error) {
				const err = error.message;
				if (err?.startsWith("Handle")) {
					form.setFieldError("handle", err);
				} else {
					form.setErrors({ handle: err, password: err });
				}
			},
		});
	}

	const handleLogin = () => {
		setSignupOpened(false);
		setLoginOpened(true);
	};
	return (
		<Modal
			centered
			title="Sign up"
			opened={signupOpened}
			onClose={() => setSignupOpened(false)}
		>
			<LoadingOverlay visible={isLoading} />
			<form onSubmit={form.onSubmit(handleSubmit)}>
				<TextInput
					withAsterisk
					label="Handle"
					{...form.getInputProps("handle")}
					mb="md"
				/>
				<PasswordInput
					withAsterisk
					label="Password"
					{...form.getInputProps("password")}
					mb="xl"
				/>
				<Button fullWidth radius="xl" type="submit" mb="xl">
					Submit
				</Button>
			</form>
			<Text>
				{"Already have an account? "}
				<Anchor onClick={() => handleLogin()}>Log in</Anchor>
			</Text>
		</Modal>
	);
}
