import { Avatar, Group, Text } from "@mantine/core";
import { Link } from "react-router-dom";
import { UserHoverCard } from "../UserHoverCard";

interface PostAuthorProps {
	handle: string;
}

export const PostAuthor = ({ handle }: PostAuthorProps) => {
	return (
		<Group align="center">
			<UserHoverCard handle={handle}>
				<Avatar radius="xl" size="md" component={Link} to={`/${handle}`} />
			</UserHoverCard>
			<UserHoverCard handle={handle}>
				<Text component={Link} to={`/${handle}`} weight="bold">
					{handle}
				</Text>
			</UserHoverCard>
		</Group>
	);
};
