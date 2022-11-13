import {
	Anchor,
	Button,
	LoadingOverlay,
	Modal,
	PasswordInput,
	TextInput,
} from "@mantine/core";
import { useForm } from "@mantine/form";
import { Text } from "@mantine/core";
import { Dispatch, SetStateAction, useContext } from "react";
import { AuthUser } from "../models";
import { usePostLogin } from "../spec.gen";
import AuthContext from "../context/AuthContext";

interface LoginProps {
	loginOpened: boolean;
	setLoginOpened: Dispatch<SetStateAction<boolean>>;
	setSignupOpened: Dispatch<SetStateAction<boolean>>;
}

export default function Login({
	loginOpened,
	setLoginOpened,
	setSignupOpened,
}: LoginProps) {
	const form = useForm<AuthUser>({
		initialValues: { handle: "", password: "" },
		validate: {
			handle: (value) => (value.length < 5 ? "Invalid Handle" : null),
			password: (value) => (value.length < 12 ? "Invalid Password" : null),
		},
	});
	const { setAuth } = useContext(AuthContext)!;

	const { mutate, isLoading } = usePostLogin();
	const handleSubmit = (values: AuthUser) => {
		const data = { data: values };
		mutate(data, {
			onSuccess(data) {
				setAuth(data);
				setLoginOpened(false);
			},
			onError(error) {
				const err = error.response?.data.message;
				if (err?.startsWith("Handle")) {
					form.setFieldError("handle", err);
				} else {
					form.setErrors({ handle: err, password: err });
				}
			},
		});
	};

	const handleSignup = () => {
		setLoginOpened(false);
		setSignupOpened(true);
	};
	return (
		<Modal
			centered
			title="Log in"
			opened={loginOpened}
			onClose={() => setLoginOpened(false)}
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
				{"Don't have an account? "}
				<Anchor onClick={() => handleSignup()}> Sign up</Anchor>
			</Text>
		</Modal>
	);
}
