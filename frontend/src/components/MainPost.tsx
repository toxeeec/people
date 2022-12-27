import { Text, Paper } from "@mantine/core";
import { forwardRef, useContext } from "react";
import { PostAuthor } from "./post/PostAuthor";
import { PostsContext } from "../context/PostsContext";
import { MainPostActions } from "./post/MainPostActions";
import { Images } from "./images/Images";

interface MainPostProps {
	id: number;
	handle: string;
}

export const MainPost = forwardRef<HTMLDivElement, MainPostProps>(
	({ id, handle }, ref) => {
		const { posts } = useContext(PostsContext);
		const post = posts[id]!;

		return (
			post && (
				<Paper p="xs" radius="xs" withBorder ref={ref}>
					<PostAuthor handle={handle} />
					<Text my="xs">{post.content}</Text>
					<Images images={post.images} />
					<MainPostActions id={id} handle={handle} />
				</Paper>
			)
		);
	}
);

MainPost.displayName = "MainPost";
