import { QueryKey, useInfiniteQuery } from "@tanstack/react-query";
import { Fragment, useContext, useEffect } from "react";
import { Users as UsersType } from "../../models";
import { useInView } from "react-intersection-observer";
import { User } from "./User";
import { UsersContext } from "../../context/UsersContext";
import { CenterLoader } from "../CenterLoader";

const queryLimit = 10;

interface PaginationParams {
	limit?: number;
	before?: string;
	after?: string;
}

export type Query = (params: PaginationParams) => Promise<UsersType>;

interface UsersProps {
	enabled?: boolean;
	query: Query;
	queryKey: QueryKey;
	onClick: (handle: string) => void;
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
	const { setUser } = useContext(UsersContext);
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
		onSuccess: (data) => {
			data.pages.forEach((users) =>
				users.data.forEach((user) => setUser(user))
			);
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
