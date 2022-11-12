import { Avatar, Text, Group, Paper } from "@mantine/core";
import { forwardRef } from "react";
import { User } from "../models";

interface PostProps {
	author: User;
	content: string;
}

const Post = forwardRef<HTMLDivElement, PostProps>(
	({ author, content }, ref) => {
		Post.displayName = "Post";
		return (
			<Paper ref={ref} p="xs" radius="xs" withBorder>
				<Group align="center">
					<Avatar radius="xl" />
					<b>{author.handle}</b>
				</Group>
				<Text>{content}</Text>
			</Paper>
		);
	}
);

export default Post;
