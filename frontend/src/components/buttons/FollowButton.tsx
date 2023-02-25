import { Button } from "@mantine/core";
import { useQueryClient } from "@tanstack/react-query";
import { useCallback } from "react";
import { type User } from "@/models";
import { useDeleteMeFollowingHandle, usePutMeFollowingHandle } from "@/spec.gen";

type FollowButtonProps = {
	user: User;
};

export function FollowButton({ user }: FollowButtonProps) {
	const queryClient = useQueryClient();
	const { mutate: follow, isLoading: isFollowLoading } = usePutMeFollowingHandle({
		mutation: { retry: 1 },
	});
	const { mutate: unfollow, isLoading: isUnfollowLoading } = useDeleteMeFollowingHandle({
		mutation: { retry: 1 },
	});

	const isLoading = isFollowLoading || isUnfollowLoading;

	const handleFollow = useCallback(() => {
		(user.status?.isFollowed ? unfollow : follow)(
			{ handle: user.handle },
			{
				onSuccess: () => queryClient.invalidateQueries(),
			}
		);
	}, [follow, unfollow, user, queryClient]);

	return (
		<Button
			loading={isLoading}
			loaderPosition="center"
			onClick={handleFollow}
			variant={user?.status?.isFollowed ? "outline" : "filled"}
			radius="xl"
		>
			{user?.status?.isFollowed ? "Unfollow" : "Follow"}
		</Button>
	);
}
