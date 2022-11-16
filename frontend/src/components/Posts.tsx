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

function Posts({ query }: PostsProps) {
	const usersCtx = useContext(UsersContext);
	const { ref, inView } = useInView();

	function queryFn({ pageParam }: QueryFunctionArgs) {
		pageParam = { ...pageParam, limit: queryLimit };
		return query(pageParam);
	}
	const { isLoading, data, hasNextPage, isFetching, fetchNextPage } =
		useInfiniteQuery({
			queryKey: ["feed"],
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
			{isLoading ? (
				<CenterLoader />
			) : (
				data?.pages.map((group, i) =>
					i > 0 && i === data.pages.length - 1 && isFetching ? (
						<CenterLoader key={i} />
					) : (
						<Fragment key={i}>
							{group.data.map((post) => {
								usersCtx?.setUser(post.user!.handle, post.user!);
								return <Post data={post} key={post.createdAt} ref={ref} />;
							})}
						</Fragment>
					)
				)
			)}
		</Container>
	);
}

export default Posts;
