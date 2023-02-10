import { Avatar, Group, Text } from "@mantine/core";
import { Link } from "react-router-dom";
import { User } from "../../models";
import { UserHoverCard } from "../UserHoverCard";

interface PostAuthorProps {
	user: User;
}

export const PostAuthor = ({ user }: PostAuthorProps) => {
	return (
		<Group align="center">
			<UserHoverCard user={user}>
				<Avatar radius="xl" size="md" component={Link} to={`/${user.handle}`} />
			</UserHoverCard>
			<UserHoverCard user={user}>
				<Text component={Link} to={`/${user.handle}`} weight="bold">
					{user.handle}
				</Text>
			</UserHoverCard>
		</Group>
	);
};
