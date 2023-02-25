import { type QueryKey, useInfiniteQuery } from "@tanstack/react-query";
import { Fragment, useContext, useEffect, useMemo } from "react";
import { type Messages } from "@/models";
import { useInView } from "react-intersection-observer";
import { CenterLoader } from "@/components/utils";
import { Message } from "@/components/messages/Message";
import { AuthContext } from "@/context/AuthContext";

const LIMIT = 20;

type PaginationParams = {
	limit?: number;
	before?: number;
	after?: number;
};

export type Query = (params: PaginationParams) => Promise<Messages>;

type UserMessagesProps = {
	query: Query;
	queryKey: QueryKey;
	enabled?: boolean;
};

type QueryFunctionArgs = {
	pageParam?: PaginationParams;
};

export function UserMessages({ query, queryKey, enabled = true }: UserMessagesProps) {
	const { ref, inView } = useInView();
	const { getAuth } = useContext(AuthContext);
	const handle = useMemo(() => getAuth().handle, [getAuth]);
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
		refetchOnMount: false,
		refetchOnWindowFocus: false,
	});

	useEffect(() => {
		if (inView && !isFetchingNextPage) {
			fetchNextPage();
		}
	}, [fetchNextPage, inView, isFetchingNextPage]);

	return (
		<>
			{(enabled && status === "loading") || !data ? (
				<CenterLoader />
			) : (
				data.pages.map((page, i) => (
					<Fragment key={i}>
						{page.data.map((message) => (
							<Message
								ref={ref}
								key={message.id}
								message={message}
								own={message.from.handle === handle}
							/>
						))}
					</Fragment>
				))
			)}
			{isFetchingNextPage && <CenterLoader />}
		</>
	);
}
