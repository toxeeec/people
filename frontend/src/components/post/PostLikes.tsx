import { Modal, Text } from "@mantine/core";
import { Dispatch, SetStateAction } from "react";
import { getPostsPostIDLikes } from "../../spec.gen";
import { Users, Query } from "../Users";

interface PostLikesProps {
	opened: boolean;
	setOpened: Dispatch<SetStateAction<boolean>>;
	id: number;
}

export const PostLikes = ({ opened, setOpened, id }: PostLikesProps) => {
	const query: Query = (params) => {
		return getPostsPostIDLikes(id, params);
	};
	return (
		<Modal
			opened={opened}
			onClose={() => {
				setOpened(false);
			}}
			centered
			title={<Text weight="bold">Liked by</Text>}
		>
			<Users queryKey={["users", id, "likes"]} query={query} />
		</Modal>
	);
};
