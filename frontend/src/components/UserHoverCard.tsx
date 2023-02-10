import { HoverCard, Paper } from "@mantine/core";
import { useContext } from "react";
import { AuthContext } from "../context/AuthContext";
import { User } from "../models";
import { FollowButton } from "./FollowButton";
import { UserInfo } from "./UserInfo";

interface UserHoverCardProps {
	user: User;
	children: React.ReactNode;
}

export const UserHoverCard = ({ user, children }: UserHoverCardProps) => {
	const { getAuth } = useContext(AuthContext);
	return (
		<HoverCard>
			<HoverCard.Target>{children}</HoverCard.Target>
			<HoverCard.Dropdown p={0}>
				<Paper p="md">
					<UserInfo user={user}>
						{user?.handle === getAuth().handle ? null : (
							<FollowButton user={user} />
						)}
					</UserInfo>
				</Paper>
			</HoverCard.Dropdown>
		</HoverCard>
	);
};
