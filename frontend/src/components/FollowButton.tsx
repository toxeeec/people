import { Button } from "@mantine/core";
import { useQueryClient } from "@tanstack/react-query";
import { useCallback } from "react";
import { User } from "../models";
import {
	useDeleteMeFollowingHandle,
	usePutMeFollowingHandle,
} from "../spec.gen";

interface FollowButtonProps {
	user: User;
}

export const FollowButton = ({ user }: FollowButtonProps) => {
	const queryClient = useQueryClient();
	const { mutate: follow, isLoading: isFollowLoading } =
		usePutMeFollowingHandle({
			mutation: { retry: 1 },
		});
	const { mutate: unfollow, isLoading: isUnfollowLoading } =
		useDeleteMeFollowingHandle({
			mutation: { retry: 1 },
		});

	const isLoading = isFollowLoading || isUnfollowLoading;

	const handleFollow = useCallback(() => {
		const fn = user!.status?.isFollowed ? unfollow : follow;
		fn(
			{ handle: user!.handle! },
			{
				onSuccess: () => {
					queryClient.invalidateQueries();
				},
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
};
