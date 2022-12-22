import { CloseButton, Image } from "@mantine/core";
import { Dispatch, SetStateAction, useEffect } from "react";
import { ImageResponse } from "../../models";
import { postImages } from "../../spec.gen";

interface CreateImageProps {
	file: File;
	setFiles: Dispatch<SetStateAction<File[]>>;
	setImageResponses: Dispatch<SetStateAction<Promise<ImageResponse>[]>>;
}
export const CreateImage = ({
	file,
	setFiles,
	setImageResponses,
}: CreateImageProps) => {
	const handleRemove = () => {
		setFiles((files) => files.filter((f) => f !== file));
	};
	useEffect(() => {
		const imageResponse = postImages({ image: file });
		setImageResponses((irs) => [...irs, imageResponse]);
		return () =>
			setImageResponses((irs) => irs.filter((ir) => ir !== imageResponse));
	}, [file, setImageResponses]);
	const src = URL.createObjectURL(file);
	return (
		<div style={{ position: "relative" }}>
			<Image
				src={src}
				styles={{ image: { aspectRatio: "3 / 2" } }}
				radius="lg"
			/>
			<CloseButton
				onClick={handleRemove}
				pos="absolute"
				top={6}
				left={6}
				bg="dark"
				radius="xl"
				size="lg"
			/>
		</div>
	);
};
