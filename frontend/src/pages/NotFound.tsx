import { Flex, Title, Text, Group, Button } from "@mantine/core";
import { Link } from "react-router-dom";

export default function NotFound() {
	return (
		<Flex direction="column" h="100%" justify="center">
			<Text align="center" color="gray" lh={1.2} weight={900} style={{ fontSize: 200 }}>
				404
			</Text>
			<Title align="center" weight={900}>
				Page not found
			</Title>
			<Text color="dimmed" size="lg" align="center" my="xl">
				The page you are looking for does not exist. It might have been moved or deleted.
			</Text>
			<Group position="center">
				<Button variant="subtle" size="md" component={Link} to={"/home"}>
					Take me back to home page
				</Button>
			</Group>
		</Flex>
	);
}
