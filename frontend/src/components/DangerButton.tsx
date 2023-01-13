import { Button, Tooltip } from "@mantine/core";

interface DangerButtonProps {
	label?: string;
	onClick: () => void;
	text: string;
}

export const DangerButton = ({ label, onClick, text }: DangerButtonProps) => {
	return (
		<Tooltip label={label} zIndex={9999} display={label ? "unset" : "none"}>
			<Button
				fullWidth
				variant="subtle"
				c="red"
				styles={{ inner: { justifyContent: "start" } }}
				onClick={onClick}
				mb="md"
			>
				{text}
			</Button>
		</Tooltip>
	);
};
