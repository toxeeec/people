import { AuthContext } from "@/context/AuthContext";
import { useDeletePostsPostID } from "@/spec.gen";
import { ActionIcon, Popover, Stack, Text } from "@mantine/core";
import { IconDots } from "@tabler/icons";
import { useQueryClient } from "@tanstack/react-query";
import { useContext } from "react";

type PostEditProps = {
	postID: number;
	handle: string;
	onSuccess?: () => void;
};

export function PostEdit({ postID, handle, onSuccess }: PostEditProps) {
	const { getAuth } = useContext(AuthContext);
	const ownPost = handle === getAuth().handle;
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
	return ownPost ? (
		<Popover position="bottom">
			<Popover.Target>
				<ActionIcon>
					<IconDots />
				</ActionIcon>
			</Popover.Target>
			<Popover.Dropdown p={0} style={{ cursor: "pointer" }}>
				<Stack spacing="xs">
					<Text onClick={() => mutate({ postID })} p="xs">
						Delete Post
					</Text>
				</Stack>
			</Popover.Dropdown>
		</Popover>
	) : null;
}
