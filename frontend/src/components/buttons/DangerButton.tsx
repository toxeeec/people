import { Button, Tooltip } from "@mantine/core";

type DangerButtonProps = {
	label?: string;
	onClick: () => void;
	text: string;
};

export function DangerButton({ label, onClick, text }: DangerButtonProps) {
	return (
		<Tooltip label={label} zIndex={9999} hidden={!label}>
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
}
