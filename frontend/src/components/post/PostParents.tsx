import { useGetPostsPostID } from "@/spec.gen";
import { Post } from "@/components/post";

type PostParentsProps = {
	parentID: number | undefined;
	scroll: () => void;
};

export function PostParents({ parentID, scroll }: PostParentsProps) {
	const enabled = !!parentID;
	const { data: post, isLoading } = useGetPostsPostID(parentID ?? 0, {
		query: {
			onSuccess: (data) => !data.data.repliesTo && scroll(),
			enabled,
		},
	});
	return enabled && !isLoading && post ? (
		<>
			<PostParents parentID={post.data.repliesTo} scroll={scroll} />
			<Post post={post.data} user={post.user} />
		</>
	) : null;
}
