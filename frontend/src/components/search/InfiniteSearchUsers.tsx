import { Avatar } from "@/components/user";
import { CenterLoader } from "@/components/utils";
import { type User } from "@/models";
import { getUsersSearch } from "@/spec.gen";
import { Container, Group, Text } from "@mantine/core";
import { useInfiniteQuery } from "@tanstack/react-query";
import { forwardRef, Fragment, useEffect } from "react";
import { useInView } from "react-intersection-observer";

const LIMIT = 10;

type UserProps = {
	user: User;
	onClick?: (handle: string) => void;
};

const SearchUser = forwardRef<HTMLDivElement, UserProps>(({ user, onClick }, ref) => {
	return (
		<Group
			ref={ref}
			onClick={() => onClick && onClick(user.handle)}
			align="stretch"
			p="md"
			style={{ cursor: "pointer" }}
		>
			<Avatar user={user} size="lg" />
			<Text weight="bold">@{user.handle}</Text>
		</Group>
	);
});

SearchUser.displayName = "SearchUser";

type InfiniteSearchUsersProps = {
	hidden: boolean;
	query: string;
	onClick: (handle: string) => void;
	enabled: boolean;
};

export const InfiniteSearchUsers = forwardRef<HTMLDivElement, InfiniteSearchUsersProps>(
	({ hidden, query, onClick, enabled }, ref) => {
		const { ref: inViewRef, inView } = useInView();

		const { status, data, isFetchingNextPage, fetchNextPage } = useInfiniteQuery({
			queryKey: ["users", query],
			queryFn: ({ pageParam }) => getUsersSearch({ query, limit: LIMIT, ...pageParam }),
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
			<Container ref={ref} p={0} hidden={hidden} w="100%" h="100%" style={{ overflowY: "auto" }}>
				{status === "loading" && enabled ? (
					<CenterLoader />
				) : (
					data?.pages.map((page, i) => (
						<Fragment key={i}>
							{page.data.map((user) => (
								<SearchUser key={user.handle} user={user} ref={inViewRef} onClick={onClick} />
							))}
						</Fragment>
					))
				)}
				{isFetchingNextPage && <CenterLoader />}
			</Container>
		);
	}
);

InfiniteSearchUsers.displayName = "SearchUsers";
