import { InfiniteUsers, type UsersQuery } from "@/components/user";
import { Wrapper } from "@/components/utils";
import { Tabs } from "@mantine/core";
import { useNavigate, useParams } from "react-router-dom";
import { getUsersHandleFollowers, getUsersHandleFollowing } from "@/spec.gen";
import { useContext, useEffect } from "react";
import { RouteContext } from "@/context/RouteContext";

export type FollowsPage = "followers" | "following";

function Follows({ value }: { value: FollowsPage }) {
	const navigate = useNavigate();
	const params = useParams();
	const handle = params.handle ?? "";
	const { setRouteName } = useContext(RouteContext);
	useEffect(() => {
		setRouteName(handle);
	}, [setRouteName, handle]);

	const queryFollowing: UsersQuery = (queryParams) => getUsersHandleFollowing(handle, queryParams);
	const queryFollowers: UsersQuery = (queryParams) => getUsersHandleFollowers(handle, queryParams);

	return (
		<Wrapper>
			<Tabs
				value={value}
				onTabChange={(value) => navigate(`/${params.handle}/${value}`, { replace: true })}
			>
				<Tabs.List grow position="center">
					<Tabs.Tab value="following">Following</Tabs.Tab>
					<Tabs.Tab value="followers">Followers</Tabs.Tab>
				</Tabs.List>
				<Tabs.Panel value="following">
					<InfiniteUsers queryKey={["users", handle, "following"]} query={queryFollowing} />
				</Tabs.Panel>
				<Tabs.Panel value="followers">
					<InfiniteUsers queryKey={["users", handle, "followers"]} query={queryFollowers} />
				</Tabs.Panel>
			</Tabs>
		</Wrapper>
	);
}

export default Follows;
