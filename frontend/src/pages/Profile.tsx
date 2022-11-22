import {
	Avatar,
	Container,
	Group,
	Paper,
	Text,
	UnstyledButton,
} from "@mantine/core";
import { useContext, useState } from "react";
import { Link, useLoaderData } from "react-router-dom";
import FollowButton from "../components/FollowButton";
import Posts, { Query } from "../components/Posts";
import UsersContext from "../context/UsersContext";
import { User } from "../models";
import { getUsersHandlePosts } from "../spec.gen";
import { stopPropagation } from "../utils";

export default function Profile() {
	const data = useLoaderData();
	const handle = (data as User).handle!;
	const [user, setUser] = useState<Partial<User>>(data as User);
	const usersCtx = useContext(UsersContext)!;

	const updateUser = (u: Partial<User>) => {
		setUser(usersCtx.setUser(handle, u));
	};

	const query: Query = (params) => {
		return getUsersHandlePosts(handle, params);
	};

	return (
		<Paper withBorder radius="xs">
			<Container p="xs">
				<Group align="center" position="apart">
					<Avatar size="xl" radius={999} mb="xs" />
					<FollowButton user={user} updateUser={updateUser} />
				</Group>
				<Text weight="bold">@{handle}</Text>
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
			</Container>
			<Posts query={query} user={user as User} queryKey={["posts", handle]} />
		</Paper>
	);
}
