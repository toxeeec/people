import { Container } from "@mantine/core";
import { forwardRef, useCallback } from "react";
import { Users } from "./Users";
import { Query as UsersQuery } from "../../components/messages/Users";
import { getUsersSearch } from "../../spec.gen";

interface SearchUsersProps {
	hidden: boolean;
	debounced: string;
	onClick: (handle: string) => void;
}

export const SearchUsers = forwardRef<HTMLDivElement, SearchUsersProps>(
	({ hidden, debounced, onClick }, ref) => {
		const usersQuery: UsersQuery = useCallback(
			(params) => getUsersSearch({ query: debounced, ...params }),
			[debounced]
		);

		return (
			<Container
				ref={ref}
				p={0}
				w="100%"
				hidden={hidden}
				mih="calc(100% - 60px)"
			>
				<Users
					queryKey={["users", debounced, "messages"]}
					enabled={debounced.length > 0}
					query={usersQuery}
					onClick={onClick}
				/>
			</Container>
		);
	}
);

SearchUsers.displayName = "SearchUsers";