import { type Users } from "@/models";
import { type QueryKey, useInfiniteQuery } from "@tanstack/react-query";
import { Fragment, useEffect } from "react";
import { useInView } from "react-intersection-observer";
import { CenterLoader } from "@/components/utils";
import { User } from "@/components/user/User";

const LIMIT = 10;

type PaginationParams = {
	limit?: number;
	before?: string;
	after?: string;
};

export type UsersQuery = (params: PaginationParams) => Promise<Users>;

type UsersProps = {
	query: UsersQuery;
	queryKey: QueryKey;
	enabled?: boolean;
};

type QueryFunctionArgs = {
	pageParam?: PaginationParams;
};

export function InfiniteUsers({ query, queryKey, enabled = true }: UsersProps) {
	const { ref, inView } = useInView();
	const queryFn = ({ pageParam }: QueryFunctionArgs) => {
		pageParam = { ...pageParam, limit: LIMIT };
		return query(pageParam);
	};

	const { status, data, isFetchingNextPage, fetchNextPage } = useInfiniteQuery({
		queryKey,
		queryFn,
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
				data?.pages.map((page, i) => (
					<Fragment key={i}>
						{page.data.map((user) => (
							<User key={user.handle} user={user} ref={ref} />
						))}
					</Fragment>
				))
			)}
			{isFetchingNextPage && <CenterLoader />}
		</>
	);
}
