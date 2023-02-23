import { Button, Modal, TextInput } from "@mantine/core";
import { useForm } from "@mantine/form";
import { useQueryClient } from "@tanstack/react-query";
import { useContext, useState } from "react";
import { useNavigate } from "react-router-dom";
import { AuthContext } from "../context/AuthContext";
import { UpdatedUser, User } from "../models";
import { usePutMe } from "../spec.gen";
import { EditProfilePicture } from "./images/EditProfilePicture";

interface EditButtonProps {
	user: User;
}

export const EditButton = ({ user }: EditButtonProps) => {
	const { setAuth } = useContext(AuthContext);
	const queryClient = useQueryClient();
	const navigate = useNavigate();
	const form = useForm<UpdatedUser>({
		initialValues: { handle: user.handle },
		validate: {
			handle: (value) =>
				value!.length < 5
					? "Username must have at least 5 characters"
					: value!.length > 15
					? "Username cannot have more than 15 characters"
					: null,
		},
	});
	const [opened, setOpened] = useState(false);
	const [cropOpened, setCropOpened] = useState(false);
	const { mutate, status } = usePutMe({
		mutation: {
			retry: 1,
			onSuccess: (newUser) => {
				setOpened(false);
				if (newUser.handle !== user.handle) {
					setAuth({ handle: newUser.handle });
					navigate(`/${newUser.handle}`, { replace: true });
				} else {
					queryClient.invalidateQueries();
				}
			},
			onError: (error) => {
				const err = error.response?.data.message;
				form.setErrors({ handle: err });
			},
		},
	});

	const handleSubmit = (u: UpdatedUser) => {
		mutate({ data: u });
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
				onClose={() => {
					setOpened(false);
					form.reset();
				}}
			>
				<form onSubmit={form.onSubmit(handleSubmit)}>
					<EditProfilePicture
						user={user}
						setImage={(image) => {
							form.setValues({ ...form.values, image });
						}}
						removeImage={() => {
							form.setValues({ ...form.values, image: undefined });
							if (user.image) form.setDirty({ image: true });
						}}
						setCropOpened={setCropOpened}
					/>
					{!cropOpened && (
						<>
							<TextInput
								label="Handle"
								{...form.getInputProps("handle")}
								mb="md"
							/>
							<Button
								fullWidth
								radius="xl"
								type="submit"
								mb="xl"
								disabled={
									!form.isValid() ||
									(!form.isDirty("image") && !form.isDirty("handle"))
								}
								loading={status === "loading"}
							>
								Save
							</Button>
						</>
					)}
				</form>
			</Modal>
		</>
	);
};
