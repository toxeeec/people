import { createContext, useState } from "react";
import { Post } from "../models";

export type Posts = { [id: number]: Post | undefined };
interface PostsContextType {
	posts: Posts;
	setPost: (post: Post) => void;
}

export const PostsContext = createContext<PostsContextType>(null!);

interface PostsContextProviderProps {
	children: React.ReactNode;
}

export const PostsContextProvider = ({
	children,
}: PostsContextProviderProps) => {
	const [posts, setPosts] = useState<Posts>({});
	const setPost = (post: Post) => {
		setPosts((posts) => ({ ...posts, [post.id]: post }));
	};
	return (
		<PostsContext.Provider value={{ posts, setPost }}>
			{children}
		</PostsContext.Provider>
	);
};
