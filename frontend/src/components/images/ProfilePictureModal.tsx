import { Button, Flex, Modal } from "@mantine/core";
import { useState } from "react";
import "react-image-crop/dist/ReactCrop.css";
import { User } from "../../models";
import { putMe } from "../../spec.gen";
import { EditProfilePicture } from "./EditProfilePicture";

interface ProfilePictureModalProps {
	user: User;
	opened: boolean;
	onClose: () => void;
}

export const ProfilePictureModal = ({
	user,
	opened,
	onClose,
}: ProfilePictureModalProps) => {
	const [image, setImage] = useState<number | undefined>(undefined);
	const [cropOpened, setCropOpened] = useState(false);

	const handleSave = () => {
		putMe({ image: image })
			.then(onClose)
			.catch(() => {
				setImage(undefined);
			});
	};

	return (
		<Modal
			centered
			title={!cropOpened && "Set profile picture"}
			opened={opened}
			onClose={onClose}
			overflow="inside"
			styles={{ body: { overflow: "hidden" } }}
			withCloseButton={!cropOpened}
		>
			<EditProfilePicture
				user={user}
				setImage={(id) => setImage(id)}
				setCropOpened={setCropOpened}
			/>
			<Flex justify="flex-end" mt="md" hidden={cropOpened}>
				<Button disabled={!image} onClick={handleSave}>
					Save
				</Button>
			</Flex>
		</Modal>
	);
};
