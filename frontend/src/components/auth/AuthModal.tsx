import {
	Button,
	Text,
	LoadingOverlay,
	Modal,
	PasswordInput,
	TextInput,
	Anchor,
} from "@mantine/core";
import { type FormErrors, type UseFormReturnType } from "@mantine/form";
import { type Dispatch, type SetStateAction, useContext } from "react";
import { type AuthResponse, type AuthUser, type Error } from "@/models";
import { type UseMutateFunction } from "@tanstack/react-query";
import { type ErrorType } from "@/custom-instance";
import { AuthContext } from "@/context/AuthContext";

type AuthModalProps = {
	title: string;
	opened: boolean;
	setOpened: Dispatch<SetStateAction<boolean>>;
	isLoading: boolean;
	form: UseFormReturnType<AuthUser, (values: AuthUser) => AuthUser>;
	text: string;
	handleChange: () => void;
	buttonText: string;
	mutate: UseMutateFunction<AuthResponse, ErrorType<Error>, { data: AuthUser }, unknown>;
	onSuccess?: () => void;
};

export function AuthModal({
	title,
	opened,
	setOpened,
	isLoading,
	form,
	text,
	handleChange,
	buttonText,
	mutate,
	onSuccess,
}: AuthModalProps) {
	const { setAuth } = useContext(AuthContext);

	const handleSubmit = (data: AuthUser) => {
		mutate(
			{ data },
			{
				onSuccess: ({ tokens, user }) => {
					onSuccess?.();
					setOpened(false);
					setAuth({ tokens, handle: user.handle });
				},
				onError: ({ response }) => {
					const msg = response?.data.message
						.replaceAll("Handle", "Username")
						.replaceAll("handle", "username");
					const errors: FormErrors = {};
					if (msg?.toLowerCase().includes("username")) {
						errors.handle = msg;
					}
					if (msg?.toLowerCase().includes("password")) {
						errors.password = msg;
					}
					form.setErrors(errors);
				},
			}
		);
	};

	return (
		<Modal
			centered
			title={title}
			opened={opened}
			onClose={() => {
				form.reset();
				setOpened(false);
			}}
		>
			<LoadingOverlay visible={isLoading} />
			<form onSubmit={form.onSubmit(handleSubmit)}>
				<TextInput withAsterisk label="Username" {...form.getInputProps("handle")} mb="md" />
				<PasswordInput withAsterisk label="Password" {...form.getInputProps("password")} mb="xl" />
				<Button fullWidth radius="xl" type="submit" mb="xl">
					Submit
				</Button>
			</form>
			<Text>
				{`${text} `}
				<Anchor onClick={handleChange}>{buttonText}</Anchor>
			</Text>
		</Modal>
	);
}
