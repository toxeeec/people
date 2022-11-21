import { useCallback, useContext, useState, MouseEvent } from "react";
import { Link, useLoaderData } from "react-router-dom";
import Posts, { Query } from "../components/Posts";
import UsersContext from "../context/UsersContext";
import { Post as PostType } from "../models";
import { useParams } from "react-router";
import {
	getPostsPostIDReplies,
	usePostPostsPostIDReplies,
	getPostsPostID,
} from "../spec.gen";
import {
	Avatar,
	Button,
	Flex,
	Group,
	Paper,
	Text,
	Textarea,
} from "@mantine/core";
import ProfileHoverCard from "../components/post/ProfileHoverCard";
import PostActions from "../components/post/PostActions";
import { useQueryClient } from "@tanstack/react-query";

export default function MainPost() {
	const params = useParams();
	const data = useLoaderData();
	const [post, setPost] = useState(data as PostType);
	const [content, setContent] = useState("");
	const [error, setError] = useState("");
	const usersCtx = useContext(UsersContext)!;
	const queryClient = useQueryClient();

	usersCtx.setUser(post.user!.handle, post.user!);
	const { mutate } = usePostPostsPostIDReplies({
		mutation: { retry: 1 },
	});

	const query: Query = (queryParams) => {
		return getPostsPostIDReplies(parseInt(params.postID!), queryParams);
	};

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
					},
					onError: (error) => {
						const err = error.response?.data.message;
						setError(err!);
					},
				}
			);
		},
		[content, post, setPost, mutate, queryClient]
	);

	return (
		<>
			<Paper p="xs" radius="xs" withBorder>
				<Group align="center">
					<ProfileHoverCard handle={post.user!.handle}>
						<Avatar
							radius="xl"
							size="md"
							component={Link}
							to={`/${post.user!.handle}`}
						/>
					</ProfileHoverCard>
					<ProfileHoverCard handle={post.user!.handle}>
						<Text component={Link} to={`/${post.user!.handle}`} weight="bold">
							{post.user?.handle}
						</Text>
					</ProfileHoverCard>
				</Group>
				<Text my="xs">{post.content}</Text>
				<PostActions post={post} setPost={setPost} />
			</Paper>
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
			<Posts query={query} queryKey={["replies", post.id.toString()]} />
		</>
	);
}
