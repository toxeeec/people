import { Text, Paper, Group } from "@mantine/core";
import { forwardRef } from "react";
import { useNavigate } from "react-router-dom";
import { Post as PostType, User } from "../models";
import { Images } from "./images/Images";
import { PostActions } from "./post/PostActions";
import { PostAuthor } from "./post/PostAuthor";
import { PostEdit } from "./post/PostEdit";

interface PostProps {
	post: PostType;
	user: User;
}

export const Post = forwardRef<HTMLDivElement, PostProps>(
	({ post, user }, ref) => {
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
					<PostEdit id={post.id} handle={user.handle} exact={true} />
				</Group>
				<Text my="xs" style={{ display: "inline-block" }}>
					{post.content}
				</Text>
				<Images images={post.images} />
				<PostActions post={post} handle={user.handle} />
			</Paper>
		);
	}
);

Post.displayName = "Post";
