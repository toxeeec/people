import { Tabs } from "@mantine/core";
import { useNavigate, useParams } from "react-router-dom";
import Profiles, { Query } from "../components/Profiles";
import { getUsersHandleFollowers, getUsersHandleFollowing } from "../spec.gen";

export enum FollowsPage {
	// eslint-disable-next-line
	Followers = "followers",
	// eslint-disable-next-line
	Following = "following",
}

export default function Follows({ value }: { value: FollowsPage }) {
	const navigate = useNavigate();
	const params = useParams();

	const queryFollowing: Query = (queryParams) => {
		return getUsersHandleFollowing(params.handle!, queryParams);
	};

	const queryFollowers: Query = (queryParams) => {
		return getUsersHandleFollowers(params.handle!, queryParams);
	};

	return (
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
				<Profiles
					query={queryFollowing}
					queryKey={["following", params.handle!]}
				/>
			</Tabs.Panel>
			<Tabs.Panel value="followers">
				<Profiles
					query={queryFollowers}
					queryKey={["followers", params.handle!]}
				/>
			</Tabs.Panel>
		</Tabs>
	);
}
