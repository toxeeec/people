import { Container, Paper, type PaperProps } from "@mantine/core";

type WrapperProps = {
	children: React.ReactNode;
};
export function Wrapper({ children, ...style }: WrapperProps & PaperProps) {
	return (
		<Container p={0}>
			<Paper withBorder radius={0} mt={-1} {...style}>
				{children}
			</Paper>
		</Container>
	);
}
