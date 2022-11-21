import {
	Avatar,
	Button,
	Flex,
	Group,
	Modal,
	Text,
	Textarea,
} from "@mantine/core";
import { useFocusTrap } from "@mantine/hooks";
import { useQueryClient } from "@tanstack/react-query";
import {
	Dispatch,
	MouseEvent,
	SetStateAction,
	useCallback,
	useState,
} from "react";
import { Post } from "../../models";
import { getPostsPostID, usePostPostsPostIDReplies } from "../../spec.gen";

interface PostReplyProps {
	isReply: boolean;
	opened: boolean;
	setOpened: Dispatch<SetStateAction<boolean>>;
	post: Post;
	setPost: Dispatch<SetStateAction<Post>>;
}

export default function PostReply({
	opened,
	setOpened,
	post,
	setPost,
}: PostReplyProps) {
	const [content, setContent] = useState("");
	const [error, setError] = useState("");
	const focusTrapRef = useFocusTrap();
	const queryClient = useQueryClient();
	const { mutate } = usePostPostsPostIDReplies({
		mutation: { retry: 1 },
	});

	const handleSubmit = useCallback(
		(e: MouseEvent) => {
			e.stopPropagation();
			mutate(
				{ postID: post.id, data: { content: content.trim() } },
				{
					onSuccess: () => {
						setContent("");
						getPostsPostID(post.id).then((p) => setPost(p));
						queryClient.invalidateQueries({
							queryKey: ["replies", post.id.toString()],
						});
						setOpened(false);
					},
					onError: (error) => {
						const err = error.response?.data.message;
						setError(err!);
					},
				}
			);
		},
		[content, post, setPost, setOpened, mutate, queryClient]
	);

	return (
		<Modal
			opened={opened}
			onClose={() => {
				setContent("");
				setOpened(false);
			}}
			onClick={(e: MouseEvent) => e.stopPropagation()}
			centered
			title={
				<Group align="center">
					<Avatar radius="xl" size="md" />
					<Text weight="bold">{post.user?.handle}</Text>
				</Group>
			}
		>
			<Text my="xs">{post.content}</Text>
			<Textarea
				ref={focusTrapRef}
				data-autofocus
				placeholder="Create reply"
				variant="unstyled"
				value={content}
				onChange={(e) => setContent(e.currentTarget.value)}
				error={error}
			/>
			<Flex justify="end" align="center" mt="md">
				<Text mr="md">{`${content.trim().length}/280`}</Text>
				<Button
					onClick={handleSubmit}
					variant="filled"
					radius="xl"
					disabled={content.trim().length === 0 || content.trim().length > 280}
				>
					Reply
				</Button>
			</Flex>
		</Modal>
	);
}
