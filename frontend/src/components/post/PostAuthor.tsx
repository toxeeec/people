import { Avatar, Group, Text } from "@mantine/core";
import { Link } from "react-router-dom";
import { stopPropagation } from "../../utils";
import { ProfileHoverCard } from "../ProfileHoverCard";

interface PostAuthorProps {
	handle: string;
}

export const PostAuthor = ({ handle }: PostAuthorProps) => {
	return (
		<Group align="center">
			<ProfileHoverCard handle={handle}>
				<Avatar
					radius="xl"
					size="md"
					component={Link}
					to={`/${handle}`}
					onClick={stopPropagation}
				/>
			</ProfileHoverCard>
			<ProfileHoverCard handle={handle}>
				<Text
					component={Link}
					to={`/${handle}`}
					weight="bold"
					onClick={stopPropagation}
				>
					{handle}
				</Text>
			</ProfileHoverCard>
		</Group>
	);
};
