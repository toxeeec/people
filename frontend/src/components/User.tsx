import { Group, Stack, Text } from "@mantine/core";
import { Link, useNavigate } from "react-router-dom";
import { UserHoverCard } from "./UserHoverCard";
import { forwardRef, useContext } from "react";
import { AuthContext } from "../context/AuthContext";
import { FollowButton } from "./FollowButton";
import { User as UserType } from "../models";
import { Avatar } from "../Avatar";

interface UserProps {
	user: UserType;
}
export const User = forwardRef<HTMLDivElement, UserProps>(({ user }, ref) => {
	const { isAuthenticated, getAuth } = useContext(AuthContext);
	const navigate = useNavigate();
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
				{isAuthenticated && getAuth().handle !== user.handle ? (
					<FollowButton user={user} />
				) : null}
			</Group>
		</>
	);
});

User.displayName = "User";
