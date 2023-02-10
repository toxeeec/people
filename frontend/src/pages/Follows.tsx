import { Tabs } from "@mantine/core";
import { useNavigate, useParams } from "react-router-dom";
import { Users, Query } from "../components/Users";
import { Wrapper } from "../components/Wrapper";
import { getUsersHandleFollowers, getUsersHandleFollowing } from "../spec.gen";

export type FollowsPage = "followers" | "following";

const Follows = ({ value }: { value: FollowsPage }) => {
	const navigate = useNavigate();
	const params = useParams();

	const queryFollowing: Query = (queryParams) => {
		return getUsersHandleFollowing(params.handle!, queryParams);
	};

	const queryFollowers: Query = (queryParams) => {
		return getUsersHandleFollowers(params.handle!, queryParams);
	};

	return (
		<Wrapper>
			<Tabs
				value={value}
				onTabChange={(value) =>
					navigate(`/${params.handle}/${value}`, { replace: true })
				}
			>
				<Tabs.List grow position="center">
					<Tabs.Tab value="following">Following</Tabs.Tab>
					<Tabs.Tab value="followers">Followers</Tabs.Tab>
				</Tabs.List>

				<Tabs.Panel value="following">
					<Users
						queryKey={["users", params.handle, "following"]}
						query={queryFollowing}
					/>
				</Tabs.Panel>
				<Tabs.Panel value="followers">
					<Users
						queryKey={["users", params.handle, "followers"]}
						query={queryFollowers}
					/>
				</Tabs.Panel>
			</Tabs>
		</Wrapper>
	);
};

export default Follows;
