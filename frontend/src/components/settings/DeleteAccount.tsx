import { Button, Modal, PasswordInput } from "@mantine/core";
import { useForm } from "@mantine/form";
import { useContext, useState } from "react";
import { AuthContext } from "../../context/AuthContext";
import { useDeleteMe } from "../../spec.gen";
import { DangerButton } from "../DangerButton";

export const DeleteAccount = () => {
	const { getAuth, clearAuth } = useContext(AuthContext);
	const [opened, setOpened] = useState(false);
	const form = useForm<{ password: string }>({
		initialValues: { password: "" },
	});
	const { mutate } = useDeleteMe({
		mutation: {
			retry: 1,
			onSuccess: clearAuth,
			onError: (error) => {
				const err = error.response?.data.message;
				form.setErrors({ password: err });
			},
		},
	});

	const handleDelete = ({ password }: { password: string }) => {
		if (!password) return;
		mutate({
			data: { password, refreshToken: getAuth().refreshToken! },
		});
	};
	return (
		<>
			<DangerButton onClick={() => setOpened(true)} text="Delete Account" />
			<Modal
				centered
				title="This action is irreversible"
				opened={opened}
				onClose={() => setOpened(false)}
			>
				<form onSubmit={form.onSubmit(handleDelete)}>
					<PasswordInput
						withAsterisk
						label="Password"
						{...form.getInputProps("password")}
						mb="xl"
					/>
					<Button fullWidth radius="xl" type="submit" mb="xl">
						Delete Account
					</Button>
				</form>
			</Modal>
		</>
	);
};
