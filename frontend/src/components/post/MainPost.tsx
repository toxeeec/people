import { Text, Paper, Group } from "@mantine/core";
import { forwardRef } from "react";
import { PostAuthor } from "@/components/post";
import { MainPostActions } from "@/components/post";
import { Images } from "@/components/images";
import { PostEdit } from "@/components/post";
import { useLocation, useNavigate } from "react-router";
import { type Post, type User } from "@/models";

type MainPostProps = {
	post: Post;
	user: User;
};

export const MainPost = forwardRef<HTMLDivElement, MainPostProps>(({ post, user }, ref) => {
	const navigate = useNavigate();
	const { key } = useLocation();
	const onSuccess = () => {
		key === "default" ? navigate("home") : navigate(-1);
	};

	return (
		<Paper p="xs" withBorder radius={0} m={-1} ref={ref}>
			<Group position="apart">
				<PostAuthor user={user} />
				<PostEdit postID={post.id} handle={user.handle} onSuccess={onSuccess} />
			</Group>
			<Text my="xs">{post.content}</Text>
			<Images images={post.images} />
			<MainPostActions post={post} user={user} />
		</Paper>
	);
});

MainPost.displayName = "MainPost";
