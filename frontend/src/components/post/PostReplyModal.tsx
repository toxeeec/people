import { Group, Modal, Text } from "@mantine/core";
import { useFocusTrap } from "@mantine/hooks";
import { Avatar } from "@/components/user";
import { type Post, type User } from "@/models";
import { postPostsPostIDReplies } from "@/spec.gen";
import { type MutationFn, CreatePost } from "@/components/post";

type PostReplyModalProps = {
	opened: boolean;
	handleClose: () => void;
	post: Post;
	user: User;
};

export function PostReplyModal({ opened, handleClose, post, user }: PostReplyModalProps) {
	const focusTrapRef = useFocusTrap(opened);

	const mutationFn: MutationFn = (newPost) => {
		return postPostsPostIDReplies(post.id, newPost);
	};

	return (
		<Modal
			opened={opened}
			onClose={() => handleClose()}
			title={
				<Group align="center">
					<Avatar user={user} size="md" />
					<Text weight="bold">{user.handle}</Text>
				</Group>
			}
			centered
			padding="lg"
		>
			<Text my="xs">{post.content}</Text>
			<CreatePost
				mutationFn={mutationFn}
				handleClose={handleClose}
				ref={focusTrapRef}
				placeholder={"Create reply"}
			/>
		</Modal>
	);
}
