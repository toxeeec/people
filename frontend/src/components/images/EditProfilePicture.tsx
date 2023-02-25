import { ActionIcon, FileButton, Flex, Group, Tooltip } from "@mantine/core";
import { IconPhoto, IconX } from "@tabler/icons";
import { useRef, useState } from "react";
import { Avatar } from "@/components/user";
import { type User } from "@/models";
import { postImages } from "@/spec.gen";
import "react-image-crop/dist/ReactCrop.css";
import { Crop } from "@/components/images/Crop";

type EditProfilePictureProps = {
	user: User;
	cropOpened: boolean;
	setCropOpened: (opened: boolean) => void;
	setImage?: (id: number | undefined) => void;
	handleRemoveImage?: () => void;
};

export function EditProfilePicture({
	user,
	setImage,
	cropOpened,
	setCropOpened,
	handleRemoveImage,
}: EditProfilePictureProps) {
	const [src, setSrc] = useState(user.image ?? "");
	const resetRef = useRef<() => void>(null);

	const handleChange = async (file: File | null) => {
		if (!file) return;
		const src = URL.createObjectURL(file);
		setSrc(src);
		const dimensions = await new Promise<{ width: number; height: number }>((resolve) => {
			const img = new Image();
			img.onload = () => {
				const { width, height } = img;
				resolve({ width, height });
			};
			img.src = src;
		});
		if (dimensions.width !== dimensions.height) {
			setCropOpened(true);
			return;
		}
		postImages({ image: file })
			.then((img) => {
				resetRef.current?.();
				setImage?.(img.id);
			})
			.catch((e) => e);
	};

	return cropOpened ? (
		<Crop
			user={user}
			src={src}
			handleChange={handleChange}
			handleClose={() => setCropOpened(false)}
		/>
	) : (
		<FileButton
			onChange={handleChange}
			accept="image/png,image/jpeg,image/webp"
			resetRef={resetRef}
		>
			{(props) => (
				<Flex direction="column" align="center" gap="md">
					<Avatar size={120} m="auto" user={{ ...user, image: src }} />
					<Group>
						<Tooltip label="Add picture">
							<ActionIcon {...props}>
								<IconPhoto />
							</ActionIcon>
						</Tooltip>
						<Tooltip label="Remove picture">
							<ActionIcon
								onClick={() => {
									resetRef.current?.();
									setSrc("");
									handleRemoveImage?.();
								}}
							>
								<IconX />
							</ActionIcon>
						</Tooltip>
					</Group>
				</Flex>
			)}
		</FileButton>
	);
}
