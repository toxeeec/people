import { Button, Paper, Flex, Text, Textarea } from "@mantine/core";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { MouseEvent, useCallback, useContext, useState } from "react";
import { PostsContext } from "../../context/PostsContext";
import { UsersContext } from "../../context/UsersContext";
import { ErrorType } from "../../custom-instance";
import { NewPost, PostResponse } from "../../models";

export type MutationFn = (newPost: NewPost) => Promise<PostResponse>;
interface PostCreateProps {
	mutationFn: MutationFn;
	queryKey: readonly unknown[];
}

export const PostCreate = ({ mutationFn, queryKey }: PostCreateProps) => {
	const [content, setContent] = useState("");
	const [error, setError] = useState("");
	const queryClient = useQueryClient();
	const { setUser } = useContext(UsersContext);
	const { setPost } = useContext(PostsContext);
	const { mutate } = useMutation({
		mutationFn,
		retry: 1,
		onSuccess: (postResponse) => {
			setPost(postResponse.data);
			setUser(postResponse.user);
			setContent("");
			queryClient.invalidateQueries({ queryKey });
		},
		onError: (error) => {
			const err = (error as ErrorType<Error>).response?.data.message;
			setError(err!);
		},
	});

	const handleSubmit = useCallback(
		(e: MouseEvent) => {
			e.stopPropagation();
			mutate({ content: content.trim() });
		},
		[content, mutate]
	);

	const disabled = content.trim().length === 0 || content.trim().length > 280;
	return (
		<Paper p="lg" withBorder>
			<Textarea
				placeholder="Create new post"
				value={content}
				variant="unstyled"
				onChange={(e) => setContent(e.currentTarget.value)}
				error={error}
			/>
			<Flex justify="end" align="center" mt="md">
				<Text mr="md">{`${content.trim().length}/280`}</Text>
				<Button
					onClick={handleSubmit}
					variant="filled"
					radius="xl"
					disabled={disabled}
				>
					Post
				</Button>
			</Flex>
		</Paper>
	);
};
