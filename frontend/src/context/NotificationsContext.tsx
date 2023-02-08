import {
	createContext,
	useCallback,
	useContext,
	useEffect,
	useMemo,
	useState,
} from "react";
import { baseURL } from "../custom-instance";
import { AuthContext } from "./AuthContext";

const NotificationTypes = ["message"] as const;
type NotificationType = typeof NotificationTypes[number];

type Notification = {
	type: NotificationType;
	from: string;
	content?: Message;
};

export type Message = {
	message: string;
};

export type UserMessage = Message & {
	to: string;
};

export type ServerMessage = Message & {
	from: string;
};

type Messages = Map<string, ServerMessage[]>;

interface NotificationsContextType {
	newMessages: Messages;
	addNewMessage: (from: string, message: UserMessage) => void;
	sendMessage: (message: UserMessage) => void;
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
	const socket = useMemo(() => {
		if (isAuthenticated)
			return new WebSocket(
				`ws://${baseURL}/ws?access_token=${getAuth().accessToken}`
			);
	}, [isAuthenticated, getAuth]);
	const [newMessages, setNewMessages] = useState<Messages>(new Map());

	const addNewMessage = useCallback(
		(from: string, message: UserMessage) => {
			const userMessages = newMessages.get(message.to);
			const messages = userMessages
				? newMessages.set(message.to, [
						...userMessages,
						{ message: message.message, from },
				  ])
				: newMessages.set(message.to, [{ message: message.message, from }]);
			setNewMessages(new Map(messages));
		},
		[newMessages]
	);

	const handleMessage = useCallback(
		(e: MessageEvent<string>) => {
			const notif = JSON.parse(e.data) as Notification;
			if (notif.type === "message" && notif.content) {
				const message = { message: notif.content.message, to: notif.from };
				addNewMessage(notif.from, message);
			}
		},
		[addNewMessage]
	);

	useEffect(() => {
		socket?.addEventListener("message", handleMessage);
		return () => {
			socket?.removeEventListener("message", handleMessage);
		};
	}, [getAuth, isAuthenticated, addNewMessage, handleMessage, socket]);

	const sendMessage = (msg: UserMessage) => {
		const { message, to } = msg;
		if (!isAuthenticated) return;
		const userMessage: UserMessage = {
			message,
			to,
		};
		socket?.send(JSON.stringify({ ...userMessage, type: "message" }));
		addNewMessage(getAuth().handle!, {
			message: message,
			to,
		});
	};

	return (
		<NotificationsContext.Provider
			value={{ newMessages, addNewMessage, sendMessage }}
		>
			{children}
		</NotificationsContext.Provider>
	);
};
