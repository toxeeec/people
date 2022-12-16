import { Avatar, Group, Stack, Text } from "@mantine/core";
import { Link } from "react-router-dom";
import { ProfileHoverCard } from "./ProfileHoverCard";
import { stopPropagation } from "../utils";
import { forwardRef, useContext } from "react";
import { AuthContext } from "../context/AuthContext";
import { FollowButton } from "./FollowButton";

interface UserProps {
	handle: string;
}
export const User = forwardRef<HTMLDivElement, UserProps>(({ handle }, ref) => {
	const { isAuthenticated, getAuth } = useContext(AuthContext);
	return (
		<>
			<Group ref={ref} p="md" align="start" position="apart">
				<Group align="stretch">
					<ProfileHoverCard handle={handle}>
						<Avatar
							radius="xl"
							size="lg"
							component={Link}
							to={`/${handle}`}
							onClick={stopPropagation}
						/>
					</ProfileHoverCard>
					<Stack>
						<ProfileHoverCard handle={handle}>
							<Text
								component={Link}
								to={`/${handle}`}
								weight="bold"
								onClick={stopPropagation}
							>
								@{handle}
							</Text>
						</ProfileHoverCard>
					</Stack>
				</Group>
				{isAuthenticated && getAuth().handle !== handle ? (
					<FollowButton handle={handle} />
				) : null}
			</Group>
		</>
	);
});

User.displayName = "User";
