import { Text, Paper, Group } from "@mantine/core";
import { QueryKey } from "@tanstack/react-query";
import { forwardRef, useContext } from "react";
import { useNavigate } from "react-router-dom";
import { PostsContext } from "../context/PostsContext";
import { Images } from "./images/Images";
import { PostActions } from "./post/PostActions";
import { PostAuthor } from "./post/PostAuthor";

interface PostProps {
	id: number;
	handle: string;
	queryKey: QueryKey;
}

export const Post = forwardRef<HTMLDivElement, PostProps>(
	({ id, handle, queryKey }, ref) => {
		const navigate = useNavigate();
		const { posts } = useContext(PostsContext);
		const post = posts[id]!;

		return (
			post && (
				<Paper
					onClick={() => {
						navigate(`/${handle}/${id}`);
					}}
					ref={ref}
					p="xs"
					m={-1}
					withBorder
					radius={0}
					style={{ cursor: "pointer" }}
				>
					<Group>
						<PostAuthor handle={handle} />
					</Group>
					<Text my="xs">{post.content}</Text>
					<Images images={post.images} />
					<PostActions id={id} handle={handle} queryKey={queryKey} />
				</Paper>
			)
		);
	}
);

Post.displayName = "Post";
