import { Header, Space, Group, Text, UnstyledButton } from "@mantine/core";
import { useWindowScroll } from "@mantine/hooks";
import { useEffect, useState } from "react";
import { useLocation, useParams } from "react-router";
import { Avatar } from "../Avatar";
import { User } from "../models";
import { LayoutDrawer } from "./LayoutDrawer";
import { LayoutGoBack } from "./LayoutGoBack";
import { MessagesIcon } from "./MessagesIcon";
import { SearchIcon } from "./SearchIcon";

interface LayoutHeaderProps {
	isAuthenticated: boolean;
	user: User;
}

export const LayoutHeader = ({ user, isAuthenticated }: LayoutHeaderProps) => {
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
		} else if (location.pathname === "/settings") {
			setRoute("Settings");
		} else if (location.pathname.includes("messages")) {
			setRoute("Messages");
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
	const isSettings = route === "Settings";
	const isSearch = location.pathname.includes("search");
	const isMessages = location.pathname.includes("messages");
	return (
		<>
			<Space h={60} />
			<Header height={60} fixed withBorder>
				<Group h={60} align="center" px="xs" position="apart">
					<Group>
						{isHome && isAuthenticated ? (
							<UnstyledButton onClick={() => setOpened(true)}>
								<Avatar user={user} onClick={() => setOpened(true)} />
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
					<Group
						display={isSearch || isMessages || isSettings ? "none" : "flex"}
					>
						<SearchIcon />
						<MessagesIcon />
					</Group>
				</Group>
			</Header>
			{isAuthenticated ? (
				<LayoutDrawer user={user} isOpened={opened} setIsOpened={setOpened} />
			) : null}
		</>
	);
};
