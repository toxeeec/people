import { EditProfilePicture } from "@/components/images";
import { AuthContext } from "@/context/AuthContext";
import { type UpdatedUser, type User } from "@/models";
import { usePutMe } from "@/spec.gen";
import { Button, Modal, TextInput } from "@mantine/core";
import { useForm } from "@mantine/form";
import { useQueryClient } from "@tanstack/react-query";
import { useContext, useState } from "react";
import { useNavigate } from "react-router-dom";

type EditButtonProps = {
	user: User;
};

export function EditButton({ user }: EditButtonProps) {
	const { setAuth } = useContext(AuthContext);
	const navigate = useNavigate();
	const form = useForm<UpdatedUser>({
		initialValues: { handle: user.handle },
		validate: {
			handle: (value) => {
				if (!value || value.length < 5) return "Username must contain at least 5 characters";
				if (!value || value.length > 15) return "Username cannot contain more than 15 characters";
			},
		},
	});
	const [opened, setOpened] = useState(false);
	const [cropOpened, setCropOpened] = useState(false);
	const queryClient = useQueryClient();
	const { mutate, status } = usePutMe({
		mutation: {
			retry: 1,
			onSuccess: (newUser) => {
				setOpened(false);
				if (newUser.handle !== user.handle) {
					setAuth({ handle: newUser.handle });
					navigate(`/${newUser.handle}`, { replace: true });
					form.resetDirty();
				} else {
					queryClient.resetQueries();
				}
			},
			onError: ({ response }) => {
				const msg = response?.data.message;
				form.setErrors({ handle: msg });
			},
		},
	});

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
				<form onSubmit={form.onSubmit((data) => mutate({ data }))}>
					<EditProfilePicture
						user={user}
						setImage={(image) => form.setValues((prev) => ({ ...prev, image }))}
						handleRemoveImage={() => {
							form.setValues((prev) => ({ ...prev, image: 0 }));
							if (user.image) form.setDirty({ image: true });
						}}
						cropOpened={cropOpened}
						setCropOpened={setCropOpened}
					/>
					{!cropOpened && (
						<>
							<TextInput label="Handle" {...form.getInputProps("handle")} mb="md" />
							<Button
								fullWidth
								radius="xl"
								type="submit"
								mb="xl"
								disabled={!form.isValid() || (!form.isDirty("image") && !form.isDirty("handle"))}
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
}
