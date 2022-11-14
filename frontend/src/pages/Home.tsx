import Posts from "../components/Posts";
import { getMeFeed } from "../spec.gen";

export default function Home() {
	return <Posts query={getMeFeed} />;
}
