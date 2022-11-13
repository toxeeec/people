import { useInfiniteQuery } from "@tanstack/react-query";
import { Fragment, useEffect } from "react";
import { getMeFeed } from "../spec.gen";
import { GetMeFeedParams } from "../models/getMeFeedParams";
import { Container } from "@mantine/core";
import Post from "../components/Post";
import { useInView } from "react-intersection-observer";
import CenterLoader from "../components/CenterLoader";

const queryLimit = 10;

interface QueryFnProps {
	pageParam?: GetMeFeedParams;
}

export default function Home() {
	function queryFn({ pageParam }: QueryFnProps) {
		pageParam = { ...pageParam, limit: queryLimit };
		return getMeFeed(pageParam);
	}
	const {
		data,
		refetch,
		fetchNextPage,
		hasNextPage,
		isLoading,
		isFetchingNextPage,
		isError,
	} = useInfiniteQuery({
		queryKey: ["feed"],
		queryFn,
		getNextPageParam: (lastPage) => {
			if (!lastPage.meta) return undefined;
			return { before: lastPage.meta?.oldest };
		},
		enabled: false,
	});
	useEffect(() => {
		refetch();
	}, [refetch]);

	const { ref, inView } = useInView();
	useEffect(() => {
		if (inView && hasNextPage && !isFetchingNextPage && !isError) {
			fetchNextPage();
		}
	}, [fetchNextPage, inView, hasNextPage, isFetchingNextPage, isError]);

	return (
		<Container px={0}>
			{isLoading ? (
				<CenterLoader />
			) : (
				data?.pages.map((group, i) =>
					i > 0 && i === data.pages.length - 1 && isFetchingNextPage ? (
						<CenterLoader key={i} />
					) : (
						<Fragment key={i}>
							{group.data.map((post) => (
								<Post data={post} key={post.createdAt} ref={ref} />
							))}
						</Fragment>
					)
				)
			)}
		</Container>
	);
}
