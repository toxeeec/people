import { ActionIcon, Box, Flex, Textarea } from "@mantine/core";
import { useEventListener } from "@mantine/hooks";
import { IconMoodSmile, IconSend } from "@tabler/icons";
import Picker from "@emoji-mart/react";
import { type Dispatch, type SetStateAction, useContext, useMemo, useState } from "react";
import { NotificationsContext } from "@/context/NotificationsContext";

type InputProps = {
	message: string;
	setMessage: Dispatch<SetStateAction<string>>;
	threadID: number;
};
export function Input({ message, setMessage, threadID }: InputProps) {
	const { sendMessage } = useContext(NotificationsContext);
	const [opened, setOpened] = useState(false);
	const empty = useMemo(() => message.trim().length === 0, [message]);
	const handleSubmit = () => {
		if (empty) return;
		sendMessage({ threadID, content: message.trim() });
		setMessage("");
	};

	const handleKeypress = (e: KeyboardEvent) => {
		if (e.key !== "Enter") return;
		e.preventDefault();
		handleSubmit();
	};
	const ref = useEventListener("keypress", handleKeypress);

	return (
		<>
			<Flex w="100%" align="flex-end" my="sm" pos="relative">
				<Textarea
					ref={ref}
					autosize
					minRows={1}
					maxRows={5}
					value={message}
					onChange={(e) => setMessage(e.currentTarget.value)}
					ml="xs"
					styles={{
						input: {
							border: "none",
							borderRadius: 21,
							overflow: "hidden",
							paddingRight: 28,
							fontSize: 16,
						},
						root: { flex: 1 },
					}}
				/>
				<Box hidden={!opened} pos="absolute" right={39} bottom={42}>
					<Picker
						onClickOutside={() => opened && setOpened(false)}
						onEmojiSelect={(e: { native: string }) => {
							setMessage((msg) => msg.concat(e.native));
						}}
					/>
				</Box>
				<ActionIcon
					pos="absolute"
					right={39}
					bottom={8}
					ml="auto"
					onClick={() => setOpened((o) => !o)}
				>
					<IconMoodSmile />
				</ActionIcon>
				<ActionIcon
					m={4}
					mb={8}
					onClick={handleSubmit}
					disabled={empty}
					style={{ background: "none", border: "none" }}
				>
					<IconSend />
				</ActionIcon>
			</Flex>
		</>
	);
}
