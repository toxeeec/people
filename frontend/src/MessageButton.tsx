import { getUsersHandleThread } from "@/spec.gen";
import { ActionIcon } from "@mantine/core";
import { IconMessage } from "@tabler/icons";
import { useNavigate } from "react-router-dom";

type MessageButtonProps = {
	handle: string;
};

export function MessageButton({ handle }: MessageButtonProps) {
	const navigate = useNavigate();
	const handleClick = () => {
		getUsersHandleThread(handle)
			.then((thread) => navigate(`/messages/${thread.id}`))
			.catch((e) => e);
	};
	return (
		<ActionIcon onClick={handleClick}>
			<IconMessage />
		</ActionIcon>
	);
}
