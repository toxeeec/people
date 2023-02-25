import { type User } from "@/models";
import { Header as MantineHeader } from "@mantine/core";
import { UserInfo } from "@/components/messages/UserInfo";

type HeaderProps = {
	user?: User;
};

export function Header({ user }: HeaderProps) {
	return (
		<>
			{user && (
				<MantineHeader
					height={42}
					pos="absolute"
					display="flex"
					style={{ alignItems: "center" }}
					zIndex={1}
				>
					{user && <UserInfo user={user} />}
				</MantineHeader>
			)}
		</>
	);
}
