import { Avatar, Text, Group, Paper } from "@mantine/core";
import { forwardRef } from "react";
import { Link } from "react-router-dom";
import { Post as PostsData } from "../../models";
import PostActions from "./PostActions";
import ProfileHoverCard from "./ProfileHoverCard";

interface PostProps {
	data: PostsData;
}

const Post = forwardRef<HTMLDivElement, PostProps>(({ data }, ref) => {
	Post.displayName = "Post";
	const { content, user } = data;
	return (
		<Paper ref={ref} p="xs" radius="xs" withBorder>
			<Group align="center">
				<ProfileHoverCard handle={user!.handle}>
					<Avatar
						radius="xl"
						size="md"
						component={Link}
						to={`/${user!.handle}`}
					/>
				</ProfileHoverCard>
				<ProfileHoverCard handle={user!.handle}>
					<Text component={Link} to={`/${user!.handle}`} weight="bold">
						{user?.handle}
					</Text>
				</ProfileHoverCard>
			</Group>
			<Text my="xs">{content}</Text>
			<PostActions {...data} />
		</Paper>
	);
});

export default Post;
