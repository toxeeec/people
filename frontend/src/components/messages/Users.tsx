import { QueryKey, useInfiniteQuery } from "@tanstack/react-query";
import { Fragment, useEffect } from "react";
import { Users as UsersType } from "../../models";
import { useInView } from "react-intersection-observer";
import { User } from "./User";
import { CenterLoader } from "../CenterLoader";

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
	onClick: (handle: string) => void;
	enabled?: boolean;
}

interface QueryFunctionArgs {
	pageParam?: PaginationParams;
}

export const Users = ({
	query,
	queryKey,
	onClick,
	enabled = true,
}: UsersProps) => {
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
		if (inView) {
			fetchNextPage();
		}
	}, [fetchNextPage, inView]);

	return (
		<>
			{enabled && status === "loading" ? (
				<CenterLoader />
			) : (
				data?.pages.map((page, i) => (
					<Fragment key={i}>
						{page.data.map((user) => (
							<User
								key={user.handle}
								handle={user.handle}
								ref={ref}
								onClick={onClick}
							/>
						))}
					</Fragment>
				))
			)}
			{isFetchingNextPage && <CenterLoader />}
		</>
	);
};
