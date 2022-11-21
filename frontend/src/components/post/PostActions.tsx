import { ActionIcon, Group, Text } from "@mantine/core";
import { IconHeart, IconMessageCircle2 } from "@tabler/icons";
import {
	Dispatch,
	MouseEvent,
	SetStateAction,
	useCallback,
	useState,
} from "react";
import { Post } from "../../models/post";
import {
	useDeletePostsPostIDLikes,
	usePutPostsPostIDLikes,
} from "../../spec.gen";
import PostReply from "./PostReply";

interface PostActionsProps {
	post: Post;
	setPost: Dispatch<SetStateAction<Post>>;
}

export default function PostActions({ post, setPost }: PostActionsProps) {
	const { mutate: like } = usePutPostsPostIDLikes({ mutation: { retry: 1 } });
	const { mutate: unlike } = useDeletePostsPostIDLikes({
		mutation: { retry: 1 },
	});
	const handleLike = useCallback(
		(e: MouseEvent) => {
			e.stopPropagation();
			const fn = post.isLiked ? unlike : like;
			fn(
				{ postID: post.id },
				{
					onSuccess({ likes, isLiked }) {
						setPost((p) => ({ ...p, likes, isLiked }));
					},
				}
			);
		},
		[post, setPost, like, unlike]
	);

	const [opened, setOpened] = useState(false);

	const handleOpen = (e: MouseEvent) => {
		e.stopPropagation();
		setOpened(true);
	};

	return (
		<Group position="apart" align="center" mx="30%">
			<Group align="center" spacing="xs">
				<ActionIcon onClick={handleOpen}>
					<IconMessageCircle2 size={18} />
				</ActionIcon>
				<Text size="sm">{post.replies}</Text>
			</Group>
			<Group align="center" spacing="xs">
				<ActionIcon onClick={handleLike}>
					<IconHeart size={18} fill={post.isLiked ? "currentColor" : "none"} />
				</ActionIcon>
				<Text size="sm">{post.likes}</Text>
			</Group>
			<PostReply
				opened={opened}
				setOpened={setOpened}
				isReply={true}
				post={post}
				setPost={setPost}
			/>
		</Group>
	);
}
