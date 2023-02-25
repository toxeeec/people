import { useInterval } from "@mantine/hooks";
import { createContext, useCallback, useContext, useEffect, useMemo, useState } from "react";
import { wsURL } from "@/custom-instance";
import { type Message } from "@/models";
import { AuthContext } from "@/context/AuthContext";

const NotificationTypes = ["message"] as const;
type NotificationType = typeof NotificationTypes[number];

type Notification<T> = {
	type: NotificationType;
	data?: T;
};

type Messages = Map<number, Message[]>;

export type UserMessage = {
	type: NotificationType;
	content: string;
	threadID: number;
};

type NotificationsContextType = {
	newMessages: Messages;
	sendMessage: (message: Omit<UserMessage, "type">) => void;
	addMessageCallback: (cb: MessageCallback) => void;
	removeMessageCallback: (cb: MessageCallback) => void;
};

export const NotificationsContext = createContext<NotificationsContextType>(
	{} as NotificationsContextType
);

type NotificationsContextProviderProps = {
	children: React.ReactNode;
};

type MessageCallback = (msg: Message) => void;

export function NotificationsContextProvider({ children }: NotificationsContextProviderProps) {
	const { getAuth, isAuthenticated } = useContext(AuthContext);
	const [socket, setSocket] = useState<WebSocket>();
	const createConnection = useCallback(() => {
		if (!isAuthenticated) return;
		setSocket(new WebSocket(`${wsURL}?access_token=${getAuth().accessToken}`));
	}, [getAuth, isAuthenticated]);
	useEffect(() => {
		createConnection();
	}, [createConnection]);

	const [newMessages, setNewMessages] = useState<Messages>(new Map());
	const addNewMessage = useCallback(
		(msg: Message) => {
			const threadMessages = newMessages.get(msg.threadID);
			const messages = threadMessages
				? newMessages.set(msg.threadID, [...threadMessages, msg])
				: newMessages.set(msg.threadID, [msg]);
			setNewMessages(new Map(messages));
		},
		[newMessages]
	);

	let messageCallbacks: MessageCallback[] = useMemo(() => [], []);
	const addMessageCallback = (cb: MessageCallback) => {
		messageCallbacks.push(cb);
	};
	const removeMessageCallback = (cb: MessageCallback) => {
		messageCallbacks = messageCallbacks.filter((c) => c !== cb);
	};
	const interval = useInterval(() => createConnection(), 2000);
	const handleOpen = useCallback(() => interval.stop(), [interval]);
	const handleClose = useCallback(() => interval.start(), [interval]);
	const handleMessage = useCallback(
		(e: MessageEvent<string>) => {
			const { type, data } = JSON.parse(e.data) as Notification<unknown>;
			if (type === "message") {
				const msg = data as Message;
				messageCallbacks.forEach((cb) => cb(msg));
				addNewMessage(msg);
			}
		},
		[messageCallbacks, addNewMessage]
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

	const sendMessage = (msg: Omit<UserMessage, "type">) => {
		if (!isAuthenticated) return;
		const { content, threadID } = msg;
		const userMessage: UserMessage = {
			type: "message",
			threadID,
			content,
		};
		socket?.send(JSON.stringify(userMessage));
	};

	return (
		<NotificationsContext.Provider
			value={{
				newMessages,
				sendMessage,
				addMessageCallback,
				removeMessageCallback,
			}}
		>
			{children}
		</NotificationsContext.Provider>
	);
}
