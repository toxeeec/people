import { ActionIcon, Box, Flex, Tabs } from "@mantine/core";
import { useClickOutside, useDebouncedValue, useMediaQuery } from "@mantine/hooks";
import { IconArrowLeft } from "@tabler/icons";
import { useCallback, useContext, useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { SearchBar } from "@/components/search";
import { NotificationsContext } from "@/context/NotificationsContext";
import { type Message, type Thread } from "@/models";
import { getUsersHandleThread } from "@/spec.gen";
import { InfiniteSearchUsers } from "@/components/search/InfiniteSearchUsers";
import { InfiniteTabs } from "@/components/messages/InfiniteTabs";
import { Tab } from "@/components/messages/Tab";
import { HEADER_HEIGHT } from "@/layout/Header";
import { RouteContext } from "@/context/RouteContext";

function Messages() {
	const { setRouteName } = useContext(RouteContext);
	useEffect(() => {
		setRouteName("Messages");
	}, [setRouteName]);

	const params = useParams();
	const navigate = useNavigate();
	const [query, setQuery] = useState("");
	const [hidden, setHidden] = useState(true);
	const [threads, setThreads] = useState<Thread[]>([]);
	const matches = useMediaQuery("(min-width: 720px)");
	const { newMessages, addMessageCallback, removeMessageCallback } =
		useContext(NotificationsContext);
	const [currentThread, setCurrentThread] = useState<string | null>(params.thread ?? null);

	const close = () => {
		setHidden(true);
		setQuery("");
	};

	const handleChange = useCallback(
		(thread: string) => {
			setCurrentThread(thread);
			navigate(`/messages/${thread}`, { replace: matches });
			close();
		},
		[navigate, matches]
	);

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

	const getNewThread = (handle: string) => {
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

	useEffect(() => {
		if (params.thread) return;
		if (matches && currentThread === null && threads.length > 0) {
			handleChange("" + threads[0].id);
		} else if (!params.thread) {
			setCurrentThread(null);
		}
	}, [currentThread, threads, matches, handleChange, params]);

	useEffect(() => {
		const cb = (msg: Message) => {
			setThreads((threads) =>
				threads.map((thread) => {
					return thread.id === msg.threadID ? { ...thread, latest: msg } : thread;
				})
			);
			sortThreads();
		};
		addMessageCallback(cb);
		return () => removeMessageCallback(cb);
	}, [addMessageCallback, removeMessageCallback, threads]);

	const ref = useClickOutside(() => setHidden(true));
	const [debounced] = useDebouncedValue(query, 200);

	return (
		<Tabs
			orientation="vertical"
			h={`calc(100% - ${HEADER_HEIGHT}px)`}
			value={currentThread}
			onTabChange={handleChange}
			styles={{
				tabLabel: { overflow: "hidden" },
				tab: { padding: "8px" },
				tabsList: { width: matches ? "360px" : params.thread ? 0 : "100vw" },
				panel: { display: !matches && !params.thread ? "none" : "initial" },
			}}
		>
			<Tabs.List h="100%" mih="0" style={{ overflow: "hidden" }}>
				<Flex direction="column" h="100%">
					<Flex onFocus={() => setHidden(false)} align="center">
						<ActionIcon hidden={hidden} mx="xs" onClick={close}>
							<IconArrowLeft />
						</ActionIcon>
						<SearchBar value={query} setValue={setQuery} />
					</Flex>
					<InfiniteSearchUsers
						ref={ref}
						hidden={hidden}
						query={debounced}
						onClick={getNewThread}
						enabled={debounced.length > 0}
					/>
					<Box hidden={!hidden} w="100%">
						<InfiniteTabs
							threads={threads}
							setThreads={setThreads}
							initialThread={params.thread}
							sortThreads={sortThreads}
						/>
					</Box>
				</Flex>
			</Tabs.List>
			{threads.map((thread) => (
				<Tab key={thread.id} thread={thread} newMessages={newMessages.get(thread.id)} />
			))}
		</Tabs>
	);
}

export default Messages;
