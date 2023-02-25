import { CloseButton, Image } from "@mantine/core";
import { type Dispatch, type SetStateAction, useEffect } from "react";
import { type ImageResponse } from "@/models";
import { usePostImages } from "@/spec.gen";

type CreateImageProps = {
	image: File;
	handleRemove: () => void;
	setImageResponses: Dispatch<SetStateAction<Promise<ImageResponse>[]>>;
};
export function CreateImage({ image, handleRemove, setImageResponses }: CreateImageProps) {
	const { mutateAsync, status } = usePostImages({ mutation: { retry: 1 } });
	useEffect(() => {
		const imageResponse = mutateAsync({ data: { image } });
		setImageResponses((irs) => [...irs, imageResponse]);
		return () => setImageResponses((irs) => irs.filter((ir) => ir !== imageResponse));
	}, [mutateAsync, image, setImageResponses, status]);
	const src = URL.createObjectURL(image);
	return (
		<div style={{ position: "relative" }}>
			<Image src={src} styles={{ image: { aspectRatio: "3 / 2" } }} radius="lg" />
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
}
