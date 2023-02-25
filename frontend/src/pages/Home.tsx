import { Wrapper } from "@/components/utils";
import { Container } from "@mantine/core";
import { CreatePost } from "@/components/post";
import { InfinitePosts } from "@/components/post";
import { getMeFeed, postPosts } from "@/spec.gen";
import { useContext, useEffect } from "react";
import { RouteContext } from "@/context/RouteContext";

function Home() {
	const { setRouteName } = useContext(RouteContext);
	useEffect(() => {
		setRouteName("Home");
	}, [setRouteName]);
	return (
		<Wrapper>
			<Container p="md" pos="relative">
				<CreatePost mutationFn={postPosts} placeholder={"Create post"} />
			</Container>
			<InfinitePosts query={getMeFeed} queryKey={["posts", "feed"]} />
		</Wrapper>
	);
}

export default Home;
