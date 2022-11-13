import { Avatar, Text, Group, Paper } from "@mantine/core";
import { forwardRef } from "react";
import { Post as PostData } from "../models";
import PostActions from "./PostActions";

interface PostProps {
	data: PostData;
}

const Post = forwardRef<HTMLDivElement, PostProps>(({ data }, ref) => {
	Post.displayName = "Post";
	const { user, content } = data;
	return (
		<Paper ref={ref} p="xs" radius="xs" withBorder>
			<Group align="center">
				<Avatar radius="xl" />
				<b>{user?.handle}</b>
			</Group>
			<Text my="xs">{content}</Text>
			<PostActions {...data} />
		</Paper>
	);
});

export default Post;
