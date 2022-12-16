import { Button } from "@mantine/core";
import { MouseEvent, useCallback, useContext } from "react";
import { UsersContext } from "../context/UsersContext";
import {
	useDeleteMeFollowingHandle,
	usePutMeFollowingHandle,
} from "../spec.gen";

interface FollowButtonProps {
	handle: string;
}

export const FollowButton = ({ handle }: FollowButtonProps) => {
	const { users, setUser } = useContext(UsersContext);
	const user = users[handle];
	const { mutate: follow, isLoading: isFollowLoading } =
		usePutMeFollowingHandle({
			mutation: { retry: 1 },
		});
	const { mutate: unfollow, isLoading: isUnfollowLoading } =
		useDeleteMeFollowingHandle({
			mutation: { retry: 1 },
		});

	const isLoading = isFollowLoading || isUnfollowLoading;

	const handleFollow = useCallback(
		(e: MouseEvent) => {
			e.stopPropagation();
			const fn = user!.status?.isFollowed ? unfollow : follow;
			fn(
				{ handle: user!.handle! },
				{
					onSuccess: (user) => {
						setUser(user);
					},
				}
			);
		},
		[follow, unfollow, user, setUser]
	);

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
