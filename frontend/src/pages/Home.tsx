import { PostCreate } from "../components/post/PostCreate";
import { Posts } from "../components/Posts";
import { QueryKey } from "../query-key";
import { getMeFeed, postPosts } from "../spec.gen";

const Home = () => {
	return (
		<>
			<PostCreate mutationFn={postPosts} queryKey={[QueryKey.FEED]} />
			<Posts query={getMeFeed} queryKey={[QueryKey.FEED]} />
		</>
	);
};

export default Home;
