import { Avatar as MantineAvatar, type AvatarProps } from "@mantine/core";
import { type PolymorphicComponentProps } from "@mantine/utils";
import { baseURL } from "@/custom-instance";
import { type User } from "@/models";

export function Avatar({
	user,
	...props
}: Omit<PolymorphicComponentProps<"div", AvatarProps>, "radius" | "src"> & {
	user: User;
}) {
	const src = user.image && (user.image?.startsWith("blob:") ? user.image : baseURL + user.image);
	return <MantineAvatar {...props} radius={9999} src={src} />;
}
