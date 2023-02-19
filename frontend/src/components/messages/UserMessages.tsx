import { QueryKey, useInfiniteQuery } from "@tanstack/react-query";
import { Fragment, useContext, useEffect, useMemo } from "react";
import { Messages } from "../../models";
import { useInView } from "react-intersection-observer";
import { CenterLoader } from "../../components/CenterLoader";
import { Message } from "./Message";
import { AuthContext } from "../../context/AuthContext";
import { Header } from "./Header";

const queryLimit = 20;

interface PaginationParams {
	limit?: number;
	before?: number;
	after?: number;
}

export type Query = (params: PaginationParams) => Promise<Messages>;

interface UserMessagesProps {
	query: Query;
	queryKey: QueryKey;
	enabled?: boolean;
}

interface QueryFunctionArgs {
	pageParam?: PaginationParams;
}

export const UserMessages = ({
	query,
	queryKey,
	enabled = true,
}: UserMessagesProps) => {
	const { ref, inView } = useInView();
	const { getAuth } = useContext(AuthContext);
	const handle = useMemo(() => getAuth().handle, [getAuth]);
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
};
