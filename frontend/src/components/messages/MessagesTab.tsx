import { Flex, Space, Tabs } from "@mantine/core";
import { useRef, useState } from "react";
import { Message, Thread } from "../../models";
import { getThreadsThreadID } from "../../spec.gen";
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
		return getThreadsThreadID(thread.id, params);
	};

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
					<Space h={42} />
				</Flex>
				<Input message={message} setMessage={setMessage} threadID={thread.id} />
			</Flex>
		</Tabs.Panel>
	);
};
