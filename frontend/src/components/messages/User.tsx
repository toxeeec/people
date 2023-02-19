import { Avatar, Group, Text } from "@mantine/core";
import { forwardRef } from "react";
import { User as UserType } from "../../models";

interface UserProps {
	user: UserType;
	onClick?: (handle: string) => void;
}

export const User = forwardRef<HTMLDivElement, UserProps>(
	({ user, onClick }, ref) => {
		return (
			<Group
				ref={ref}
				onClick={() => onClick && onClick(user.handle)}
				align="stretch"
				p="md"
				style={{ cursor: "pointer" }}
			>
				<Avatar radius="xl" size="lg" />
				<Text weight="bold">@{user.handle}</Text>
			</Group>
		);
	}
);

User.displayName = "User";
