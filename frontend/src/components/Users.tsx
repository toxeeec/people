import { QueryKey, useInfiniteQuery } from "@tanstack/react-query";
import { Fragment, useEffect } from "react";
import { Users as UsersType } from "../models";
import { useInView } from "react-intersection-observer";
import { CenterLoader } from "../components/CenterLoader";
import { User } from "./User";

const queryLimit = 10;

interface PaginationParams {
	limit?: number;
	before?: string;
	after?: string;
}

export type Query = (params: PaginationParams) => Promise<UsersType>;

interface UsersProps {
	query: Query;
	queryKey: QueryKey;
	enabled?: boolean;
}

interface QueryFunctionArgs {
	pageParam?: PaginationParams;
}

export const Users = ({ query, queryKey, enabled = true }: UsersProps) => {
	const { ref, inView } = useInView();
	const queryFn = ({ pageParam }: QueryFunctionArgs) => {
		pageParam = { ...pageParam, limit: queryLimit };
		return query(pageParam);
	};

	const { status, data, isFetchingNextPage, fetchNextPage } = useInfiniteQuery({
		queryKey,
		queryFn,
		enabled,
		getNextPageParam: (lastPage) => {
			if (!lastPage.meta || lastPage.data.length < queryLimit) return undefined;
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
};
