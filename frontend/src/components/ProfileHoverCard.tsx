import { HoverCard } from "@mantine/core";
import { Profile } from "./Profile";

interface ProfileHoverCardProps {
	handle: string;
	children: React.ReactNode;
}

export const ProfileHoverCard = ({
	handle,
	children,
}: ProfileHoverCardProps) => {
	return (
		<HoverCard>
			<HoverCard.Target>{children}</HoverCard.Target>
			<HoverCard.Dropdown p={0}>
				<Profile handle={handle} />
			</HoverCard.Dropdown>
		</HoverCard>
	);
};
