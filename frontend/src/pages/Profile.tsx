import {
	Avatar,
	Badge,
	Container,
	Group,
	Paper,
	Text,
	UnstyledButton,
} from "@mantine/core";
import { useContext } from "react";
import { Link, useParams } from "react-router-dom";
import { CenterLoader } from "../components/CenterLoader";
import { FollowButton } from "../components/FollowButton";
import { Posts, Query } from "../components/Posts";
import { AuthContext } from "../context/AuthContext";
import { UsersContext } from "../context/UsersContext";
import { QueryKey } from "../query-key";
import { getUsersHandlePosts, useGetUsersHandle } from "../spec.gen";
import { stopPropagation } from "../utils";

const Profile = () => {
	const params = useParams();
	const { users, setUser } = useContext(UsersContext);
	const { getAuth } = useContext(AuthContext);

	const query: Query = (params) => {
		return getUsersHandlePosts(user!.handle, params);
	};
	const { isLoading } = useGetUsersHandle(params.handle!, {
		query: {
			onSuccess: (u) => {
				setUser(u);
			},
		},
	});
	const user = users[params.handle!];

	return isLoading || !user ? (
		<CenterLoader />
	) : (
		<Paper withBorder radius="xs">
			<Container p="xs">
				<Group align="center" position="apart">
					<Avatar size="xl" radius={999} mb="xs" />
					{getAuth().handle === user.handle ? null : (
						<FollowButton handle={user.handle} />
					)}
				</Group>
				<Group>
					<Text weight="bold">@{user.handle}</Text>
					{user.status?.isFollowing ? <Badge>follows you</Badge> : null}
				</Group>
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
			<Posts query={query} queryKey={[QueryKey.POSTS, user.handle]} />
		</Paper>
	);
};

export default Profile;
