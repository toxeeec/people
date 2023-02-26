import { EditButton, FollowButton } from "@/components/buttons";
import { InfinitePosts, type PostsQuery } from "@/components/post";
import { Avatar } from "@/components/user";
import { CenterLoader, Wrapper } from "@/components/utils";
import { AuthContext } from "@/context/AuthContext";
import { RouteContext } from "@/context/RouteContext";
import { getUsersHandleLikes, getUsersHandlePosts, useGetUsersHandle } from "@/spec.gen";
import { Badge, Box, Group, Tabs, Text, UnstyledButton } from "@mantine/core";
import { useContext, useEffect } from "react";
import { Link, useNavigate, useParams } from "react-router-dom";

export type UserPage = "posts" | "likes";

type UserProps = {
	value: UserPage;
};

export default function User({ value }: UserProps) {
	const params = useParams();
	const handle = params.handle ?? "";
	const { setRouteName } = useContext(RouteContext);
	useEffect(() => {
		setRouteName(handle);
	}, [setRouteName, handle]);
	const { isAuthenticated, getAuth } = useContext(AuthContext);
	const navigate = useNavigate();
	const { data: user, isLoading } = useGetUsersHandle(handle, {
		query: {
			onError: () => navigate("/404"),
		},
	});

	const ownProfile = handle === getAuth().handle;
	const postsQuery: PostsQuery = (params) => getUsersHandlePosts(handle, params);
	const likesQuery: PostsQuery = (params) => getUsersHandleLikes(handle, params);

	return isLoading || !user ? (
		<CenterLoader />
	) : (
		<Wrapper>
			<Box p="xs">
				<Group align="center" position="apart">
					<Avatar user={user} size={120} mb="xs" />
					{isAuthenticated &&
						(ownProfile ? <EditButton user={user} /> : <FollowButton user={user} />)}
				</Group>
				<Group>
					<Text weight="bold">{user.handle}</Text>
					{user.status?.isFollowing && <Badge>follows you</Badge>}
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
				onTabChange={(value) => navigate(`/${handle}/${value}`, { replace: true })}
			>
				<Tabs.List grow position="center">
					<Tabs.Tab value="posts">Posts</Tabs.Tab>
					<Tabs.Tab value="likes">Likes</Tabs.Tab>
				</Tabs.List>
				<Tabs.Panel value="posts">
					<InfinitePosts query={postsQuery} queryKey={["posts", user.handle]} />
				</Tabs.Panel>
				<Tabs.Panel value="likes">
					<InfinitePosts query={likesQuery} queryKey={["posts", user.handle, "likes"]} />
				</Tabs.Panel>
			</Tabs>
		</Wrapper>
	);
}
