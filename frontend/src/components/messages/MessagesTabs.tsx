import { Tabs } from "@mantine/core";
import { useInfiniteQuery } from "@tanstack/react-query";
import { Dispatch, SetStateAction, useEffect } from "react";
import { useInView } from "react-intersection-observer";
import { Thread as ThreadType } from "../../models";
import { getThreads } from "../../spec.gen";
import { Thread } from "./Thread";

const queryLimit = 10;

interface MessagesTabProps {
	threads: ThreadType[];
	setThreads: Dispatch<SetStateAction<ThreadType[]>>;
}

export const MessagesTabs = ({ threads, setThreads }: MessagesTabProps) => {
	const { ref, inView } = useInView();
	const { isFetchingNextPage, fetchNextPage, hasNextPage } = useInfiniteQuery({
		queryKey: ["messages", "latest"],
		queryFn({ pageParam }) {
			return getThreads({ ...pageParam, limit: queryLimit });
		},
		onSuccess({ pages }) {
			setThreads(pages.flatMap((p) => p.data));
		},
		getNextPageParam: (lastPage) => {
			if (!lastPage.meta || lastPage.data.length < queryLimit) return undefined;
			return { before: lastPage.meta?.oldest };
		},
		refetchOnMount: false,
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
