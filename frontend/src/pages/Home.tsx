import { Container } from "@mantine/core";
import { CreatePost } from "../components/post/CreatePost";
import { Posts } from "../components/Posts";
import { Wrapper } from "../components/Wrapper";
import { QueryKey } from "../query-key";
import { getMeFeed, postPosts } from "../spec.gen";

const Home = () => {
	return (
		<Wrapper>
			<Container p="md" pos="relative">
				<CreatePost
					mutationFn={postPosts}
					queryKey={[QueryKey.FEED]}
					placeholder={"Create post"}
				/>
			</Container>
			<Posts query={getMeFeed} queryKey={[QueryKey.FEED]} />
		</Wrapper>
	);
};

export default Home;
