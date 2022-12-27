import {
	Avatar,
	Header,
	Space,
	Group,
	Text,
	UnstyledButton,
	ActionIcon,
} from "@mantine/core";
import { useWindowScroll } from "@mantine/hooks";
import { IconSearch } from "@tabler/icons";
import { useEffect, useState } from "react";
import { useLocation, useParams } from "react-router";
import { Link } from "react-router-dom";
import { LayoutDrawer } from "./LayoutDrawer";
import { LayoutGoBack } from "./LayoutGoBack";

interface LayoutHeaderProps {
	isAuthenticated: boolean;
}

export const LayoutHeader = ({ isAuthenticated }: LayoutHeaderProps) => {
	const [opened, setOpened] = useState(false);
	const params = useParams();
	const location = useLocation();
	const [route, setRoute] = useState("");
	const [, scrollTo] = useWindowScroll();

	useEffect(() => {
		if (location.pathname === "/home") {
			setRoute("Home");
		} else if (params.postID) {
			setRoute("Post");
		} else if (params.handle) {
			setRoute(params.handle);
		} else {
			setRoute("");
		}
	}, [params, location]);

	useEffect(() => {
		setOpened(false);
	}, [location]);

	const isHome = route === "Home";
	const isSearch = location.pathname.includes("search");
	return (
		<>
			<Space h={60} />
			<Header height={60} fixed>
				<Group h={60} align="center" px="xs" position="apart">
					<Group>
						{isHome && isAuthenticated ? (
							<UnstyledButton onClick={() => setOpened(true)}>
								<Avatar radius="xl" onClick={() => setOpened(true)} />
							</UnstyledButton>
						) : (
							<LayoutGoBack />
						)}
						<UnstyledButton onClick={() => scrollTo({ y: 0 })}>
							<Text fz="xl" fw={700}>
								{route}
							</Text>
						</UnstyledButton>
					</Group>
					<ActionIcon
						component={Link}
						to={`/search/posts`}
						display={isSearch ? "none" : "unset"}
					>
						<IconSearch />
					</ActionIcon>
				</Group>
			</Header>
			{isAuthenticated ? (
				<LayoutDrawer isOpened={opened} setIsOpened={setOpened} />
			) : null}
		</>
	);
};
