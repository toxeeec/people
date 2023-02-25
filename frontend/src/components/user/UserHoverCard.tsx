import { FollowButton } from "@/components/buttons";
import { UserInfo } from "@/components/user/UserInfo";
import { AuthContext } from "@/context/AuthContext";
import { type User } from "@/models";
import { HoverCard, Paper } from "@mantine/core";
import { useContext } from "react";

type UserHoverCardProps = {
	user: User;
	children: React.ReactNode;
};

export function UserHoverCard({ user, children }: UserHoverCardProps) {
	const { getAuth } = useContext(AuthContext);
	const ownProfile = user.handle === getAuth().handle;
	return (
		<HoverCard>
			<HoverCard.Target>{children}</HoverCard.Target>
			<HoverCard.Dropdown p={0}>
				<Paper p="md">
					<UserInfo user={user}>{!ownProfile && <FollowButton user={user} />}</UserInfo>
				</Paper>
			</HoverCard.Dropdown>
		</HoverCard>
	);
}
