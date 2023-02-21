import { Group, Text } from "@mantine/core";
import { Link } from "react-router-dom";
import { Avatar } from "../../Avatar";
import { User } from "../../models";
import { UserHoverCard } from "../UserHoverCard";

interface PostAuthorProps {
	user: User;
}

export const PostAuthor = ({ user }: PostAuthorProps) => {
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
};
