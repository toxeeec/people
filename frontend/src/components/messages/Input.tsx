import { ActionIcon, Flex, Textarea } from "@mantine/core";
import { IconSend } from "@tabler/icons";
import {
	ChangeEvent,
	Dispatch,
	SetStateAction,
	useContext,
	useMemo,
} from "react";
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
		sendMessage({ to, message });
		setMessage("");
	};

	const handleChange = (e: ChangeEvent<HTMLTextAreaElement>) => {
		if (e.currentTarget.value.endsWith("\n")) {
			if (empty) return;
			handleSubmit();
			return;
		}
		setMessage(e.currentTarget.value);
	};

	return (
		<Flex bottom={0} w="100%" align="flex-end" pt="sm">
			<Textarea
				autosize
				minRows={1}
				maxRows={5}
				style={{ flex: 1 }}
				value={message}
				onChange={handleChange}
			/>
			{!empty && (
				<ActionIcon m={4} onClick={handleSubmit}>
					<IconSend />
				</ActionIcon>
			)}
		</Flex>
	);
};
