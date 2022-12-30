import { Avatar, Group, Modal, Text } from "@mantine/core";
import { useFocusTrap } from "@mantine/hooks";
import { Dispatch, SetStateAction, useContext } from "react";
import { PostsContext } from "../../context/PostsContext";
import { postPostsPostIDReplies } from "../../spec.gen";
import { stopPropagation } from "../../utils";
import { MutationFn, CreatePost } from "./CreatePost";

interface PostReplyModalProps {
	opened: boolean;
	setOpened: Dispatch<SetStateAction<boolean>>;
	id: number;
	handle: string;
	queryKey: readonly unknown[];
}

export const PostReplyModal = ({
	opened,
	setOpened,
	id,
	handle,
	queryKey,
}: PostReplyModalProps) => {
	const { posts } = useContext(PostsContext);
	const post = posts[id];
	const mutationFn: MutationFn = (newPost) => {
		return postPostsPostIDReplies(id, newPost);
	};
	const focusTrapRef = useFocusTrap(opened);
	return (
		<Modal
			opened={opened}
			onClose={() => {
				setOpened(false);
			}}
			onClick={stopPropagation}
			title={
				<Group align="center">
					<Avatar radius="xl" size="md" />
					<Text weight="bold">{handle}</Text>
				</Group>
			}
			centered
			padding="lg"
		>
			<Text my="xs">{post?.content}</Text>
			<CreatePost
				mutationFn={mutationFn}
				queryKey={queryKey}
				setOpened={setOpened}
				ref={focusTrapRef}
				placeholder={"Create reply"}
			/>
		</Modal>
	);
};
