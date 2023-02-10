import { useGetPostsPostID } from "../../spec.gen";
import { Post } from "../Post";
import { Post as PostType } from "../../models";

interface PostParentsProps {
	post: PostType;
	scroll: () => void;
}
export const PostParents = ({ post, scroll }: PostParentsProps) => {
	const enabled = !!post?.repliesTo;
	const replyID = post?.repliesTo || 0;
	const { data, isLoading } = useGetPostsPostID(replyID, {
		query: {
			onSuccess: (data) => !data.data.repliesTo && scroll(),
			enabled,
		},
	});
	return enabled && !isLoading && data ? (
		<>
			<PostParents post={data.data} scroll={scroll} />
			<Post post={data.data} user={data.user} />
		</>
	) : null;
};
