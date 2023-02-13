import { Flex } from "@mantine/core";
import { useContext, useEffect, useMemo, useRef } from "react";
import { AuthContext } from "../../context/AuthContext";
import { ServerMessage } from "../../context/NotificationsContext";
import { Message } from "./Message";

interface MessagesProps {
	messages: ServerMessage[];
}
export const Messages = ({ messages }: MessagesProps) => {
	const { getAuth } = useContext(AuthContext);
	const handle = useMemo(() => getAuth().handle, [getAuth]);
	const ref = useRef<HTMLDivElement>(null);

	useEffect(() => {
		const element = ref.current;
		if (element) {
			element.scrollTop = element.scrollHeight;
		}
	}, [messages]);

	return (
		<Flex
			ref={ref}
			h="100%"
			w="100%"
			direction="column"
			align="flex-start"
			style={{ overflowY: "auto" }}
		>
			{messages.map((message, i) => (
				<Message
					key={i}
					message={message.content}
					own={message.from === handle}
				/>
			))}
		</Flex>
	);
};
