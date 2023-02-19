import { Tabs } from "@mantine/core";
import { useInfiniteQuery } from "@tanstack/react-query";
import { Dispatch, SetStateAction, useEffect } from "react";
import { useInView } from "react-intersection-observer";
import { Thread as ThreadType } from "../../models";
import { getThreads, getThreadsThreadID } from "../../spec.gen";
import { Thread } from "./Thread";

const queryLimit = 1;

interface MessagesTabProps {
	threads: ThreadType[];
	setThreads: Dispatch<SetStateAction<ThreadType[]>>;
	initialThread?: string;
	sortThreads: () => void;
}

export const MessagesTabs = ({
	threads,
	setThreads,
	initialThread,
	sortThreads,
}: MessagesTabProps) => {
	const { ref, inView } = useInView();
	const { isFetchingNextPage, fetchNextPage, hasNextPage } = useInfiniteQuery({
		queryKey: ["messages", "latest"],
		queryFn({ pageParam }) {
			return getThreads({ ...pageParam, limit: queryLimit });
		},
		async onSuccess({ pages }) {
			const newThreads = pages.flatMap((p) => p.data);
			if (initialThread && !newThreads.find((t) => t.id === +initialThread)) {
				try {
					const t = await getThreadsThreadID(+initialThread);
					newThreads.push(t);
				} catch (e) {
					return e;
				}
			}
			setThreads(newThreads);
			sortThreads();
		},
		getNextPageParam: (lastPage) => {
			if (!lastPage.meta || lastPage.data.length < queryLimit) return undefined;
			return { before: lastPage.meta?.oldest };
		},
		refetchOnMount: true,
		refetchOnWindowFocus: false,
	});

	useEffect(() => {
		if (inView && !isFetchingNextPage) {
			fetchNextPage();
		}
	}, [fetchNextPage, inView, hasNextPage, isFetchingNextPage]);

	return (
		<>
			{threads.map((thread) => (
				<Tabs.Tab ref={ref} key={thread.id} value={"" + thread.id} w="100%">
					<Thread thread={thread} />
				</Tabs.Tab>
			))}
		</>
	);
};
