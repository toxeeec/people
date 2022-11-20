import {
	Avatar,
	Button,
	Flex,
	Group,
	Modal,
	Text,
	Textarea,
} from "@mantine/core";
import {
	Dispatch,
	MouseEvent,
	SetStateAction,
	useCallback,
	useState,
} from "react";
import { ErrorType } from "../../custom-instance";
import { Error as CustomError, Post } from "../../models";
import {
	getPostsPostID,
	usePostPosts,
	usePostPostsPostIDReplies,
} from "../../spec.gen";

interface PostInputProps {
	isReply: boolean;
	opened: boolean;
	setOpened: Dispatch<SetStateAction<boolean>>;
	post: Post;
	setPost: Dispatch<SetStateAction<Post>>;
}

export default function PostInput({
	isReply,
	opened,
	setOpened,
	post,
	setPost,
}: PostInputProps) {
	const [content, setContent] = useState("");
	const [error, setError] = useState("");
	const { mutate: reply } = usePostPostsPostIDReplies({
		mutation: { retry: 1 },
	});
	const { mutate: create } = usePostPosts({ mutation: { retry: 1 } });

	const handleSubmit = useCallback(
		(e: MouseEvent) => {
			e.stopPropagation();
			const onError = (error: ErrorType<CustomError>) => {
				const err = error.response?.data.message;
				setError(err!);
			};

			if (isReply) {
				reply(
					{ postID: post.id, data: { content: content.trim() } },
					{
						onSuccess: () => {
							getPostsPostID(post.id).then((p) => setPost(p));
							setOpened(false);
						},
						onError,
					}
				);
			} else {
				create(
					{ data: { content: content.trim() } },
					{
						onSuccess: () => {
							setOpened(false);
						},
						onError,
					}
				);
			}
		},
		[isReply, create, reply, content, post, setPost, setOpened]
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
					disabled={content.trim().length === 0}
				>
					Submit
				</Button>
			</Flex>
		</Modal>
	);
}
