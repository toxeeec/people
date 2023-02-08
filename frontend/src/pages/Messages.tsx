import { ActionIcon, Container, Flex, Tabs } from "@mantine/core";
import { useClickOutside, useDebouncedValue } from "@mantine/hooks";
import { useCallback, useContext, useState } from "react";
import { Search } from "../components/Search";
import { Query as UsersQuery, Users } from "../components/messages/Users";
import { Messages as MessagesComponent } from "../components/messages/Messages";
import { QueryKey } from "../query-key";
import { getUsersSearch } from "../spec.gen";
import { IconArrowLeft } from "@tabler/icons";
import { useNavigate, useParams } from "react-router-dom";
import { NotificationsContext } from "../context/NotificationsContext";

const Messages = () => {
	const params = useParams();
	const navigate = useNavigate();
	const [query, setQuery] = useState("");
	const [currentUser, setCurrentUser] = useState<string | null>(
		params.handle ?? null
	);
	const [hidden, setHidden] = useState(true);
	const ref = useClickOutside(() => setHidden(true));
	const [debounced] = useDebouncedValue(query, 200);
	const usersQuery: UsersQuery = useCallback(
		(params) => getUsersSearch({ query: debounced, ...params }),
		[debounced]
	);
	const close = () => {
		setHidden(true);
		setQuery("");
	};
	const handleClick = (handle: string) => {
		setCurrentUser(handle);
		navigate(`/messages/${handle}`);
		close();
	};
	const { newMessages } = useContext(NotificationsContext);
	return (
		<Tabs
			orientation="vertical"
			h="calc(100% - 60px)"
			value={currentUser}
			onTabChange={setCurrentUser}
		>
			<Tabs.List w="25%" style={{ flexWrap: "nowrap" }}>
				<Flex onFocus={() => setHidden(false)} align="center">
					<ActionIcon hidden={hidden} mx="xs" onClick={close}>
						<IconArrowLeft />
					</ActionIcon>
					<Search value={query} setValue={setQuery} />
				</Flex>
				<Container
					ref={ref}
					p={0}
					w="100%"
					hidden={hidden}
					mih="calc(100% - 60px)"
				>
					<Users
						enabled={debounced.length > 0}
						query={usersQuery}
						queryKey={[QueryKey.USERS, QueryKey.MESSAGES, debounced]}
						onClick={handleClick}
					/>
				</Container>
				<Flex direction="column" hidden={!hidden}>
					{[...newMessages.keys()].map((handle) => (
						<Tabs.Tab key={handle} value={handle}>
							{handle}
						</Tabs.Tab>
					))}
				</Flex>
			</Tabs.List>
			{[...newMessages.entries()].map(([to, messages]) => (
				<Tabs.Panel key={to} value={to}>
					<MessagesComponent messages={messages} currentUser={currentUser} />
				</Tabs.Panel>
			))}
		</Tabs>
	);
};

export default Messages;
