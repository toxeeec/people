import {
	ActionIcon,
	Divider,
	Flex,
	Group,
	UnstyledButton,
} from "@mantine/core";
import { IconHeart, IconMessageCircle2 } from "@tabler/icons";
import { useQueryClient } from "@tanstack/react-query";
import { useCallback, useContext, useState } from "react";
import { PostsContext } from "../../context/PostsContext";
import { UsersContext } from "../../context/UsersContext";
import { QueryKey } from "../../query-key";
import {
	useDeletePostsPostIDLikes,
	usePutPostsPostIDLikes,
} from "../../spec.gen";
import { PostLikes } from "./PostLikes";
import { PostReplyModal } from "./PostReplyModal";

interface MainPostActionsProps {
	id: number;
	handle: string;
}

export const MainPostActions = ({ id, handle }: MainPostActionsProps) => {
	const { posts, setPost } = useContext(PostsContext);
	const { setUser } = useContext(UsersContext);
	const queryClient = useQueryClient();
	const post = posts[id]!;
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
				onSuccess: (postResponse) => {
					setPost(postResponse.data);
					setUser(postResponse.user);
					queryClient.resetQueries({ queryKey: [QueryKey.POSTS] });
					queryClient.resetQueries({ queryKey: [QueryKey.USERS] });
				},
			}
		);
	}, [like, unlike, post, queryClient, setPost, setUser]);

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
						<IconHeart
							size={18}
							fill={post.status?.isLiked ? "currentColor" : "none"}
						/>
					</ActionIcon>
				</Group>
			</Flex>
			<PostLikes opened={likesOpened} setOpened={setLikesOpened} id={id} />
			<PostReplyModal
				opened={replyOpened}
				setOpened={setReplyOpened}
				id={id}
				handle={handle}
				queryKey={[QueryKey.POSTS, QueryKey.REPLIES, id]}
			/>
		</>
	);
};
