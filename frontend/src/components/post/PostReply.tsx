import { Button, Flex, Paper, Text, Textarea } from "@mantine/core";
import { useFocusTrap } from "@mantine/hooks";
import { useQueryClient } from "@tanstack/react-query";
import {
	Dispatch,
	MouseEvent,
	SetStateAction,
	useCallback,
	useContext,
	useState,
} from "react";
import { PostsContext } from "../../context/PostsContext";
import { UsersContext } from "../../context/UsersContext";
import { QueryKey } from "../../query-key";
import { getPostsPostID, usePostPostsPostIDReplies } from "../../spec.gen";

interface PostReplyProps {
	id: number;
	setOpened?: Dispatch<SetStateAction<boolean>>;
}

export const PostReply = ({ id, setOpened }: PostReplyProps) => {
	const [content, setContent] = useState("");
	const [error, setError] = useState("");
	const focusTrapRef = useFocusTrap();
	const queryClient = useQueryClient();
	const { posts, setPost } = useContext(PostsContext);
	const { setUser } = useContext(UsersContext);
	const { mutate } = usePostPostsPostIDReplies({
		mutation: {
			retry: 1,
			onSuccess: (postResponse) => {
				setPost(postResponse.data);
				setUser(postResponse.user);
				setContent("");
				queryClient.invalidateQueries({
					queryKey: [QueryKey.REPLIES, id],
				});
				getPostsPostID(postResponse.data.repliesTo!).then((postResponse) => {
					setPost(postResponse.data);
					setUser(postResponse.user);
				});
			},
			onError: (error) => {
				const err = error.response?.data.message;
				setError(err!);
			},
		},
	});

	const handleSubmit = useCallback(
		(e: MouseEvent) => {
			e.stopPropagation();
			mutate({ postID: id, data: { content: content.trim() } });
			if (setOpened) {
				setOpened(false);
			}
		},
		[mutate, content, id, setOpened]
	);

	const post = posts[id]!;
	return (
		<Paper>
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
		</Paper>
	);
};
