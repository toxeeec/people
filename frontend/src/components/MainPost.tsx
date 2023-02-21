import { Text, Paper, Group } from "@mantine/core";
import { forwardRef } from "react";
import { PostAuthor } from "./post/PostAuthor";
import { MainPostActions } from "./post/MainPostActions";
import { Images } from "./images/Images";
import { PostEdit } from "./post/PostEdit";
import { useLocation, useNavigate } from "react-router";
import { Post, User } from "../models";

interface MainPostProps {
	post: Post;
	user: User;
}

export const MainPost = forwardRef<HTMLDivElement, MainPostProps>(
	({ post, user }, ref) => {
		const navigate = useNavigate();
		const { key } = useLocation();
		const onSuccess = () => {
			key === "default" ? navigate("home") : navigate(-1);
		};

		return (
			<Paper p="xs" withBorder radius={0} m={-1} ref={ref}>
				<Group position="apart">
					<PostAuthor user={user} />
					<PostEdit
						id={post.id}
						handle={user.handle}
						exact={false}
						onSuccess={onSuccess}
					/>
				</Group>
				<Text my="xs">{post.content}</Text>
				<Images images={post.images} />
				<MainPostActions post={post} user={user} />
			</Paper>
		);
	}
);

MainPost.displayName = "MainPost";
