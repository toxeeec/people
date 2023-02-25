import { Modal, Text } from "@mantine/core";
import { type Dispatch, type SetStateAction } from "react";
import { getPostsPostIDLikes } from "@/spec.gen";
import { InfiniteUsers, type UsersQuery } from "@/components/user";

type PostLikesProps = {
	opened: boolean;
	setOpened: Dispatch<SetStateAction<boolean>>;
	id: number;
};

export function PostLikes({ opened, setOpened, id }: PostLikesProps) {
	const query: UsersQuery = (params) => {
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
			<InfiniteUsers queryKey={["users", id, "likes"]} query={query} />
		</Modal>
	);
}
