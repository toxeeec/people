import { useInfiniteQuery } from "@tanstack/react-query";
import { Fragment, useContext, useEffect } from "react";
import { Posts as PostsData } from "../models";
import { Container } from "@mantine/core";
import Post from "../components/post";
import { useInView } from "react-intersection-observer";
import CenterLoader from "../components/CenterLoader";
import UsersContext from "../context/UsersContext";

const queryLimit = 10;

interface PaginationParams {
	limit?: number;
	before?: number;
	after?: number;
}

type Query = (_props: PaginationParams) => Promise<PostsData>;

interface PostsProps {
	query: Query;
}

interface QueryFunctionArgs {
	pageParam?: PaginationParams;
}

export default function Posts({ query }: PostsProps) {
	function queryFn({ pageParam }: QueryFunctionArgs) {
		pageParam = { ...pageParam, limit: queryLimit };
		return query(pageParam);
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

	const users = useContext(UsersContext);

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
							{group.data.map((post) => {
								users?.setUser(post.user!.handle, post.user!);
								return <Post data={post} key={post.createdAt} ref={ref} />;
							})}
						</Fragment>
					)
				)
			)}
		</Container>
	);
}
