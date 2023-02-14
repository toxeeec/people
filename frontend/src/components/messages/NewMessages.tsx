import { useContext, useEffect, useMemo } from "react";
import { AuthContext } from "../../context/AuthContext";
import { ServerMessage } from "../../context/NotificationsContext";
import { Message } from "./Message";

interface MessagesProps {
	element: HTMLDivElement | null;
	messages: ServerMessage[];
}
export const NewMessages = ({ element, messages }: MessagesProps) => {
	const { getAuth } = useContext(AuthContext);
	const handle = useMemo(() => getAuth().handle, [getAuth]);

	useEffect(() => {
		if (element) {
			element.scrollTop = element.scrollHeight;
		}
	}, [element, messages]);

	return (
		<>
			{messages
				.slice()
				.reverse()
				.map((message, i) => (
					<Message
						key={i}
						message={message.message}
						own={message.from === handle}
					/>
				))}
		</>
	);
};
