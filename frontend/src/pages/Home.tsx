import { Container } from "@mantine/core";
import { CreatePost } from "../components/post/CreatePost";
import { Posts } from "../components/Posts";
import { QueryKey } from "../query-key";
import { getMeFeed, postPosts } from "../spec.gen";

const Home = () => {
	return (
		<>
			<Container p="md" pos="relative">
				<CreatePost mutationFn={postPosts} queryKey={[QueryKey.FEED]} />
			</Container>
			<Posts query={getMeFeed} queryKey={[QueryKey.FEED]} />
		</>
	);
};

export default Home;
