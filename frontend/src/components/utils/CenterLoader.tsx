import { Center, Loader } from "@mantine/core";

export function CenterLoader() {
	return (
		<Center pos="absolute" left="50%" top="50%" style={{ transform: "translate(-50%, -50%)" }}>
			<Loader />
		</Center>
	);
}
