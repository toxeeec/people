import {
	Button,
	Text,
	LoadingOverlay,
	Modal,
	PasswordInput,
	TextInput,
	Anchor,
} from "@mantine/core";
import { UseFormReturnType } from "@mantine/form";
import { Dispatch, SetStateAction } from "react";
import { AuthUser } from "../../models";

interface AuthModalProps {
	title: string;
	opened: boolean;
	setOpened: Dispatch<SetStateAction<boolean>>;
	isLoading: boolean;
	form: UseFormReturnType<AuthUser, (values: AuthUser) => AuthUser>;
	handleSubmit: (user: AuthUser) => void;
	text: string;
	handleChange: () => void;
	buttonText: string;
}

export const AuthModal = ({
	title,
	opened,
	setOpened,
	isLoading,
	form,
	handleSubmit,
	text,
	handleChange,
	buttonText,
}: AuthModalProps) => {
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
				<TextInput
					withAsterisk
					label="Username"
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
				{text}
				<Anchor onClick={handleChange}>{buttonText}</Anchor>
			</Text>
		</Modal>
	);
};
