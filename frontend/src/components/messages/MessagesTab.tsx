import { Flex } from "@mantine/core";
import { useState } from "react";
import { ServerMessage } from "../../context/NotificationsContext";
import { Input } from "./Input";
import { Messages } from "./Messages";

interface MessagesTabProps {
	messages: ServerMessage[];
	to: string;
}

export const MessagesTab = ({ messages, to }: MessagesTabProps) => {
	const [message, setMessage] = useState("");
	return (
		<Flex direction="column" align="flex-start" h="100%">
			<Messages messages={messages} />
			<Input message={message} setMessage={setMessage} to={to} />
		</Flex>
	);
};
