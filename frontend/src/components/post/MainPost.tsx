import {
	Avatar,
	Text,
	Group,
	Paper,
	ActionIcon,
	Flex,
	Divider,
	UnstyledButton,
} from "@mantine/core";
import { useCallback, useState } from "react";
import { Link } from "react-router-dom";
import { Post as PostData } from "../../models";
import { stopPropagation } from "../../utils";
import ProfileHoverCard from "../ProfileHoverCard";
import {
	useDeletePostsPostIDLikes,
	usePutPostsPostIDLikes,
} from "../../spec.gen";
import { IconHeart, IconMessageCircle2 } from "@tabler/icons";
import PostReply from "./PostReply";
import PostLikes from "./PostLikes";

interface MainPostProps {
	post: PostData;
}

export default function MainPost({ post: initialPost }: MainPostProps) {
	const [post, setPost] = useState(initialPost);
	const [replyOpened, setReplyOpened] = useState(false);
	const [likesOpened, setLikesOpened] = useState(false);
	const { mutate: like } = usePutPostsPostIDLikes({ mutation: { retry: 1 } });
	const { mutate: unlike } = useDeletePostsPostIDLikes({
		mutation: { retry: 1 },
	});
	const handleLike = useCallback(() => {
		const fn = post.isLiked ? unlike : like;
		fn(
			{ postID: post.id },
			{
				onSuccess({ likes, isLiked }) {
					setPost((p) => ({ ...p, likes, isLiked }));
				},
			}
		);
	}, [post, setPost, like, unlike]);

	const handleOpen = () => {
		setReplyOpened(true);
	};

	return (
		<Paper p="xs" radius="xs" withBorder>
			<Group align="center">
				<ProfileHoverCard handle={post.user!.handle}>
					<Avatar
						radius="xl"
						size="md"
						component={Link}
						to={`/${post.user!.handle}`}
						onClick={stopPropagation}
					/>
				</ProfileHoverCard>
				<ProfileHoverCard handle={post.user!.handle}>
					<Text
						component={Link}
						to={`/${post.user!.handle}`}
						weight="bold"
						onClick={stopPropagation}
					>
						{post.user?.handle}
					</Text>
				</ProfileHoverCard>
			</Group>
			<Text my="xs">{post.content}</Text>
			<Divider />
			<UnstyledButton my="md" onClick={() => setLikesOpened(true)}>
				<b>{post.likes}</b>
				{post.likes === 1 ? " Like" : " Likes"}
			</UnstyledButton>
			<Divider mb="xs" />
			<Flex justify="space-around">
				<Group align="center" spacing="xs">
					<ActionIcon onClick={handleOpen}>
						<IconMessageCircle2 size={18} />
					</ActionIcon>
				</Group>
				<Group align="center" spacing="xs">
					<ActionIcon onClick={handleLike}>
						<IconHeart
							size={18}
							fill={post.isLiked ? "currentColor" : "none"}
						/>
					</ActionIcon>
				</Group>
			</Flex>
			<PostReply
				opened={replyOpened}
				setOpened={setReplyOpened}
				isReply={true}
				post={post}
				setPost={setPost}
			/>
			<PostLikes
				opened={likesOpened}
				setOpened={setLikesOpened}
				postID={post.id}
			/>
		</Paper>
	);
}
