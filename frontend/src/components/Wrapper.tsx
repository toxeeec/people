import { Container, Paper, PaperProps } from "@mantine/core";

interface WrapperProps {
	children: React.ReactNode;
}
export const Wrapper = ({ children, ...style }: WrapperProps & PaperProps) => {
	return (
		<Container p={0}>
			<Paper withBorder radius={0} mt={-1} {...style}>
				{children}
			</Paper>
		</Container>
	);
};
