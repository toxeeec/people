import {
	useCallback,
	useContext,
	useState,
	MouseEvent,
	useEffect,
} from "react";
import Posts, { Query } from "../components/Posts";
import UsersContext from "../context/UsersContext";
import MainPost from "../components/post/MainPost";
import { useParams } from "react-router";
import {
	getPostsPostIDReplies,
	usePostPostsPostIDReplies,
	useGetPostsPostID,
} from "../spec.gen";
import { Button, Flex, Paper, Text, Textarea } from "@mantine/core";
import { useQueryClient } from "@tanstack/react-query";
import CenterLoader from "../components/CenterLoader";

export default function Post() {
	const params = useParams();
	const { data, isLoading, refetch, isRefetching } = useGetPostsPostID(
		parseInt(params.postID!)
	);
	const [content, setContent] = useState("");
	const [error, setError] = useState("");
	const usersCtx = useContext(UsersContext)!;
	const queryClient = useQueryClient();
	const { mutate } = usePostPostsPostIDReplies({
		mutation: { retry: 1 },
	});

	useEffect(() => {
		if (!isLoading && data) {
			usersCtx.setUser(data.user!.handle, data.user!);
		}
	}, [isLoading, data, usersCtx]);

	const query: Query = (queryParams) => {
		return getPostsPostIDReplies(parseInt(params.postID!), queryParams);
	};

	const handleSubmit = useCallback(
		(e: MouseEvent) => {
			e.stopPropagation();
			if (data) {
				mutate(
					{ postID: data.id, data: { content: content.trim() } },
					{
						onSuccess: () => {
							setContent("");
							refetch();

							queryClient.invalidateQueries({
								queryKey: ["replies", data.id],
							});
						},
						onError: (error) => {
							const err = error.response?.data.message;
							setError(err!);
						},
					}
				);
			}
		},
		[content, mutate, queryClient, data, refetch]
	);

	return isLoading || isRefetching ? (
		<CenterLoader />
	) : (
		<>
			<MainPost post={data!} />
			<Paper withBorder p="xs">
				<Textarea
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
						disabled={
							content.trim().length === 0 || content.trim().length > 280
						}
					>
						Reply
					</Button>
				</Flex>
			</Paper>
			<Posts query={query} queryKey={["replies", parseInt(params.postID!)]} />
		</>
	);
}
