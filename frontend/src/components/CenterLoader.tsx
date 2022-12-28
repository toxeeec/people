import { Center, Loader } from "@mantine/core";

export const CenterLoader = () => {
	return (
		<Center
			pos="absolute"
			inline
			left="50%"
			top="50%"
			style={{ transform: "translate(-50%, -50%)" }}
		>
			<Loader />
		</Center>
	);
};
