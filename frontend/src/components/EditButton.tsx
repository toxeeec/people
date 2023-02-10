import { Button, Modal, TextInput } from "@mantine/core";
import { useForm } from "@mantine/form";
import { useQueryClient } from "@tanstack/react-query";
import { useContext, useState } from "react";
import { useNavigate } from "react-router-dom";
import { AuthContext } from "../context/AuthContext";
import { Handle } from "../models";
import { usePutMe } from "../spec.gen";

interface EditButtonProps {
	handle: string;
}

export const EditButton = ({ handle }: EditButtonProps) => {
	const { setAuth } = useContext(AuthContext);
	const queryClient = useQueryClient();
	const navigate = useNavigate();
	const form = useForm<Handle>({ initialValues: { handle } });
	const [opened, setOpened] = useState(false);
	const { mutate } = usePutMe({
		mutation: {
			retry: 1,
			onSuccess: (user) => {
				queryClient.invalidateQueries();
				setOpened(false);
				navigate(`/${user.handle}`);
				setAuth({ handle: user.handle });
			},
			onError: (error) => {
				const err = error.response?.data.message;
				form.setErrors({ handle: err });
			},
		},
	});

	const handleSubmit = (h: Handle) => {
		if (h.handle !== handle) {
			mutate({ data: h });
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
