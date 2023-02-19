import { Avatar, Button, Group, Text } from "@mantine/core";
import { Link } from "react-router-dom";
import { User } from "../../models";

interface PostAuthorProps {
	user: User;
}

export const UserInfo = ({ user }: PostAuthorProps) => {
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
				<Avatar radius="xl" size="md" />
				<Text weight="bold">{user.handle}</Text>
			</Group>
		</Button>
	);
};
