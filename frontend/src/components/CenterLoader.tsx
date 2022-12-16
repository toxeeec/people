import { Center, Loader } from "@mantine/core";

export const CenterLoader = () => {
	return (
		<Center mt="xl" style={{ height: "100%" }}>
			<Loader />
		</Center>
	);
};
