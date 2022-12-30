import { useContext } from "react";
import { Posts, Query } from "../components/Posts";
import { UsersContext } from "../context/UsersContext";
import { MainPost } from "../components/MainPost";
import { useParams } from "react-router";
import {
	getPostsPostIDReplies,
	useGetPostsPostID,
	postPostsPostIDReplies,
} from "../spec.gen";
import { QueryKey } from "../query-key";
import { PostsContext } from "../context/PostsContext";
import { CenterLoader } from "../components/CenterLoader";
import { MutationFn, CreatePost } from "../components/post/CreatePost";
import { Container } from "@mantine/core";
import { PostParents } from "../components/post/PostParents";
import { useScrollIntoView } from "@mantine/hooks";
import { Wrapper } from "../components/Wrapper";

const Post = () => {
	const params = useParams();
	const postID = parseInt(params.postID!)!;
	const { setUser } = useContext(UsersContext);
	const { setPost } = useContext(PostsContext);

	const { data, isLoading } = useGetPostsPostID(postID, {
		query: {
			onSuccess: (postResponse) => {
				setPost(postResponse.data);
				setUser(postResponse.user);
			},
		},
	});

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
			<PostParents id={postID} scroll={scroll} />
			<MainPost id={postID} handle={data.user.handle} ref={targetRef} />
			<Container p="md" pos="relative">
				<CreatePost
					mutationFn={mutationFn}
					queryKey={[QueryKey.REPLIES, postID]}
					placeholder={"Create reply"}
				/>
			</Container>
			<Posts query={query} queryKey={[QueryKey.REPLIES, postID]} />
		</Wrapper>
	);
};

export default Post;
