import { Grid } from "@mantine/core";
import { Image } from "./Image";

interface ImagesProps {
	images?: string[];
}

export const Images = ({ images }: ImagesProps) => {
	return (
		<Grid grow gutter="xs" my="xs" style={{ cursor: "pointer" }}>
			{images?.map((path) => (
				<Grid.Col span={6} key={path}>
					<Image path={path} />
				</Grid.Col>
			))}
		</Grid>
	);
};
