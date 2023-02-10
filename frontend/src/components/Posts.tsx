import { QueryKey, useInfiniteQuery } from "@tanstack/react-query";
import { Fragment, useEffect } from "react";
import { PostsResponse } from "../models";
import { Post } from "../components/Post";
import { useInView } from "react-intersection-observer";
import { CenterLoader } from "../components/CenterLoader";

const queryLimit = 10;

interface PaginationParams {
	limit?: number;
	before?: number;
	after?: number;
}

export type Query = (params: PaginationParams) => Promise<PostsResponse>;

interface PostsProps {
	enabled?: boolean;
	query: Query;
	queryKey: QueryKey;
}

interface QueryFunctionArgs {
	pageParam?: PaginationParams;
}

export const Posts = ({ query, queryKey, enabled = true }: PostsProps) => {
	const { ref, inView } = useInView();

	const queryFn = ({ pageParam }: QueryFunctionArgs) => {
		pageParam = { ...pageParam, limit: queryLimit };
		return query(pageParam);
	};

	const { status, data, isFetchingNextPage, fetchNextPage } = useInfiniteQuery({
		queryFn,
		queryKey,
		enabled,
		getNextPageParam: (lastPage) => {
			if (!lastPage.meta || lastPage.data.length < queryLimit) return undefined;
			return { before: lastPage.meta?.oldest };
		},
	});

	useEffect(() => {
		if (inView) {
			fetchNextPage();
		}
	}, [fetchNextPage, inView]);

	return (
		<>
			{enabled && status === "loading" ? (
				<CenterLoader />
			) : (
				data?.pages?.map((page, i) => (
					<Fragment key={i}>
						{page.data?.map(({ data, user }) => (
							<Post key={data.id} ref={ref} post={data} user={user} />
						))}
					</Fragment>
				))
			)}
			{isFetchingNextPage && <CenterLoader />}
		</>
	);
};
