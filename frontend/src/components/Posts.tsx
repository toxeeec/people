import { QueryKey, useInfiniteQuery } from "@tanstack/react-query";
import { Fragment, useContext, useEffect } from "react";
import { PostsResponse, User } from "../models";
import { Container } from "@mantine/core";
import PostComponent from "../components/post/Post";
import { useInView } from "react-intersection-observer";
import CenterLoader from "../components/CenterLoader";
import UsersContext from "../context/UsersContext";

const queryLimit = 10;

interface PaginationParams {
	limit?: number;
	before?: number;
	after?: number;
}

export type Query = (_params: PaginationParams) => Promise<PostsResponse>;

interface PostsProps {
	query: Query;
	queryKey: QueryKey;
	user?: User;
}

interface QueryFunctionArgs {
	pageParam?: PaginationParams;
}

function Posts({ query, user, queryKey }: PostsProps) {
	const usersCtx = useContext(UsersContext);
	const { ref, inView } = useInView();

	function queryFn({ pageParam }: QueryFunctionArgs) {
		pageParam = { ...pageParam, limit: queryLimit };
		return query(pageParam);
	}
	const {
		isLoading,
		data,
		hasNextPage,
		isFetching,
		fetchNextPage,
		isRefetching,
	} = useInfiniteQuery({
		queryKey,
		queryFn,
		getNextPageParam: (lastPage) => {
			if (!lastPage.meta) return undefined;
			return { before: lastPage.meta?.oldest };
		},
	});

	useEffect(() => {
		if (inView && hasNextPage) {
			fetchNextPage();
		}
	}, [fetchNextPage, inView, hasNextPage]);

	return (
		<Container px={0}>
			{isLoading || isRefetching ? (
				<CenterLoader />
			) : (
				data?.pages?.map((group, i) =>
					i > 0 && i === data.pages.length - 1 && isFetching ? (
						<CenterLoader key={i} />
					) : (
						<Fragment key={i}>
							{group.data?.map((postResponse) => {
								return (
									<PostComponent
										post={postResponse}
										key={postResponse.data.createdAt}
										ref={ref}
									/>
								);
							})}
						</Fragment>
					)
				)
			)}
		</Container>
	);
}

export default Posts;
