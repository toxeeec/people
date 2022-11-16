import { ActionIcon, Group, Text } from "@mantine/core";
import { IconHeart, IconMessageCircle2 } from "@tabler/icons";
import { useCallback, useState } from "react";
import {
	useDeletePostsPostIDLikes,
	usePutPostsPostIDLikes,
} from "../../spec.gen";

interface PostActionsProps {
	id: number;
	likes: number;
	replies: number;
	isLiked: boolean;
}

export default function PostActions({
	likes: initialLikes,
	replies,
	isLiked: initialIsLiked,
	id,
}: PostActionsProps) {
	const { mutate: like } = usePutPostsPostIDLikes({ mutation: { retry: 1 } });
	const { mutate: unlike } = useDeletePostsPostIDLikes({
		mutation: { retry: 1 },
	});
	const [isLiked, setIsLiked] = useState(initialIsLiked);
	const [likes, setLikes] = useState(initialLikes);
	const handleLike = useCallback(() => {
		const fn = isLiked ? unlike : like;
		fn(
			{ postID: id },
			{
				onSuccess(likes) {
					setIsLiked((liked) => !liked);
					setLikes(likes.likes);
				},
			}
		);
	}, [id, isLiked, like, unlike]);
	return (
		<Group position="apart" align="center" mx="30%">
			<Group align="center" spacing="xs">
				<ActionIcon>
					<IconMessageCircle2 size={18} />
				</ActionIcon>
				<Text size="sm">{replies}</Text>
			</Group>
			<Group align="center" spacing="xs">
				<ActionIcon onClick={handleLike}>
					<IconHeart size={18} fill={isLiked ? "currentColor" : "none"} />
				</ActionIcon>
				<Text size="sm">{likes}</Text>
			</Group>
		</Group>
	);
}
