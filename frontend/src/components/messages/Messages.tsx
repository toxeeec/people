import { Flex } from "@mantine/core";
import { useContext, useEffect, useMemo, useRef, useState } from "react";
import { AuthContext } from "../../context/AuthContext";
import { ServerMessage } from "../../context/NotificationsContext";
import { Input } from "./Input";
import { Message } from "./Message";

interface MessagesProps {
	messages: ServerMessage[];
	to: string;
}
export const Messages = ({ messages, to }: MessagesProps) => {
	const { getAuth } = useContext(AuthContext);
	const [message, setMessage] = useState("");
	const handle = useMemo(() => getAuth().handle, [getAuth]);
	const ref = useRef<HTMLDivElement>(null);

	useEffect(() => {
		const element = ref.current;
		if (element) {
			element.scrollTop = element.scrollHeight;
		}
	}, [messages]);

	return (
		<Flex direction="column" align="flex-start" h="100%">
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
						message={message.message}
						own={message.from === handle}
					/>
				))}
			</Flex>
			<Input message={message} setMessage={setMessage} to={to} />
		</Flex>
	);
};
