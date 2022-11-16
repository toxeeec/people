import { HoverCard } from "@mantine/core";
import { useContext, useState } from "react";
import UsersContext from "../../context/UsersContext";
import { User } from "../../models";
import AccountInfo from "../AccountInfo";
import FollowButton from "../FollowButton";

interface ProfileHoverCardProps {
	handle: string;
	children: React.ReactNode;
}

export default function ProfileHoverCard({
	children,
	handle,
}: ProfileHoverCardProps) {
	const usersCtx = useContext(UsersContext)!;
	const [user, setUser] = useState(usersCtx.users.get(handle)!);

	const updateUser = (u: Partial<User>) => {
		setUser(usersCtx.setUser(handle, u));
	};

	return (
		<HoverCard onOpen={() => setUser(usersCtx.users.get(handle)!)}>
			<HoverCard.Target>{children}</HoverCard.Target>
			<HoverCard.Dropdown>
				<AccountInfo user={user!}>
					<FollowButton user={user} updateUser={updateUser} />
				</AccountInfo>
			</HoverCard.Dropdown>
		</HoverCard>
	);
}
