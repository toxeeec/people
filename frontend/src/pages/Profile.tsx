import {
	Avatar,
	Badge,
	Container,
	Group,
	Paper,
	Tabs,
	Text,
	UnstyledButton,
} from "@mantine/core";
import { useContext } from "react";
import { Link, useNavigate, useParams } from "react-router-dom";
import { CenterLoader } from "../components/CenterLoader";
import { FollowButton } from "../components/FollowButton";
import { Posts, Query } from "../components/Posts";
import { AuthContext } from "../context/AuthContext";
import { UsersContext } from "../context/UsersContext";
import { QueryKey } from "../query-key";
import {
	getUsersHandleLikes,
	getUsersHandlePosts,
	useGetUsersHandle,
} from "../spec.gen";
import { stopPropagation } from "../utils";

export type ProfilePage = "posts" | "likes";

interface ProfileProps {
	value: ProfilePage;
}

const Profile = ({ value }: ProfileProps) => {
	const params = useParams();
	const { users, setUser } = useContext(UsersContext);
	const { isAuthenticated, getAuth } = useContext(AuthContext);
	const navigate = useNavigate();

	const postsQuery: Query = (params) => {
		return getUsersHandlePosts(user!.handle, params);
	};
	const likesQuery: Query = (params) => {
		return getUsersHandleLikes(user!.handle, params);
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
					{!isAuthenticated || getAuth().handle === user.handle ? null : (
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
			<Tabs
				value={value}
				onTabChange={(value) => {
					const url =
						value === "likes"
							? `/${params.handle}/${value}`
							: `/${params.handle}`;
					navigate(url, { replace: true });
				}}
			>
				<Tabs.List grow position="center">
					<Tabs.Tab value="posts">Posts</Tabs.Tab>
					<Tabs.Tab value="likes">Likes</Tabs.Tab>
				</Tabs.List>

				<Tabs.Panel value="posts">
					<Posts query={postsQuery} queryKey={[QueryKey.POSTS, user.handle]} />
				</Tabs.Panel>
				<Tabs.Panel value="likes">
					<Posts query={likesQuery} queryKey={[QueryKey.LIKES, user.handle]} />
				</Tabs.Panel>
			</Tabs>
		</Paper>
	);
};

export default Profile;
