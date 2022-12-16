import { Avatar, Group, Modal, Text } from "@mantine/core";
import { Dispatch, MouseEvent, SetStateAction } from "react";
import { PostReply } from "./PostReply";

interface PostReplyProps {
	isReply: boolean;
	opened: boolean;
	setOpened: Dispatch<SetStateAction<boolean>>;
	id: number;
	handle: string;
}

export const PostReplyModal = ({
	opened,
	setOpened,
	id,
	handle,
}: PostReplyProps) => {
	return (
		<Modal
			opened={opened}
			onClose={() => {
				setOpened(false);
			}}
			onClick={(e: MouseEvent) => e.stopPropagation()}
			centered
			title={
				<Group align="center">
					<Avatar radius="xl" size="md" />
					<Text weight="bold">{handle}</Text>
				</Group>
			}
		>
			<PostReply id={id} setOpened={setOpened} />
		</Modal>
	);
};
