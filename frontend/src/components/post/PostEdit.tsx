import { ActionIcon, Popover, Stack, Text } from "@mantine/core";
import { IconDots } from "@tabler/icons";
import { useQueryClient } from "@tanstack/react-query";
import { useContext } from "react";
import { AuthContext } from "../../context/AuthContext";
import { useDeletePostsPostID } from "../../spec.gen";

interface PostEditProps {
	id: number;
	handle: string;
	exact: boolean;
	onSuccess?: () => void;
}

export const PostEdit = ({ id, handle, onSuccess }: PostEditProps) => {
	const { isAuthenticated, getAuth } = useContext(AuthContext);
	const queryClient = useQueryClient();
	const { mutate } = useDeletePostsPostID({
		mutation: {
			onSuccess: () => {
				queryClient.invalidateQueries();
				onSuccess && onSuccess();
			},
			retry: 1,
		},
	});
	return isAuthenticated && getAuth().handle === handle ? (
		<Popover position="bottom">
			<Popover.Target>
				<ActionIcon>
					<IconDots />
				</ActionIcon>
			</Popover.Target>
			<Popover.Dropdown p={0} style={{ cursor: "pointer" }}>
				<Stack spacing="xs">
					<Text onClick={() => mutate({ postID: id })} p="xs">
						Delete Post
					</Text>
				</Stack>
			</Popover.Dropdown>
		</Popover>
	) : null;
};
