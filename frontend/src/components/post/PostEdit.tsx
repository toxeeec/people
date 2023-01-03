import { ActionIcon, Popover, Stack, Text } from "@mantine/core";
import { IconDots } from "@tabler/icons";
import { QueryKey, useQueryClient } from "@tanstack/react-query";
import { useContext } from "react";
import { AuthContext } from "../../context/AuthContext";
import { deletePostsPostID } from "../../spec.gen";

interface PostEditProps {
	id: number;
	handle: string;
	queryKey: QueryKey;
	exact: boolean;
	onSuccess?: () => void;
}

export const PostEdit = ({
	id,
	handle,
	queryKey,
	exact,
	onSuccess,
}: PostEditProps) => {
	const { isAuthenticated, getAuth } = useContext(AuthContext);
	const queryClient = useQueryClient();
	const handleDelete = () => {
		deletePostsPostID(id).then(() => {
			queryClient.resetQueries({ queryKey, exact });
			if (onSuccess) {
				onSuccess();
			}
		});
	};
	return isAuthenticated && getAuth().handle === handle ? (
		<Popover position="bottom">
			<Popover.Target>
				<ActionIcon>
					<IconDots />
				</ActionIcon>
			</Popover.Target>
			<Popover.Dropdown p={0} style={{ cursor: "pointer" }}>
				<Stack spacing="xs">
					<Text onClick={handleDelete} p="xs">
						Delete Post
					</Text>
				</Stack>
			</Popover.Dropdown>
		</Popover>
	) : null;
};
