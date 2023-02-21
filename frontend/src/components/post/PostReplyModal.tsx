import { Group, Modal, Text } from "@mantine/core";
import { useFocusTrap } from "@mantine/hooks";
import { Dispatch, SetStateAction } from "react";
import { Avatar } from "../../Avatar";
import { Post, User } from "../../models";
import { postPostsPostIDReplies } from "../../spec.gen";
import { MutationFn, CreatePost } from "./CreatePost";

interface PostReplyModalProps {
	opened: boolean;
	setOpened: Dispatch<SetStateAction<boolean>>;
	post: Post;
	user: User;
}

export const PostReplyModal = ({
	opened,
	setOpened,
	post,
	user,
}: PostReplyModalProps) => {
	const mutationFn: MutationFn = (newPost) => {
		return postPostsPostIDReplies(post.id, newPost);
	};
	const focusTrapRef = useFocusTrap(opened);
	return (
		<Modal
			opened={opened}
			onClose={() => {
				setOpened(false);
			}}
			title={
				<Group align="center">
					<Avatar user={user} size="md" />
					<Text weight="bold">{user.handle}</Text>
				</Group>
			}
			centered
			padding="lg"
		>
			<Text my="xs">{post?.content}</Text>
			<CreatePost
				mutationFn={mutationFn}
				setOpened={setOpened}
				ref={focusTrapRef}
				placeholder={"Create reply"}
			/>
		</Modal>
	);
};
