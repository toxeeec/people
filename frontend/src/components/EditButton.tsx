import { Button, Modal, TextInput } from "@mantine/core";
import { useForm } from "@mantine/form";
import { useContext, useState } from "react";
import { useNavigate } from "react-router-dom";
import { AuthContext } from "../context/AuthContext";
import { UsersContext } from "../context/UsersContext";
import { Handle } from "../models";
import { usePutMe } from "../spec.gen";

interface EditButtonProps {
	handle: string;
}

export const EditButton = ({ handle }: EditButtonProps) => {
	const { users, setUser } = useContext(UsersContext);
	const { setAuth } = useContext(AuthContext);
	const navigate = useNavigate();
	const user = users[handle]!;
	const form = useForm<Handle>({ initialValues: { handle: user.handle } });
	const [opened, setOpened] = useState(false);
	const { mutate } = usePutMe({
		mutation: {
			retry: 1,
			onSuccess: (user) => {
				setOpened(false);
				setUser(user);
				navigate(`/${user.handle}`);
				setAuth({ handle: user.handle });
			},
			onError: (error) => {
				const err = error.response?.data.message;
				form.setErrors({ handle: err });
			},
		},
	});

	const handleSubmit = (handle: Handle) => {
		if (handle.handle !== user.handle) {
			mutate({ data: handle });
		} else {
			setOpened(false);
		}
	};

	return (
		<>
			<Button onClick={() => setOpened(true)} variant={"outline"} radius="xl">
				Edit Profile
			</Button>

			<Modal
				centered
				title="Edit profile"
				opened={opened}
				onClose={() => setOpened(false)}
			>
				<form onSubmit={form.onSubmit(handleSubmit)}>
					<TextInput label="Handle" {...form.getInputProps("handle")} mb="md" />
					<Button fullWidth radius="xl" type="submit" mb="xl">
						Save
					</Button>
				</form>
			</Modal>
		</>
	);
};
