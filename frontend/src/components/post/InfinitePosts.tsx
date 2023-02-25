import { type PostsResponse } from "@/models";
import { type QueryKey, useInfiniteQuery } from "@tanstack/react-query";
import { Fragment, useEffect } from "react";
import { useInView } from "react-intersection-observer";
import { CenterLoader } from "@/components/utils";
import { Post } from "@/components/post";

const LIMIT = 10;

type PaginationParams = {
	limit?: number;
	before?: number;
	after?: number;
};

export type PostsQuery = (params: PaginationParams) => Promise<PostsResponse>;

type PostsProps = {
	enabled?: boolean;
	query: PostsQuery;
	queryKey: QueryKey;
};

type QueryFunctionArgs = {
	pageParam?: PaginationParams;
};

export function InfinitePosts({ query, queryKey, enabled = true }: PostsProps) {
	const { ref, inView } = useInView();

	const queryFn = ({ pageParam }: QueryFunctionArgs) => {
		pageParam = { ...pageParam, limit: LIMIT };
		return query(pageParam);
	};

	const { status, data, isFetchingNextPage, fetchNextPage } = useInfiniteQuery({
		queryFn,
		queryKey,
		enabled,
		getNextPageParam: (lastPage) => {
			if (!lastPage.meta || lastPage.data.length < LIMIT) return undefined;
			return { before: lastPage.meta?.oldest };
		},
	});

	useEffect(() => {
		if (inView && !isFetchingNextPage) {
			fetchNextPage();
		}
	}, [fetchNextPage, inView, isFetchingNextPage]);

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
}
