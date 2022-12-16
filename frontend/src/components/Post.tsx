import { Text, Paper, Group } from "@mantine/core";
import { forwardRef, useContext } from "react";
import { useNavigate } from "react-router-dom";
import { PostsContext } from "../context/PostsContext";
import { PostActions } from "./post/PostActions";
import { PostAuthor } from "./post/PostAuthor";

interface PostProps {
	id: number;
	handle: string;
}

export const Post = forwardRef<HTMLDivElement, PostProps>(
	({ id, handle }, ref) => {
		const navigate = useNavigate();
		const { posts } = useContext(PostsContext);
		const post = posts[id]!;

		return (
			post && (
				<Paper
					p="xs"
					radius="xs"
					withBorder
					ref={ref}
					onClick={() => {
						navigate(`/${handle}/${id}`);
					}}
					style={{ cursor: "pointer" }}
				>
					<Group>
						<PostAuthor handle={handle} />
					</Group>
					<Text my="xs">{post.content}</Text>
					<PostActions id={id} handle={handle} />
				</Paper>
			)
		);
	}
);

Post.displayName = "Post";
