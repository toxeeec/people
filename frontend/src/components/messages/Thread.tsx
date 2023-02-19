import { Avatar, Flex, Stack, Text } from "@mantine/core";
import { forwardRef, useContext } from "react";
import { AuthContext } from "../../context/AuthContext";
import { Thread as ThreadType } from "../../models";

interface ThreadProps {
	thread: ThreadType;
}

export const Thread = forwardRef<HTMLDivElement, ThreadProps>(
	({ thread }, ref) => {
		const { getAuth } = useContext(AuthContext);
		const { handle } = getAuth();
		const user =
			thread.users.length > 1
				? thread.users.filter((u) => u.handle !== handle)[0]
				: thread.users[0];
		const prefix = thread.latest?.from.handle === handle ? "You: " : "";
		const content = thread.latest?.content;
		return (
			<Flex w="100%" align="center" ref={ref}>
				<Avatar radius="xl" size="lg" mr="sm" />
				<Stack w="100%" spacing={0} style={{ overflow: "hidden" }}>
					<Text weight="bold">{user.handle}</Text>
					<Text
						truncate
						style={{ visibility: content ? "initial" : "hidden" }}
					>{`${prefix}${content}`}</Text>
				</Stack>
			</Flex>
		);
	}
);

Thread.displayName = "User";