import { QueryKey, useInfiniteQuery } from "@tanstack/react-query";
import { Fragment, useContext, useEffect } from "react";
import { PostsResponse } from "../models";
import { Container } from "@mantine/core";
import { Post } from "../components/Post";
import { useInView } from "react-intersection-observer";
import { CenterLoader } from "../components/CenterLoader";
import { UsersContext } from "../context/UsersContext";
import { PostsContext } from "../context/PostsContext";

const queryLimit = 10;

interface PaginationParams {
	limit?: number;
	before?: number;
	after?: number;
}

export type Query = (params: PaginationParams) => Promise<PostsResponse>;

interface PostsProps {
	query: Query;
	queryKey: QueryKey;
}

interface QueryFunctionArgs {
	pageParam?: PaginationParams;
}

export const Posts = ({ query, queryKey }: PostsProps) => {
	const { ref, inView } = useInView();
	const { setUser } = useContext(UsersContext);
	const { setPost } = useContext(PostsContext);

	const queryFn = ({ pageParam }: QueryFunctionArgs) => {
		pageParam = { ...pageParam, limit: queryLimit };
		return query(pageParam);
	};

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
		onSuccess: (data) => {
			data.pages.forEach((postsResponse) =>
				postsResponse.data.forEach((postResponse) => {
					setPost(postResponse.data);
					setUser(postResponse.user);
				})
			);
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
							{group.data?.map((postResponse) => (
								<Post
									key={postResponse.data.createdAt}
									ref={ref}
									id={postResponse.data.id}
									handle={postResponse.user.handle}
									queryKey={queryKey}
								/>
							))}
						</Fragment>
					)
				)
			)}
		</Container>
	);
};
