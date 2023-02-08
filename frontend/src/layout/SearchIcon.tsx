import { ActionIcon } from "@mantine/core";
import { IconSearch } from "@tabler/icons";
import { Link } from "react-router-dom";

export const SearchIcon = () => {
	return (
		<ActionIcon component={Link} to={`/search/posts`}>
			<IconSearch />
		</ActionIcon>
	);
};
