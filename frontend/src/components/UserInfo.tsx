import { Badge, Group, Text, UnstyledButton } from "@mantine/core";
import { Link } from "react-router-dom";
import { Avatar } from "../Avatar";
import { User } from "../models";

interface UserInfoProps {
	user: User;
	children?: React.ReactNode;
}

export const UserInfo = ({ user, children }: UserInfoProps) => {
	return (
		<>
			<Group align="start" position="apart">
				<Link to={`/${user!.handle}`}>
					<Avatar size="lg" mb="xs" user={user} />
				</Link>
				{children}
			</Group>
			<Text component={Link} to={`/${user!.handle}`} weight="bold">
				{user.handle}
			</Text>
			{user.status?.isFollowing ? <Badge ml="xs">follows you</Badge> : null}
			<Group mt="xs">
				<UnstyledButton component={Link} to={`/${user.handle}/following`}>
					<b>{user.following}</b> Following
				</UnstyledButton>
				<UnstyledButton component={Link} to={`/${user.handle}/followers`}>
					<b>{user.followers}</b>
					{user.followers === 1 ? " Follower" : " Followers"}
				</UnstyledButton>
			</Group>
		</>
	);
};
