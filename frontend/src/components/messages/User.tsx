import { Avatar, Group, Text } from "@mantine/core";
import { forwardRef } from "react";

interface UserProps {
	handle: string;
	onClick: (handle: string) => void;
}

export const User = forwardRef<HTMLDivElement, UserProps>(
	({ handle, onClick }, ref) => {
		return (
			<Group
				ref={ref}
				onClick={() => onClick(handle)}
				align="stretch"
				p="md"
				style={{ cursor: "pointer" }}
			>
				<Avatar radius="xl" size="lg" />
				<Text weight="bold">@{handle}</Text>
			</Group>
		);
	}
);

User.displayName = "User";
