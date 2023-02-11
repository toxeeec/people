import { ActionIcon, Flex, Textarea } from "@mantine/core";
import { useEventListener } from "@mantine/hooks";
import { IconSend } from "@tabler/icons";
import { Dispatch, SetStateAction, useContext, useMemo } from "react";
import { NotificationsContext } from "../../context/NotificationsContext";

interface InputProps {
	message: string;
	setMessage: Dispatch<SetStateAction<string>>;
	to: string;
}
export const Input = ({ message, setMessage, to }: InputProps) => {
	const { sendMessage } = useContext(NotificationsContext);
	const empty = useMemo(() => message.trim().length === 0, [message]);
	const handleSubmit = () => {
		if (empty) return;
		sendMessage({ to, message: message.trim() });
		setMessage("");
	};

	const handleKeypress = (e: KeyboardEvent) => {
		if (e.key !== "Enter") return;
		e.preventDefault();
		handleSubmit();
	};
	const ref = useEventListener("keypress", handleKeypress);

	return (
		<Flex ref={ref} bottom={0} w="100%" align="flex-end" pt="sm">
			<Textarea
				autosize
				minRows={1}
				maxRows={5}
				style={{ flex: 1 }}
				value={message}
				onChange={(e) => setMessage(e.currentTarget.value)}
			/>
			{!empty && (
				<ActionIcon m={4} onClick={handleSubmit}>
					<IconSend />
				</ActionIcon>
			)}
		</Flex>
	);
};
