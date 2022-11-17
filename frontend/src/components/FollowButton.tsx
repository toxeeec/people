import { Button, Loader } from "@mantine/core";
import { useCallback, useEffect, useState } from "react";
import { User } from "../models";
import {
	useDeleteMeFollowingHandle,
	usePutMeFollowingHandle,
} from "../spec.gen";

interface FollowButtonProps {
	user: Partial<User>;
	updateUser: (_user: Partial<User>) => void;
}

export default function FollowButton({ user, updateUser }: FollowButtonProps) {
	const { mutate: follow, isLoading: isFollowLoading } =
		usePutMeFollowingHandle({
			mutation: { retry: 1 },
		});
	const { mutate: unfollow, isLoading: isUnfollowLoading } =
		useDeleteMeFollowingHandle({
			mutation: { retry: 1 },
		});

	const [isLoading, setIsLoading] = useState(false);
	useEffect(() => {
		setIsLoading(isFollowLoading || isUnfollowLoading);
	}, [isFollowLoading, isUnfollowLoading]);

	const handleFollow = useCallback(() => {
		const fn = user?.isFollowed ? unfollow : follow;
		fn(
			{ handle: user.handle! },
			{
				onSuccess(follows) {
					updateUser(follows);
				},
			}
		);
	}, [follow, unfollow, user, updateUser]);

	return isLoading ? (
		<Loader />
	) : (
		<Button
			onClick={handleFollow}
			variant={user.isFollowed ? "outline" : "filled"}
			radius="xl"
		>
			{user.isFollowed ? "Unfollow" : "Follow"}
		</Button>
	);
}