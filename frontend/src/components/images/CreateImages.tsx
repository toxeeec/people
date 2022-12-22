import { Grid } from "@mantine/core";
import { Dispatch, SetStateAction } from "react";
import { ImageResponse } from "../../models";
import { CreateImage } from "./CreateImage";

interface CreateImagesProps {
	files: File[];
	setFiles: Dispatch<SetStateAction<File[]>>;
	setImageResponses: Dispatch<SetStateAction<Promise<ImageResponse>[]>>;
}

export const CreateImages = ({
	files,
	setFiles,
	setImageResponses,
}: CreateImagesProps) => {
	return (
		<Grid>
			{files.map((file, i) => (
				<Grid.Col span={6} key={`${file.name}-${i}`}>
					<CreateImage
						file={file}
						setFiles={setFiles}
						setImageResponses={setImageResponses}
					/>
				</Grid.Col>
			))}
		</Grid>
	);
};
