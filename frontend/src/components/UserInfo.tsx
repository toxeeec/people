import { Avatar, Badge, Group, Text, UnstyledButton } from "@mantine/core";
import { useContext } from "react";
import { Link } from "react-router-dom";
import { UsersContext } from "../context/UsersContext";

interface UserInfoProps {
	handle: string;
	children?: React.ReactNode;
}

export const UserInfo = ({ handle, children }: UserInfoProps) => {
	const { users } = useContext(UsersContext);
	const user = users[handle]!;
	return (
		<>
			<Group align="start" position="apart">
				<Avatar
					size="lg"
					radius="xl"
					mb="xs"
					component={Link}
					to={`/${user!.handle}`}
				/>
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
