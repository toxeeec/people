import { Avatar as MantineAvatar, AvatarProps } from "@mantine/core";
import { PolymorphicComponentProps } from "@mantine/utils";
import { baseURL } from "./custom-instance";
import { User } from "./models";

export const Avatar = ({
	user,
	...props
}: Omit<PolymorphicComponentProps<"div", AvatarProps>, "radius" | "src"> & {
	user: User;
}) => {
	return (
		<MantineAvatar
			{...props}
			radius={9999}
			src={
				user.image &&
				(user.image?.includes("blob:") ? user.image : baseURL + user.image)
			}
		/>
	);
};
