import { QueryKey, useInfiniteQuery } from "@tanstack/react-query";
import { Fragment, useContext, useEffect, useMemo } from "react";
import { UserMessages as UserMessagesType } from "../../models";
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

export type Query = (params: PaginationParams) => Promise<UserMessagesType>;

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
			if (!lastPage.data.meta || lastPage.data.data.length < queryLimit)
				return undefined;
			return { before: lastPage.data.meta?.oldest };
		},
		refetchOnMount: false,
		refetchOnWindowFocus: false,
	});

	useEffect(() => {
		if (inView) {
			fetchNextPage();
		}
	}, [fetchNextPage, inView]);

	return (
		<>
			{(enabled && status === "loading") || !data ? (
				<CenterLoader />
			) : (
				data.pages.map((page, i) => (
					<Fragment key={i}>
						{i + 1 === data.pages.length && <Header user={page.user} />}
						{page.data.data.map((message) => (
							<Message
								ref={ref}
								key={message.id}
								message={message.message}
								own={message.from === handle}
							/>
						))}
					</Fragment>
				))
			)}
			{isFetchingNextPage && <CenterLoader />}
		</>
	);
};
