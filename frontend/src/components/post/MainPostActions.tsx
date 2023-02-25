import { ActionIcon, Divider, Flex, Group, UnstyledButton } from "@mantine/core";
import { IconHeart, IconMessageCircle2 } from "@tabler/icons";
import { useQueryClient } from "@tanstack/react-query";
import { useCallback, useState } from "react";
import { type Post, type User } from "@/models";
import { useDeletePostsPostIDLikes, usePutPostsPostIDLikes } from "@/spec.gen";
import { PostLikes } from "@/components/post/PostLikes";
import { PostReplyModal } from "@/components/post/PostReplyModal";

type MainPostActionsProps = {
	post: Post;
	user: User;
};

export function MainPostActions({ post, user }: MainPostActionsProps) {
	const queryClient = useQueryClient();
	const [likesOpened, setLikesOpened] = useState(false);
	const [replyOpened, setReplyOpened] = useState(false);

	const { mutate: like } = usePutPostsPostIDLikes({
		mutation: {
			retry: 1,
		},
	});
	const { mutate: unlike } = useDeletePostsPostIDLikes({
		mutation: {
			retry: 1,
		},
	});
	const handleLike = useCallback(() => {
		const fn = post.status?.isLiked ? unlike : like;
		fn(
			{ postID: post.id },
			{
				onSuccess: () => {
					queryClient.invalidateQueries();
				},
			}
		);
	}, [like, unlike, post, queryClient]);

	return (
		<>
			<Divider />
			<UnstyledButton my="md" onClick={() => setLikesOpened(true)}>
				<b>{post.likes}</b>
				{post.likes === 1 ? " Like" : " Likes"}
			</UnstyledButton>
			<Divider mb="xs" />
			<Flex justify="space-around">
				<Group align="center" spacing="xs">
					<ActionIcon onClick={() => setReplyOpened(true)}>
						<IconMessageCircle2 size={18} />
					</ActionIcon>
				</Group>
				<Group align="center" spacing="xs">
					<ActionIcon onClick={handleLike}>
						<IconHeart size={18} fill={post.status?.isLiked ? "currentColor" : "none"} />
					</ActionIcon>
				</Group>
			</Flex>
			<PostLikes opened={likesOpened} setOpened={setLikesOpened} id={post.id} />
			<PostReplyModal
				opened={replyOpened}
				handleClose={() => setReplyOpened(false)}
				post={post}
				user={user}
			/>
		</>
	);
}
