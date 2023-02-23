import { ActionIcon, FileButton, Flex, Group, Tooltip } from "@mantine/core";
import { IconPhoto, IconX } from "@tabler/icons";
import { useEffect, useRef, useState } from "react";
import { Avatar } from "../../Avatar";
import { User } from "../../models";
import { postImages } from "../../spec.gen";
import "react-image-crop/dist/ReactCrop.css";
import { Crop } from "./Crop";

interface EditProfilePictureProps {
	user: User;
	setCropOpened: (opened: boolean) => void;
	setImage?: (id: number | undefined) => void;
	removeImage?: () => void;
}

export const EditProfilePicture = ({
	user: initialUser,
	setImage,
	setCropOpened,
	removeImage,
}: EditProfilePictureProps) => {
	const [user, setUser] = useState(initialUser);
	const [src, setSrc] = useState("");
	const resetRef = useRef<() => void>(null);

	useEffect(() => {
		setCropOpened(!!src);
	}, [src, setCropOpened]);

	const handleChange = async (file: File | null) => {
		if (!file) return;
		const src = URL.createObjectURL(file);
		const dimensions = await new Promise<{ width: number; height: number }>(
			(resolve) => {
				const img = new Image();
				img.onload = () => {
					const { width, height } = img;
					resolve({ width, height });
				};
				img.src = src;
			}
		);
		if (dimensions.width !== dimensions.height) {
			setSrc(src);
			return;
		}
		setUser((user) => ({ ...user, image: src }));
		postImages({ image: file })
			.then((img) => {
				setUser((user) => ({ ...user, image: src }));
				resetRef.current?.();
				setImage?.(img.id);
			})
			.catch((e) => e);
	};

	return src ? (
		<Crop
			user={user}
			src={src}
			handleChange={handleChange}
			handleCancel={() => setSrc("")}
		/>
	) : (
		<FileButton
			onChange={handleChange}
			accept="image/png,image/jpeg,image/webp"
			resetRef={resetRef}
		>
			{(props) => (
				<Flex direction="column" align="center" gap="md">
					<Avatar size={120} m="auto" user={user} />
					<Group>
						<Tooltip label="Add picture">
							<ActionIcon {...props}>
								<IconPhoto />
							</ActionIcon>
						</Tooltip>
						<Tooltip label="Remove picture">
							<ActionIcon
								onClick={() => {
									setUser((user) => ({ ...user, image: undefined }));
									resetRef.current?.();
									removeImage?.();
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
};
