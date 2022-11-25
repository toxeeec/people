import { Avatar, Badge, Group, Text, UnstyledButton } from "@mantine/core";
import { Link } from "react-router-dom";
import { User } from "../models";
import { stopPropagation } from "../utils";

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
					onClick={stopPropagation}
				/>
				{children}
			</Group>
			<Text
				component={Link}
				to={`/${user!.handle}`}
				weight="bold"
				onClick={stopPropagation}
			>
				@{user?.handle}
			</Text>
			{user.isFollowing ? <Badge ml="xs">follows you</Badge> : null}
			<Group mt="xs">
				<UnstyledButton
					component={Link}
					to={`/${user.handle}/following`}
					onClick={stopPropagation}
				>
					<b>{user?.following}</b> Following
				</UnstyledButton>
				<UnstyledButton
					component={Link}
					to={`/${user.handle}/followers`}
					onClick={stopPropagation}
				>
					<b>{user?.followers}</b>
					{user?.followers === 1 ? " Follower" : " Followers"}
				</UnstyledButton>
			</Group>
		</>
	);
}
