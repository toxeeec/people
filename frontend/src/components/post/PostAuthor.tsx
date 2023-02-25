import { Avatar, UserHoverCard } from "@/components/user";
import { type User } from "@/models";
import { Group, Text } from "@mantine/core";
import { Link } from "react-router-dom";

type PostAuthorProps = {
	user: User;
};

export function PostAuthor({ user }: PostAuthorProps) {
	return (
		<Group align="center">
			<UserHoverCard user={user}>
				<Group>
					<Link to={`/${user.handle}`}>
						<Avatar user={user} size="md" />
					</Link>
					<Text component={Link} to={`/${user.handle}`} weight="bold">
						{user.handle}
					</Text>
				</Group>
			</UserHoverCard>
		</Group>
	);
}
