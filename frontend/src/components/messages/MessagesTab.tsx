import { Flex } from "@mantine/core";
import { useRef, useState } from "react";
import { ServerMessage } from "../../context/NotificationsContext";
import { getMessagesHandle } from "../../spec.gen";
import { Input } from "./Input";
import { NewMessages } from "./NewMessages";
import { UserMessages, Query } from "./UserMessages";

interface MessagesTabProps {
	messages: ServerMessage[];
	to: string;
}

export const MessagesTab = ({ messages, to }: MessagesTabProps) => {
	const [message, setMessage] = useState("");
	const ref = useRef<HTMLDivElement>(null);

	const messagesQuery: Query = (params) => {
		return getMessagesHandle(to, params);
	};

	return (
		<Flex direction="column" align="flex-start" h="100%">
			<Flex
				ref={ref}
				h="100%"
				w="100%"
				direction="column-reverse"
				align="flex-start"
				style={{ overflowY: "auto" }}
			>
				<NewMessages messages={messages} element={ref.current} />
				<UserMessages query={messagesQuery} queryKey={["messages", to]} />
			</Flex>
			<Input message={message} setMessage={setMessage} to={to} />
		</Flex>
	);
};
