import { User } from "../../models";
import { Header as MantineHeader } from "@mantine/core";
import { UserInfo } from "./UserInfo";

interface HeaderProps {
	user?: User;
}

export const Header = ({ user }: HeaderProps) => {
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
};
