import { useContext } from "react";
import { PostsContext } from "../../context/PostsContext";
import { UsersContext } from "../../context/UsersContext";
import { QueryKey } from "../../query-key";
import { useGetPostsPostID } from "../../spec.gen";
import { Post } from "../Post";

interface PostParentsProps {
	id: number;
	scroll: () => void;
}
export const PostParents = ({ id, scroll }: PostParentsProps) => {
	const { posts, setPost } = useContext(PostsContext);
	const { setUser } = useContext(UsersContext);
	const post = posts[id];
	const enabled = !!post?.repliesTo;
	const replyID = post?.repliesTo || 0;
	const { data, isLoading } = useGetPostsPostID(replyID, {
		query: {
			onSuccess: (data) => {
				if (!data.data.repliesTo) {
					scroll();
				}
				setPost(data.data);
				setUser(data.user);
			},
			enabled,
		},
	});
	return enabled && !isLoading && data ? (
		<>
			<PostParents id={replyID} scroll={scroll} />
			<Post
				id={data.data.id}
				handle={data.user.handle}
				queryKey={[QueryKey.PARENT, data.data.id]}
			/>
		</>
	) : null;
};
