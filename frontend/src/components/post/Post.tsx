import { Images } from "@/components/images";
import { PostActions, PostAuthor, PostEdit } from "@/components/post";
import { type Post as PostType, type User } from "@/models";
import { Group, Paper, Text } from "@mantine/core";
import { forwardRef } from "react";
import { useNavigate } from "react-router-dom";

type PostProps = {
	post: PostType;
	user: User;
};

export const Post = forwardRef<HTMLDivElement, PostProps>(({ post, user }, ref) => {
	const navigate = useNavigate();
	return (
		<Paper
			onClick={(e) => {
				if (e.target === e.currentTarget) {
					navigate(`/${user.handle}/${post.id}`);
				}
			}}
			ref={ref}
			p="xs"
			m={-1}
			withBorder
			radius={0}
			style={{ cursor: "pointer" }}
		>
			<Group position="apart">
				<PostAuthor user={user} />
				<PostEdit postID={post.id} handle={user.handle} />
			</Group>
			<Text my="xs">{post.content}</Text>
			<Images images={post.images} />
			<PostActions post={post} user={user} />
		</Paper>
	);
});

Post.displayName = "Post";
