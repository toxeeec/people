import PostCreate from "../components/post/PostCreate";
import Posts from "../components/Posts";
import { getMeFeed } from "../spec.gen";

export default function Home() {
	return (
		<>
			<PostCreate />
			<Posts query={getMeFeed} queryKey={["feed"]} />
		</>
	);
}
