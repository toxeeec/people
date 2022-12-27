import { QueryKey, useInfiniteQuery } from "@tanstack/react-query";
import { Fragment, useContext, useEffect } from "react";
import { Users as UsersType } from "../models";
import { Container } from "@mantine/core";
import { useInView } from "react-intersection-observer";
import { CenterLoader } from "../components/CenterLoader";
import { UsersContext } from "../context/UsersContext";
import { User } from "./User";

const queryLimit = 10;

interface PaginationParams {
	limit?: number;
	before?: string;
	after?: string;
}

export type Query = (params: PaginationParams) => Promise<UsersType>;

interface PostsProps {
	enabled?: boolean;
	query: Query;
	queryKey: QueryKey;
}

interface QueryFunctionArgs {
	pageParam?: PaginationParams;
}

export const Users = ({ query, queryKey, enabled = true }: PostsProps) => {
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
		<Container px={0}>
			{enabled && status === "loading" ? (
				<CenterLoader />
			) : (
				data?.pages.map((page, i) => (
					<Fragment key={i}>
						{page.data.map((user) => (
							<User key={user.handle} handle={user.handle} ref={ref} />
						))}
					</Fragment>
				))
			)}
			{isFetchingNextPage && <CenterLoader />}
		</Container>
	);
};
