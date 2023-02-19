import { ActionIcon, Box, Flex, Tabs } from "@mantine/core";
import { useClickOutside, useDebouncedValue } from "@mantine/hooks";
import { useCallback, useContext, useEffect, useState } from "react";
import { Search } from "../components/Search";
import { IconArrowLeft } from "@tabler/icons";
import { useNavigate, useParams } from "react-router-dom";
import { SearchUsers } from "../components/messages/SearchUsers";
import { MessagesTabs } from "../components/messages/MessagesTabs";
import { MessagesTab } from "../components/messages/MessagesTab";
import { Message, Thread } from "../models";
import { NotificationsContext } from "../context/NotificationsContext";
import { getUsersHandleThread } from "../spec.gen";

const Messages = () => {
	const params = useParams();
	const navigate = useNavigate();
	const [query, setQuery] = useState("");
	const [currentThread, setCurrentThread] = useState<string | null>(
		params.thread ?? null
	);
	const [threads, setThreads] = useState<Thread[]>([]);

	const close = () => {
		setHidden(true);
		setQuery("");
	};

	const handleChange = useCallback(
		(thread: string) => {
			setCurrentThread(thread);
			navigate(`/messages/${thread}`, { replace: true });
			close();
		},
		[navigate]
	);

	useEffect(() => {
		if (currentThread !== null || threads.length === 0) return;
		handleChange("" + threads[0].id);
	}, [currentThread, threads, handleChange]);

	const [hidden, setHidden] = useState(true);
	const ref = useClickOutside(() => setHidden(true));
	const [debounced] = useDebouncedValue(query, 200);
	const { newMessages, addMessageCallback, removeMessageCallback } =
		useContext(NotificationsContext);

	const sortThreads = () => {
		setThreads((thread) => {
			return thread.sort((a, b) => {
				if (!a.latest) return 1;
				if (!b.latest) return -1;
				if (a.latest.id < b.latest.id) return 1;
				if (a.latest.id === b.latest.id) return 0;
				return -1;
			});
		});
	};

	useEffect(() => {
		const cb = (msg: Message) => {
			setThreads((threads) =>
				threads.map((thread) => {
					return thread.id === msg.threadID
						? { ...thread, latest: msg }
						: thread;
				})
			);
			sortThreads();
		};

		addMessageCallback(cb);
		return () => removeMessageCallback(cb);
	}, [addMessageCallback, removeMessageCallback, threads]);

	const onClick = (handle: string) => {
		getUsersHandleThread(handle).then((thread) => {
			const i = threads.findIndex((t) => t.id === thread.id);
			if (i === -1) {
				setThreads((threads) => [...threads, thread]);
			} else {
				threads[i] = thread;
			}
			sortThreads();
			handleChange("" + thread.id);
		});
	};

	return (
		<Tabs
			orientation="vertical"
			h="calc(100% - 60px)"
			value={currentThread}
			onTabChange={handleChange}
			styles={{ tabLabel: { overflow: "hidden" }, tab: { padding: "8px" } }}
		>
			<Tabs.List w="360px">
				<Flex onFocus={() => setHidden(false)} align="center">
					<ActionIcon hidden={hidden} mx="xs" onClick={close}>
						<IconArrowLeft />
					</ActionIcon>
					<Search value={query} setValue={setQuery} />
				</Flex>
				<SearchUsers
					ref={ref}
					hidden={hidden}
					debounced={debounced}
					onClick={onClick}
				/>
				<Box hidden={!hidden} w="100%">
					<MessagesTabs threads={threads} setThreads={setThreads} />
				</Box>
			</Tabs.List>
			{threads.map((thread) => (
				<MessagesTab
					key={thread.id}
					thread={thread}
					newMessages={newMessages.get(thread.id)}
				/>
			))}
		</Tabs>
	);
};

export default Messages;
