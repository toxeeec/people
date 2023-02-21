import { Badge, Box, Group, Tabs, Text, UnstyledButton } from "@mantine/core";
import { useContext } from "react";
import { Link, useNavigate, useParams } from "react-router-dom";
import { Avatar } from "../Avatar";
import { CenterLoader } from "../components/CenterLoader";
import { EditButton } from "../components/EditButton";
import { FollowButton } from "../components/FollowButton";
import { Posts, Query } from "../components/Posts";
import { Wrapper } from "../components/Wrapper";
import { AuthContext } from "../context/AuthContext";
import {
	getUsersHandleLikes,
	getUsersHandlePosts,
	useGetUsersHandle,
} from "../spec.gen";

export type UserPage = "posts" | "likes";

interface UserProps {
	value: UserPage;
}

const User = ({ value }: UserProps) => {
	const params = useParams();
	const { isAuthenticated, getAuth } = useContext(AuthContext);
	const navigate = useNavigate();

	const postsQuery: Query = (params) => {
		return getUsersHandlePosts(user!.handle, params);
	};
	const likesQuery: Query = (params) => {
		return getUsersHandleLikes(user!.handle, params);
	};
	const { data: user, isLoading } = useGetUsersHandle(params.handle!);

	return isLoading || !user ? (
		<CenterLoader />
	) : (
		<Wrapper>
			<Box p="xs">
				<Group align="center" position="apart">
					<Avatar user={user} size={120} mb="xs" />
					{isAuthenticated && getAuth().handle === user.handle ? (
						<EditButton handle={user.handle} />
					) : (
						<FollowButton user={user} />
					)}
				</Group>
				<Group>
					<Text weight="bold">{user.handle}</Text>
					{user.status?.isFollowing ? <Badge>follows you</Badge> : null}
				</Group>
				<Group mt="xs">
					<UnstyledButton component={Link} to={`/${user.handle}/following`}>
						<b>{user?.following}</b> Following
					</UnstyledButton>
					<UnstyledButton component={Link} to={`/${user.handle}/followers`}>
						<b>{user?.followers}</b>
						{user?.followers === 1 ? " Follower" : " Followers"}
					</UnstyledButton>
				</Group>
			</Box>
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
					<Posts query={postsQuery} queryKey={["posts", user.handle]} />
				</Tabs.Panel>
				<Tabs.Panel value="likes">
					<Posts
						query={likesQuery}
						queryKey={["posts", user.handle, "likes"]}
					/>
				</Tabs.Panel>
			</Tabs>
		</Wrapper>
	);
};

export default User;
