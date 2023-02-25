import { Button, Group, Text } from "@mantine/core";
import { Link } from "react-router-dom";
import { Avatar } from "@/components/user";
import { type User } from "@/models";

type PostAuthorProps = {
	user: User;
};

export function UserInfo({ user }: PostAuthorProps) {
	return (
		<Button
			component={Link}
			to={`/${user.handle}`}
			variant="default"
			radius={0}
			style={{ border: "none" }}
			h="100%"
			bg="none"
		>
			<Group>
				<Avatar user={user} size="md" />
				<Text weight="bold">{user.handle}</Text>
			</Group>
		</Button>
	);
}
