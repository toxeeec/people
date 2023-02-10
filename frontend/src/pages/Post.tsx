import { Posts, Query } from "../components/Posts";
import { MainPost } from "../components/MainPost";
import { useParams } from "react-router";
import {
	getPostsPostIDReplies,
	postPostsPostIDReplies,
	useGetPostsPostID,
} from "../spec.gen";
import { MutationFn, CreatePost } from "../components/post/CreatePost";
import { Container } from "@mantine/core";
import { PostParents } from "../components/post/PostParents";
import { useScrollIntoView } from "@mantine/hooks";
import { Wrapper } from "../components/Wrapper";
import { CenterLoader } from "../components/CenterLoader";

const Post = () => {
	const params = useParams();
	const postID = parseInt(params.postID!)!;
	const { data, isLoading } = useGetPostsPostID(postID);

	const mutationFn: MutationFn = (newPost) => {
		return postPostsPostIDReplies(postID, newPost);
	};

	const query: Query = (queryParams) => {
		return getPostsPostIDReplies(postID, queryParams);
	};

	const { scrollIntoView, targetRef } = useScrollIntoView<HTMLDivElement>({
		offset: 60,
		duration: 0,
	});
	const scroll = () => scrollIntoView({ alignment: "start" });
	return isLoading || !data ? (
		<CenterLoader />
	) : (
		<Wrapper>
			<PostParents post={data.data} scroll={scroll} />
			<MainPost post={data.data} user={data.user} ref={targetRef} />
			<Container p="md" pos="relative">
				<CreatePost mutationFn={mutationFn} placeholder={"Create reply"} />
			</Container>
			<Posts query={query} queryKey={["posts", postID, "replies"]} />
		</Wrapper>
	);
};

export default Post;
