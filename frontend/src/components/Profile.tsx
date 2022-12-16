import { Paper } from "@mantine/core";
import { useContext } from "react";
import { AuthContext } from "../context/AuthContext";
import { UsersContext } from "../context/UsersContext";
import { AccountInfo } from "./AccountInfo";
import { FollowButton } from "./FollowButton";

interface ProfileProps {
	handle: string;
}

export const Profile = ({ handle }: ProfileProps) => {
	const { getAuth } = useContext(AuthContext);
	const { users } = useContext(UsersContext);
	const user = users[handle];

	return (
		<Paper p="md">
			<AccountInfo handle={handle}>
				{user?.handle === getAuth().handle ? null : (
					<FollowButton handle={handle} />
				)}
			</AccountInfo>
		</Paper>
	);
};

Profile.displayName = "Profile";
