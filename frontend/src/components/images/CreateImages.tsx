import { Grid } from "@mantine/core";
import { type Dispatch, type SetStateAction } from "react";
import { type ImageResponse } from "@/models";
import { CreateImage } from "@/components/images/CreateImage";

type CreateImagesProps = {
	files: File[];
	setFiles: Dispatch<SetStateAction<File[]>>;
	setImageResponses: Dispatch<SetStateAction<Promise<ImageResponse>[]>>;
};

export function CreateImages({ files, setFiles, setImageResponses }: CreateImagesProps) {
	return (
		<Grid>
			{files.map((file, i) => (
				<Grid.Col span={6} key={`${file.name}-${i}`}>
					<CreateImage
						image={file}
						handleRemove={() => setFiles((files) => files.filter((f) => file !== f))}
						setImageResponses={setImageResponses}
					/>
				</Grid.Col>
			))}
		</Grid>
	);
}
