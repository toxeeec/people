import { useInterval } from "@mantine/hooks";
import {
	createContext,
	useCallback,
	useContext,
	useEffect,
	useState,
} from "react";
import { wsURL } from "../custom-instance";
import { Message } from "../models";
import { AuthContext } from "./AuthContext";

const NotificationTypes = ["message"] as const;
type NotificationType = typeof NotificationTypes[number];

type Notification = {
	type: NotificationType;
	content?: ServerMessage;
};

export type UserMessage = Message & {
	to: string;
};

export type ServerMessage = {
	message: Message;
	id: number;
	from: string;
	to: string;
	sentAt: Date;
};

type Messages = Map<string, ServerMessage[]>;

interface NotificationsContextType {
	newMessages: Messages;
	sendMessage: (message: UserMessage) => void;
	addUser: (handle: string) => void;
}

export const NotificationsContext = createContext<NotificationsContextType>(
	null!
);

interface NotificationsContextProviderProps {
	children: React.ReactNode;
}

export const NotificationsContextProvider = ({
	children,
}: NotificationsContextProviderProps) => {
	const { getAuth, isAuthenticated } = useContext(AuthContext);
	const [socket, setSocket] = useState<WebSocket>();

	const createConnection = useCallback(() => {
		if (!isAuthenticated) return;
		setSocket(new WebSocket(`${wsURL}?access_token=${getAuth().accessToken}`));
	}, [getAuth, isAuthenticated]);

	const interval = useInterval(() => createConnection(), 2000);

	useEffect(() => {
		createConnection();
	}, [createConnection]);

	const [newMessages, setNewMessages] = useState<Messages>(new Map());

	const addNewMessage = useCallback(
		(msg: ServerMessage) => {
			const userMessages = newMessages.get(msg.to);
			const messages = userMessages
				? newMessages.set(msg.to, [...userMessages, msg])
				: newMessages.set(msg.to, [msg]);
			setNewMessages(new Map(messages));
		},
		[newMessages]
	);

	const handleOpen = useCallback(() => interval.stop(), [interval]);
	const handleClose = useCallback(() => interval.start(), [interval]);
	const handleMessage = useCallback(
		(e: MessageEvent<string>) => {
			const { type, content } = JSON.parse(e.data) as Notification;
			if (type === "message" && content) {
				const handle = getAuth().handle!;
				if (content.to === handle) {
					content.to = content.from;
				}
				addNewMessage(content);
			}
		},
		[addNewMessage, getAuth]
	);

	useEffect(() => {
		socket?.addEventListener("open", handleOpen);
		socket?.addEventListener("message", handleMessage);
		socket?.addEventListener("close", handleClose);
		return () => {
			socket?.removeEventListener("open", handleOpen);
			socket?.removeEventListener("message", handleMessage);
			socket?.removeEventListener("close", handleClose);
		};
	}, [handleClose, handleMessage, handleOpen, socket]);

	const sendMessage = (msg: UserMessage) => {
		if (!isAuthenticated) return;
		const { content, to } = msg;
		const userMessage: UserMessage = {
			to,
			content,
		};
		socket?.send(JSON.stringify({ type: "message", ...userMessage }));
	};

	const addUser = (handle: string) => {
		const messages = newMessages.get(handle);
		if (messages) return;
		setNewMessages(new Map(newMessages.set(handle, [])));
	};

	return (
		<NotificationsContext.Provider
			value={{ newMessages, sendMessage, addUser }}
		>
			{children}
		</NotificationsContext.Provider>
	);
};
