import { Avatar, Text, Group, Paper } from "@mantine/core";
import { forwardRef } from "react";
import { Post as PostsData } from "../../models";
import PostActions from "./PostActions";
import ProfileHoverCard from "./ProfileHoverCard";

interface PostProps {
	data: PostsData;
}

const Post = forwardRef<HTMLDivElement, PostProps>(({ data }, ref) => {
	Post.displayName = "Post";
	const { user, content } = data;
	return (
		<Paper ref={ref} p="xs" radius="xs" withBorder>
			<Group align="center">
				<ProfileHoverCard handle={user!.handle}>
					<Avatar radius="xl" size="md" />
				</ProfileHoverCard>
				<ProfileHoverCard handle={user!.handle}>
					<b>{user?.handle}</b>
				</ProfileHoverCard>
			</Group>
			<Text my="xs">{content}</Text>
			<PostActions {...data} />
		</Paper>
	);
});

export default Post;
