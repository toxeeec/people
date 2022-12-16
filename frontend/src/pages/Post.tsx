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
import { MutationFn, PostCreate } from "../components/post/PostCreate";

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

	return isLoading || !data ? (
		<CenterLoader />
	) : (
		<>
			<MainPost id={postID} handle={data.user.handle} />
			<PostCreate
				mutationFn={mutationFn}
				queryKey={[QueryKey.REPLIES, postID]}
			/>
			<Posts query={query} queryKey={[QueryKey.REPLIES, postID]} />
		</>
	);
};

export default Post;
