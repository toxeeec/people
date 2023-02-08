import { ActionIcon } from "@mantine/core";
import { IconMessage } from "@tabler/icons";
import { useContext } from "react";
import { Link } from "react-router-dom";
import { AuthContext } from "../context/AuthContext";

export const MessagesIcon = () => {
	const { isAuthenticated } = useContext(AuthContext);
	return (
		<>
			{isAuthenticated && (
				<ActionIcon component={Link} to={`/messages`}>
					<IconMessage />
				</ActionIcon>
			)}
		</>
	);
};
