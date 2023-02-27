import {
	Header as MantineHeader,
	Space,
	Group,
	Text,
	UnstyledButton,
	ActionIcon,
	Indicator,
} from "@mantine/core";
import { useWindowScroll } from "@mantine/hooks";
import { useContext, useEffect, useState } from "react";
import { useLocation } from "react-router";
import { type User } from "@/models";
import { Drawer } from "@/layout/Drawer";
import { Avatar } from "@/components/user";
import { RouteContext } from "@/context/RouteContext";
import { AuthContext } from "@/context/AuthContext";
import { Link, useNavigate } from "react-router-dom";
import { IconArrowLeft, IconMessage, IconSearch } from "@tabler/icons";
import { NotificationsContext } from "@/context/NotificationsContext";

export const HEADER_HEIGHT = 60;

function BackButton() {
	const navigate = useNavigate();
	const { key } = useLocation();
	const handleClick = () => {
		key === "default" ? navigate("home") : navigate(-1);
	};
	return (
		<ActionIcon onClick={handleClick} radius="xl">
			<IconArrowLeft size={20} />
		</ActionIcon>
	);
}

function MessagesButton() {
	const { unreadCount, clearUnreadCount } = useContext(NotificationsContext);
	const location = useLocation();
	useEffect(() => {
		if (location.pathname.includes("/messages")) {
			clearUnreadCount();
		}
	}, [location, clearUnreadCount]);
	return (
		<Indicator inline label={unreadCount} size={16} overflowCount={9} showZero={false} dot={false}>
			<ActionIcon component={Link} to={`/messages`}>
				<IconMessage />
			</ActionIcon>
		</Indicator>
	);
}

type LayoutHeaderProps = {
	user: User;
};

export function Header({ user }: LayoutHeaderProps) {
	const [opened, setOpened] = useState(false);
	const location = useLocation();
	const [, scrollTo] = useWindowScroll();
	const { isAuthenticated } = useContext(AuthContext);
	const { routeName } = useContext(RouteContext);

	const isHome = location.pathname.includes("/home");
	const isSettings = location.pathname.includes("/settings");
	const isSearch = location.pathname.includes("/search");
	const isMessages = location.pathname.includes("/messages");

	useEffect(() => {
		setOpened(false);
	}, [location]);

	return (
		<>
			<Space h={HEADER_HEIGHT} />
			<MantineHeader height={HEADER_HEIGHT} fixed withBorder>
				<Group h="100%" align="center" px="xs" position="apart">
					<Group>
						{isHome && isAuthenticated ? (
							<Avatar user={user} onClick={() => setOpened(true)} style={{ cursor: "pointer" }} />
						) : (
							<BackButton />
						)}
						<UnstyledButton onClick={() => scrollTo({ y: 0 })}>
							<Text fz="xl" fw={700}>
								{routeName}
							</Text>
						</UnstyledButton>
					</Group>
					<Group hidden={isSearch || isMessages || isSettings}>
						{isAuthenticated && <MessagesButton />}
						<ActionIcon component={Link} to={`/search/posts`}>
							<IconSearch />
						</ActionIcon>
					</Group>
				</Group>
			</MantineHeader>
			{isAuthenticated && <Drawer user={user} isOpened={opened} setIsOpened={setOpened} />}
		</>
	);
}
