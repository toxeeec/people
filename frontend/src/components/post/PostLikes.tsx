import { Modal, Text } from "@mantine/core";
import { MouseEvent } from "react";
import { Dispatch, SetStateAction } from "react";
import { QueryKey } from "../../query-key";
import { getPostsPostIDLikes } from "../../spec.gen";
import { Profiles, Query } from "../Profiles";

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
			onClick={(e: MouseEvent) => e.stopPropagation()}
			centered
			title={<Text weight="bold">Liked by</Text>}
		>
			<Profiles query={query} queryKey={[QueryKey.LIKES, id]} />
		</Modal>
	);
};
