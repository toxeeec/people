import { Box, Button, Flex, Group } from "@mantine/core";
import { type SyntheticEvent, useRef, useState } from "react";
import ReactCrop, { centerCrop, type Crop as CropType, makeAspectCrop } from "react-image-crop";
import { type User } from "@/models";

const MAX_WIDTH = 240;

type CropProps = {
	user: User;
	src: string;
	handleClose: () => void;
	handleChange: (file: File | null) => void;
};

export function Crop({ user, src, handleClose, handleChange }: CropProps) {
	const [crop, setCrop] = useState<CropType>();
	const imgRef = useRef<HTMLImageElement>(null);
	const [dimensions, setDimensions] = useState({ width: 0, height: 0 });

	const handleLoad = (e: SyntheticEvent<HTMLImageElement, Event>) => {
		const { naturalWidth: width, naturalHeight: height } = e.currentTarget;
		const crop = centerCrop(
			makeAspectCrop({ unit: "%", width: 100 }, 1 / 1, width, height),
			width,
			height
		);
		setCrop(crop);
		setDimensions({ width, height });
	};

	const handleApply = () => {
		if (!crop || !imgRef.current) return;
		const canvas = document.createElement("canvas");
		const ctx = canvas.getContext("2d");
		if (!ctx) return;

		const width = Math.round((dimensions.width * crop.width) / 100);
		const scale = MAX_WIDTH / width;
		canvas.width = width * scale;
		canvas.height = canvas.width;
		const x = (dimensions.width * crop.x) / 100;
		const y = (dimensions.height * crop.y) / 100;
		ctx.imageSmoothingQuality = "high";
		ctx.drawImage(imgRef.current, x, y, width, width, 0, 0, width * scale, width * scale);

		canvas.toBlob((blob) => {
			if (!blob) return;
			const file = new File([blob], `${user.handle}.webp`, {
				type: "image/webp",
			});
			handleChange(file);
			handleClose();
		});
	};

	return (
		<Flex direction="column" align="flex-end" mah="inherit">
			<Group mb="md">
				<Button variant="outline" onClick={handleClose}>
					Cancel
				</Button>
				<Button onClick={handleApply}>Apply</Button>
			</Group>
			<Box h="100%" style={{ overflowY: "auto" }}>
				<ReactCrop
					crop={crop}
					onChange={(_, percentCrop) => setCrop(percentCrop)}
					locked
					circularCrop
				>
					<img
						ref={imgRef}
						src={src}
						onLoad={handleLoad}
						width="100%"
						height="100%"
						style={{ display: "block" }}
					/>
				</ReactCrop>
			</Box>
		</Flex>
	);
}
