import { ActionIcon, Group, Text } from "@mantine/core";
import { IconHeart, IconMessageCircle2 } from "@tabler/icons";
import { useQueryClient } from "@tanstack/react-query";
import { useCallback, useState } from "react";
import { type Post, type User } from "@/models";
import { useDeletePostsPostIDLikes, usePutPostsPostIDLikes } from "@/spec.gen";
import { PostReplyModal } from "@/components/post/PostReplyModal";

type PostActionsProps = {
	post: Post;
	user: User;
};

export function PostActions({ post, user }: PostActionsProps) {
	const queryClient = useQueryClient();
	const { mutate: like } = usePutPostsPostIDLikes({
		mutation: { retry: 1 },
	});
	const { mutate: unlike } = useDeletePostsPostIDLikes({
		mutation: { retry: 1 },
	});
	const isLiked = post.status?.isLiked;
	const handleLike = useCallback(() => {
		(isLiked ? unlike : like)(
			{ postID: post.id },
			{ onSuccess: () => queryClient.invalidateQueries() }
		);
	}, [isLiked, post, like, unlike, queryClient]);
	const [opened, setOpened] = useState(false);

	const handleOpen = () => {
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
					<IconHeart size={18} fill={isLiked ? "currentColor" : "none"} />
				</ActionIcon>
				<Text size="sm">{post.likes}</Text>
			</Group>
			<PostReplyModal
				post={post}
				opened={opened}
				handleClose={() => setOpened(false)}
				user={user}
			/>
		</Group>
	);
}
