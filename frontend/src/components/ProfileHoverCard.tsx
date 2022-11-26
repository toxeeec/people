import { HoverCard } from "@mantine/core";
import { useContext, useState } from "react";
import UsersContext from "../context/UsersContext";
import Profile from "./Profile";

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

	return (
		<HoverCard onOpen={() => setUser(usersCtx.users.get(handle)!)}>
			<HoverCard.Target>{children}</HoverCard.Target>
			<HoverCard.Dropdown p={0}>
				<Profile user={user} setUser={setUser} />
			</HoverCard.Dropdown>
		</HoverCard>
	);
}
