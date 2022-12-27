import { HoverCard, Paper } from "@mantine/core";
import { useContext } from "react";
import { AuthContext } from "../context/AuthContext";
import { UsersContext } from "../context/UsersContext";
import { FollowButton } from "./FollowButton";
import { UserInfo } from "./UserInfo";

interface UserHoverCardProps {
	handle: string;
	children: React.ReactNode;
}

export const UserHoverCard = ({ handle, children }: UserHoverCardProps) => {
	const { getAuth } = useContext(AuthContext);
	const { users } = useContext(UsersContext);
	const user = users[handle];
	return (
		<HoverCard>
			<HoverCard.Target>{children}</HoverCard.Target>
			<HoverCard.Dropdown p={0}>
				<Paper p="md">
					<UserInfo handle={handle}>
						{user?.handle === getAuth().handle ? null : (
							<FollowButton handle={handle} />
						)}
					</UserInfo>
				</Paper>
			</HoverCard.Dropdown>
		</HoverCard>
	);
};
