import { Button, Paper, Flex, Text, Textarea } from "@mantine/core";
import { useQueryClient } from "@tanstack/react-query";
import { MouseEvent, useCallback, useState } from "react";
import { usePostPosts } from "../../spec.gen";

export default function PostCreate() {
	const [content, setContent] = useState("");
	const [error, setError] = useState("");
	const queryClient = useQueryClient();
	const { mutate } = usePostPosts({
		mutation: { retry: 1 },
	});

	const handleSubmit = useCallback(
		(e: MouseEvent) => {
			e.stopPropagation();
			mutate(
				{ data: { content: content.trim() } },
				{
					onSuccess: () => {
						setContent("");
						queryClient.invalidateQueries({ queryKey: ["feed"] });
					},
					onError: (error) => {
						const err = error.response?.data.message;
						setError(err!);
					},
				}
			);
		},
		[content, mutate, queryClient]
	);

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
					disabled={content.trim().length === 0 || content.trim().length > 280}
				>
					Post
				</Button>
			</Flex>
		</Paper>
	);
}
