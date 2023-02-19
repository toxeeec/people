import { useContext, useEffect, useMemo } from "react";
import { AuthContext } from "../../context/AuthContext";
import { Message as MessageType } from "../../models";
import { Message } from "./Message";

interface MessagesProps {
	element: HTMLDivElement | null;
	messages?: MessageType[];
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
				?.slice()
				.reverse()
				.map((message, i) => (
					<Message
						key={i}
						message={message}
						own={message.from.handle === handle}
					/>
				))}
		</>
	);
};
