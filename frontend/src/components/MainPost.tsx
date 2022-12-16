import { Text, Paper } from "@mantine/core";
import { useContext } from "react";
import { PostAuthor } from "./post/PostAuthor";
import { PostsContext } from "../context/PostsContext";
import { MainPostActions } from "./post/MainPostActions";

interface MainPostProps {
	id: number;
	handle: string;
}

export const MainPost = ({ id, handle }: MainPostProps) => {
	const { posts } = useContext(PostsContext);
	const post = posts[id]!;

	return (
		post && (
			<Paper p="xs" radius="xs" withBorder>
				<PostAuthor handle={handle} />
				<Text my="xs">{post.content}</Text>
				<MainPostActions id={id} handle={handle} />
			</Paper>
		)
	);
};
