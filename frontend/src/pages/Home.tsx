import { CreatePost } from "../components/post/CreatePost";
import { Posts } from "../components/Posts";
import { QueryKey } from "../query-key";
import { getMeFeed, postPosts } from "../spec.gen";

const Home = () => {
	return (
		<>
			<CreatePost mutationFn={postPosts} queryKey={[QueryKey.FEED]} />
			<Posts query={getMeFeed} queryKey={[QueryKey.FEED]} />
		</>
	);
};

export default Home;
