import { Modal, Text } from "@mantine/core";
import { MouseEvent } from "react";
import { Dispatch, SetStateAction } from "react";
import { getPostsPostIDLikes } from "../../spec.gen";
import Profiles, { Query } from "../Profiles";

interface PostLikesProps {
	opened: boolean;
	setOpened: Dispatch<SetStateAction<boolean>>;
	postID: number;
}

export default function PostLikes({
	opened,
	setOpened,
	postID,
}: PostLikesProps) {
	const query: Query = (params) => {
		return getPostsPostIDLikes(postID, params);
	};
	return (
		<Modal
			opened={opened}
			onClose={() => {
				setOpened(false);
			}}
			onClick={(e: MouseEvent) => e.stopPropagation()}
			centered
			title={<Text weight="bold">Liked by</Text>}
		>
			<Profiles query={query} queryKey={["likes", postID]} />
		</Modal>
	);
}
