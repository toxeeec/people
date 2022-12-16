import {
	Avatar,
	Header,
	Space,
	Group,
	Text,
	UnstyledButton,
} from "@mantine/core";
import { useEffect, useState } from "react";
import { useLocation, useParams } from "react-router";
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

	useEffect(() => {
		if (location.pathname === "/home") {
			setRoute("Home");
		} else if (params.postID) {
			setRoute("Post");
		} else if (params.handle) {
			setRoute(params.handle);
		}
	}, [params, location]);

	useEffect(() => {
		setOpened(false);
	}, [location]);

	const isHome = route === "Home";
	return (
		<>
			<Space h={60} />
			<Header height={60} fixed>
				<Group h={60} align="center" px="xs">
					{isHome && isAuthenticated ? (
						<UnstyledButton onClick={() => setOpened(true)}>
							<Avatar radius="xl" onClick={() => setOpened(true)} />
						</UnstyledButton>
					) : (
						<LayoutGoBack />
					)}
					<Text fz="xl" fw={700}>
						{route}
					</Text>
				</Group>
			</Header>
			{isAuthenticated ? (
				<LayoutDrawer isOpened={opened} setIsOpened={setOpened} />
			) : null}
		</>
	);
};
