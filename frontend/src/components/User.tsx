import { Avatar, Group, Stack, Text } from "@mantine/core";
import { Link, useNavigate } from "react-router-dom";
import { UserHoverCard } from "./UserHoverCard";
import { stopPropagation } from "../utils";
import { forwardRef, useContext } from "react";
import { AuthContext } from "../context/AuthContext";
import { FollowButton } from "./FollowButton";

interface UserProps {
	handle: string;
}
export const User = forwardRef<HTMLDivElement, UserProps>(({ handle }, ref) => {
	const { isAuthenticated, getAuth } = useContext(AuthContext);
	const navigate = useNavigate();
	return (
		<>
			<Group
				ref={ref}
				p="md"
				align="start"
				position="apart"
				onClick={() => {
					navigate(`/${handle}`);
				}}
				style={{ cursor: "pointer" }}
			>
				<Group align="stretch">
					<UserHoverCard handle={handle}>
						<Avatar
							radius="xl"
							size="lg"
							component={Link}
							to={`/${handle}`}
							onClick={stopPropagation}
						/>
					</UserHoverCard>
					<Stack>
						<UserHoverCard handle={handle}>
							<Text
								component={Link}
								to={`/${handle}`}
								weight="bold"
								onClick={stopPropagation}
							>
								@{handle}
							</Text>
						</UserHoverCard>
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
