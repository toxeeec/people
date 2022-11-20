import { useContext, useState } from "react";
import { Link, useLoaderData } from "react-router-dom";
import Posts, { Query } from "../components/Posts";
import UsersContext from "../context/UsersContext";
import { Post as PostType } from "../models";
import { useParams } from "react-router";
import { getPostsPostIDReplies } from "../spec.gen";
import { Avatar, Group, Paper, Text } from "@mantine/core";
import ProfileHoverCard from "../components/post/ProfileHoverCard";
import PostActions from "../components/post/PostActions";

export default function MainPost() {
	const params = useParams();
	const data = useLoaderData();
	const [post, setPost] = useState(data as PostType);
	const usersCtx = useContext(UsersContext)!;
	usersCtx.setUser(post.user!.handle, post.user!);

	const query: Query = (queryParams) => {
		return getPostsPostIDReplies(parseInt(params.postID!), queryParams);
	};

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
			<Posts query={query} queryKey={["replies", post.id.toString()]} />
		</>
	);
}
