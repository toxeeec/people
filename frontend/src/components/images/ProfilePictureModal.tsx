import { Button, Flex, Modal } from "@mantine/core";
import { useState } from "react";
import "react-image-crop/dist/ReactCrop.css";
import { type User } from "@/models";
import { putMe } from "@/spec.gen";
import { EditProfilePicture } from "@/components/images/EditProfilePicture";

type ProfilePictureModalProps = {
	user: User;
	opened: boolean;
	handleChange: () => void;
	handleClose: () => void;
};

export function ProfilePictureModal({
	user,
	opened,
	handleChange,
	handleClose,
}: ProfilePictureModalProps) {
	const [image, setImage] = useState<number | undefined>();
	const [cropOpened, setCropOpened] = useState(false);

	const handleSave = () => {
		putMe({ image: image })
			.then(handleChange)
			.catch((e) => e);
		handleClose();
	};

	return (
		<Modal
			centered
			title={!cropOpened && "Set profile picture"}
			opened={opened}
			onClose={handleClose}
			overflow="inside"
			styles={{ body: { overflow: "hidden" } }}
			withCloseButton={!cropOpened}
		>
			<EditProfilePicture
				user={user}
				setImage={(id) => setImage(id)}
				cropOpened={cropOpened}
				setCropOpened={setCropOpened}
				handleRemoveImage={() => setImage(undefined)}
			/>
			<Flex justify="flex-end" mt="md" hidden={cropOpened}>
				<Button disabled={!image} onClick={handleSave}>
					Save
				</Button>
			</Flex>
		</Modal>
	);
}
