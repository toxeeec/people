import { Image as ImageComponent, Modal } from "@mantine/core";
import { useState } from "react";
import { baseURL } from "../../custom-instance";
import { stopPropagation } from "../../utils";

interface ImageProps {
	path: string;
}

export const Image = ({ path }: ImageProps) => {
	const [opened, setOpened] = useState(false);
	return (
		<>
			<ImageComponent
				src={`${baseURL}${path}`}
				styles={{ image: { aspectRatio: "3 / 2" } }}
				radius="lg"
				onClick={(e) => {
					e.stopPropagation();
					setOpened(true);
				}}
			/>
			<Modal
				centered
				size="auto"
				opened={opened}
				onClose={() => setOpened(false)}
				onClick={stopPropagation}
			>
				<ImageComponent
					src={`${baseURL}${path}`}
					styles={{ image: { maxHeight: "75vh", maxWidth: "75vw" } }}
				/>
			</Modal>
		</>
	);
};
