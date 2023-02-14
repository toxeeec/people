import { Box } from "@mantine/core";
import { forwardRef } from "react";
import { Message as MessageType } from "../../models";

interface MessageProps {
	message: MessageType;
	own: boolean;
}
export const Message = forwardRef<HTMLDivElement, MessageProps>(
	({ message, own }, ref) => {
		return (
			<Box
				ref={ref}
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
				{message.content}
			</Box>
		);
	}
);

Message.displayName = "Message";
