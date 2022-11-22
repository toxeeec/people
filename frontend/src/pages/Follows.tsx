import { Tabs } from "@mantine/core";
import { useNavigate, useParams } from "react-router-dom";
import Profiles from "../components/Profiles";
import { getMeFollowers, getMeFollowing } from "../spec.gen";

export enum FollowsPage {
	// eslint-disable-next-line
	Followers = "followers",
	// eslint-disable-next-line
	Following = "following",
}

export default function Follows({
	defaultValue,
}: {
	defaultValue: FollowsPage;
}) {
	const navigate = useNavigate();
	const params = useParams();

	return (
		<Tabs
			defaultValue={defaultValue}
			onTabChange={(value) => navigate(`/${params.handle}/${value}`)}
		>
			<Tabs.List grow position="center">
				<Tabs.Tab value="following">Following</Tabs.Tab>
				<Tabs.Tab value="followers">Followers</Tabs.Tab>
			</Tabs.List>

			<Tabs.Panel value="following">
				<Profiles
					query={getMeFollowing}
					queryKey={["following", params.handle!]}
				/>
			</Tabs.Panel>
			<Tabs.Panel value="followers">
				<Profiles
					query={getMeFollowers}
					queryKey={["followers", params.handle!]}
				/>
			</Tabs.Panel>
		</Tabs>
	);
}
