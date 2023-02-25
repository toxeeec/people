import { Group, Stack, Text } from "@mantine/core";
import { Link, useNavigate } from "react-router-dom";
import { UserHoverCard } from "@/components/user/UserHoverCard";
import { forwardRef, useContext } from "react";
import { AuthContext } from "@/context/AuthContext";
import { FollowButton } from "@/components/buttons";
import { type User as UserType } from "@/models";
import { Avatar } from "@/components/user";

type UserProps = {
	user: UserType;
};
export const User = forwardRef<HTMLDivElement, UserProps>(({ user }, ref) => {
	const { getAuth } = useContext(AuthContext);
	const navigate = useNavigate();
	const ownProfile = user.handle === getAuth().handle;
	return (
		<>
			<Group
				ref={ref}
				p="md"
				align="start"
				position="apart"
				onClick={(e) => {
					if (e.target === e.currentTarget) {
						navigate(`/${user.handle}`);
					}
				}}
				style={{ cursor: "pointer" }}
			>
				<Group align="stretch">
					<UserHoverCard user={user}>
						<Link to={`/${user.handle}`}>
							<Avatar size="lg" user={user} />
						</Link>
					</UserHoverCard>
					<Stack>
						<UserHoverCard user={user}>
							<Text component={Link} to={`/${user.handle}`} weight="bold">
								@{user.handle}
							</Text>
						</UserHoverCard>
					</Stack>
				</Group>
				{!ownProfile && <FollowButton user={user} />}
			</Group>
		</>
	);
});

User.displayName = "User";
