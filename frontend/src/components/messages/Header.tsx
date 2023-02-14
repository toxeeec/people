import { User } from "../../models";
import { Header as MantineHeader } from "@mantine/core";
import { PostAuthor } from "../post/PostAuthor";

interface HeaderProps {
	user: User;
}

export const Header = ({ user }: HeaderProps) => {
	return (
		<MantineHeader
			height={42}
			pos="absolute"
			display="flex"
			style={{ alignItems: "center" }}
			zIndex={1}
		>
			{user && <PostAuthor user={user} />}
		</MantineHeader>
	);
};
