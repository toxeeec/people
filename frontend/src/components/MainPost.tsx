import { Text, Paper, Group } from "@mantine/core";
import { forwardRef, useContext } from "react";
import { PostAuthor } from "./post/PostAuthor";
import { PostsContext } from "../context/PostsContext";
import { MainPostActions } from "./post/MainPostActions";
import { Images } from "./images/Images";
import { PostEdit } from "./post/PostEdit";
import { useLocation, useNavigate } from "react-router";
import { QueryKey } from "../query-key";

interface MainPostProps {
	id: number;
	handle: string;
}

export const MainPost = forwardRef<HTMLDivElement, MainPostProps>(
	({ id, handle }, ref) => {
		const { posts } = useContext(PostsContext);
		const post = posts[id]!;
		const navigate = useNavigate();
		const { key } = useLocation();
		const onSuccess = () => {
			key === "default" ? navigate("home") : navigate(-1);
		};

		return (
			post && (
				<Paper p="xs" withBorder radius={0} m={-1} ref={ref}>
					<Group position="apart">
						<PostAuthor handle={handle} />
						<PostEdit
							id={id}
							handle={handle}
							queryKey={[QueryKey.POSTS]}
							exact={false}
							onSuccess={onSuccess}
						/>
					</Group>
					<Text my="xs">{post.content}</Text>
					<Images images={post.images} />
					<MainPostActions id={id} handle={handle} />
				</Paper>
			)
		);
	}
);

MainPost.displayName = "MainPost";
