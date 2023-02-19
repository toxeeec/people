import { Flex, Space, Tabs } from "@mantine/core";
import { useContext, useRef, useState } from "react";
import { AuthContext } from "../../context/AuthContext";
import { Message, Thread } from "../../models";
import { getThreadsThreadIDMessages } from "../../spec.gen";
import { Header } from "./Header";
import { Input } from "./Input";
import { NewMessages } from "./NewMessages";
import { UserMessages, Query } from "./UserMessages";

interface MessagesTabProps {
	thread: Thread;
	newMessages?: Message[];
}

export const MessagesTab = ({ thread, newMessages }: MessagesTabProps) => {
	const [message, setMessage] = useState("");
	const ref = useRef<HTMLDivElement>(null);
	const messagesQuery: Query = (params) => {
		return getThreadsThreadIDMessages(thread.id, params);
	};
	const { getAuth } = useContext(AuthContext);
	const { handle } = getAuth();
	const user =
		thread.users.length > 1
			? thread.users.filter((u) => u.handle !== handle)[0]
			: thread.users[0];

	return (
		<Tabs.Panel value={"" + thread.id}>
			<Flex
				direction="column"
				h="100%"
				w="100%"
				pos="relative"
				justify="space-between"
			>
				<Flex
					ref={ref}
					w="100%"
					direction="column-reverse"
					align="flex-start"
					style={{ overflowY: "auto" }}
				>
					<NewMessages messages={newMessages} element={ref.current} />
					<UserMessages
						query={messagesQuery}
						queryKey={["messages", thread.id]}
					/>
					<Header user={user} />
					<Space h={42} />
				</Flex>
				<Input message={message} setMessage={setMessage} threadID={thread.id} />
			</Flex>
		</Tabs.Panel>
	);
};
