import { useNavigate, useParams } from "react-router";
import { getPostsPostIDReplies, postPostsPostIDReplies, useGetPostsPostID } from "@/spec.gen";
import { Container } from "@mantine/core";
import { useScrollIntoView } from "@mantine/hooks";
import { CenterLoader, Wrapper } from "@/components/utils";
import {
	CreatePost,
	type MutationFn,
	MainPost,
	PostParents,
	InfinitePosts,
	type PostsQuery,
} from "@/components/post";
import { HEADER_HEIGHT } from "@/layout/Header";
import { useContext, useEffect } from "react";
import { RouteContext } from "@/context/RouteContext";

export default function Post() {
	const navigate = useNavigate();
	const { setRouteName } = useContext(RouteContext);
	useEffect(() => {
		setRouteName("Post");
	}, [setRouteName]);
	const { scrollIntoView, targetRef } = useScrollIntoView<HTMLDivElement>({
		offset: HEADER_HEIGHT,
		duration: 0,
	});
	const scroll = () => scrollIntoView({ alignment: "start" });
	const params = useParams();
	const postID = +(params.postID ?? "");
	const { data: post, isLoading } = useGetPostsPostID(postID, {
		query: {
			onError: () => navigate("/404"),
		},
	});

	const mutationFn: MutationFn = (newPost) => postPostsPostIDReplies(postID, newPost);
	const query: PostsQuery = (queryParams) => getPostsPostIDReplies(postID, queryParams);

	return isLoading || !post ? (
		<CenterLoader />
	) : (
		<Wrapper>
			<PostParents parentID={post.data.repliesTo} scroll={scroll} />
			<MainPost post={post.data} user={post.user} ref={targetRef} />
			<Container p="md" pos="relative">
				<CreatePost mutationFn={mutationFn} placeholder={"Create reply"} />
			</Container>
			<InfinitePosts query={query} queryKey={["posts", postID, "replies"]} />
		</Wrapper>
	);
}
