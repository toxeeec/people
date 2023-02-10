import { Container } from "@mantine/core";
import { CreatePost } from "../components/post/CreatePost";
import { Posts } from "../components/Posts";
import { Wrapper } from "../components/Wrapper";
import { getMeFeed, postPosts } from "../spec.gen";

const Home = () => {
	return (
		<Wrapper>
			<Container p="md" pos="relative">
				<CreatePost mutationFn={postPosts} placeholder={"Create post"} />
			</Container>
			<Posts query={getMeFeed} queryKey={["posts", "feed"]} />
		</Wrapper>
	);
};

export default Home;
