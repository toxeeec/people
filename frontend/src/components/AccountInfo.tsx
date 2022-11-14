import { Avatar, Group } from "@mantine/core";
import { useContext } from "react";
import UsersContext from "../context/UsersContext";

interface AccountInfoProps {
	handle: string;
	children?: React.ReactNode;
}

export default function AccountInfo({ handle, children }: AccountInfoProps) {
	const user = useContext(UsersContext)?.users[handle];
	return (
		<>
			<Group align="start" position="apart">
				<Avatar size="lg" radius="xl" mb="xs" />
				{children}
			</Group>
			<b>@{user?.handle}</b>
			<Group mt="xs">
				<span>
					<b>{user?.following}</b> Following
				</span>
				<span>
					<b>{user?.followers}</b>
					{user?.followers === 1 ? " Follower" : " Followers"}
				</span>
			</Group>
		</>
	);
}
