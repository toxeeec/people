import { ActionIcon, Group, Text } from "@mantine/core";
import { IconHeart, IconMessageCircle2 } from "@tabler/icons";
import {
	QueryKey as QueryKeyType,
	useQueryClient,
} from "@tanstack/react-query";
import { MouseEvent, useCallback, useContext, useState } from "react";
import { PostsContext } from "../../context/PostsContext";
import { UsersContext } from "../../context/UsersContext";
import { QueryKey } from "../../query-key";
import {
	useDeletePostsPostIDLikes,
	usePutPostsPostIDLikes,
} from "../../spec.gen";
import { PostReplyModal } from "./PostReplyModal";

interface PostActionsProps {
	id: number;
	handle: string;
	queryKey: QueryKeyType;
}

export const PostActions = ({ id, handle, queryKey }: PostActionsProps) => {
	const queryClient = useQueryClient();
	const { setUser } = useContext(UsersContext);
	const { posts, setPost } = useContext(PostsContext);
	const { mutate: like } = usePutPostsPostIDLikes({
		mutation: { retry: 1 },
	});
	const { mutate: unlike } = useDeletePostsPostIDLikes({
		mutation: { retry: 1 },
	});
	const post = posts[id]!;
	const handleLike = useCallback(
		(e: MouseEvent) => {
			e.stopPropagation();
			const fn = post.status?.isLiked ? unlike : like;
			fn(
				{ postID: post.id },
				{
					onSuccess: (postResponse) => {
						setPost(postResponse.data);
						setUser(postResponse.user);
						queryClient.resetQueries({
							queryKey: [QueryKey.LIKES, id],
						});
						queryClient.invalidateQueries({
							queryKey: [QueryKey.LIKES, handle],
						});
					},
				}
			);
		},
		[post, setPost, setUser, like, unlike, queryClient, id, handle]
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
					<IconHeart
						size={18}
						fill={post.status?.isLiked ? "currentColor" : "none"}
					/>
				</ActionIcon>
				<Text size="sm">{post.likes}</Text>
			</Group>
			<PostReplyModal
				opened={opened}
				setOpened={setOpened}
				id={post.id}
				handle={handle}
				queryKey={queryKey}
			/>
		</Group>
	);
};
