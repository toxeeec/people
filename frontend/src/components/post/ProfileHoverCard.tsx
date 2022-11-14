import { Button, HoverCard } from "@mantine/core";
import { useCallback, useContext, useState } from "react";
import UsersContext from "../../context/UsersContext";
import {
	useDeleteMeFollowingHandle,
	usePutMeFollowingHandle,
} from "../../spec.gen";
import AccountInfo from "../AccountInfo";

interface ProfileHoverCardProps {
	handle: string;
	children: React.ReactNode;
}

export default function ProfileHoverCard({
	children,
	handle,
}: ProfileHoverCardProps) {
	const { mutate: follow } = usePutMeFollowingHandle();
	const { mutate: unfollow } = useDeleteMeFollowingHandle();
	const [isFollowed, setIsFollowed] = useState(false);
	const usersCtx = useContext(UsersContext);

	const handleOpen = () => {
		const user = usersCtx?.users[handle];
		setIsFollowed(user?.isFollowed || false);
	};

	const handleFollow = useCallback(() => {
		const fn = isFollowed ? unfollow : follow;
		fn(
			{ handle },
			{
				onSuccess(follows) {
					usersCtx?.setUser(handle, follows);
					setIsFollowed(follows.isFollowed);
				},
			}
		);
	}, [follow, unfollow, handle, isFollowed, usersCtx]);

	return (
		<HoverCard onOpen={handleOpen}>
			<HoverCard.Target>{children}</HoverCard.Target>
			<HoverCard.Dropdown>
				<AccountInfo handle={handle}>
					<Button
						onClick={handleFollow}
						variant={isFollowed ? "outline" : "filled"}
						radius="xl"
					>
						{isFollowed ? "Unfollow" : "Follow"}
					</Button>
				</AccountInfo>
			</HoverCard.Dropdown>
		</HoverCard>
	);
}
