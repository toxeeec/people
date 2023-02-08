import { Box } from "@mantine/core";

interface MessageProps {
	message: string;
	own: boolean;
}
export const Message = ({ message, own }: MessageProps) => {
	return (
		<Box
			bg={own ? "blue" : "gray"}
			c="white"
			px={12}
			py={8}
			m="sm"
			mb={0}
			style={{
				borderRadius: 18,
				overflowWrap: "break-word",
				alignSelf: own ? "flex-end" : "initial",
			}}
			maw="50%"
		>
			{message}
		</Box>
	);
};
