import { QueryKey, useInfiniteQuery } from "@tanstack/react-query";
import { Fragment, useContext, useEffect } from "react";
import { Users } from "../models";
import { Container } from "@mantine/core";
import { useInView } from "react-intersection-observer";
import CenterLoader from "../components/CenterLoader";
import UsersContext from "../context/UsersContext";
import Profile from "./Profile";

const queryLimit = 10;

interface PaginationParams {
	limit?: number;
	before?: string;
	after?: string;
}

export type Query = (_params: PaginationParams) => Promise<Users>;

interface PostsProps {
	query: Query;
	queryKey: QueryKey;
}

interface QueryFunctionArgs {
	pageParam?: PaginationParams;
}

function Posts({ query, queryKey }: PostsProps) {
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
			if (!lastPage.meta || lastPage.data.length < queryLimit) return undefined;
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
				data?.pages.map((group, i) =>
					i > 0 && i === data.pages.length - 1 && isFetching ? (
						<CenterLoader key={i} />
					) : (
						<Fragment key={i}>
							{group.data.map((user) => {
								{
									usersCtx?.setUser(user.handle, user);
								}
								return <Profile key={user.handle} user={user} ref={ref} />;
							})}
						</Fragment>
					)
				)
			)}
		</Container>
	);
}

export default Posts;
