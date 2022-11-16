import { Avatar, Group, Text } from "@mantine/core";
import { Link } from "react-router-dom";
import { User } from "../models";

interface AccountInfoProps {
	user: Partial<User>;
	children?: React.ReactNode;
}

export default function AccountInfo({ user, children }: AccountInfoProps) {
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
				@{user?.handle}
			</Text>
			<Group mt="xs">
				<span>
					<Text component={Link} to={`/${user!.handle}`} weight="bold">
						{user?.following}
					</Text>
					{" Following"}
				</span>
				<span>
					<b>{user?.followers}</b>
					{user?.followers === 1 ? " Follower" : " Followers"}
				</span>
			</Group>
		</>
	);
}
