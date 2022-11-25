import { Avatar, Text, Group, Paper } from "@mantine/core";
import { forwardRef, useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { Post as PostData } from "../../models";
import { stopPropagation } from "../../utils";
import PostActions from "./PostActions";
import ProfileHoverCard from "./ProfileHoverCard";

interface PostProps {
	post: PostData;
	clickable?: boolean;
}

const Post = forwardRef<HTMLDivElement, PostProps>(
	({ post: initialPost, clickable }, ref) => {
		const [post, setPost] = useState(initialPost);
		const navigate = useNavigate();
		const props = clickable
			? {
					onClick: () => {
						navigate(`/${post.user!.handle}/${post.id}`);
					},
					style: { cursor: "pointer" },
			  }
			: {};

		return (
			<Paper p="xs" radius="xs" withBorder ref={ref} {...props}>
				<Group align="center">
					<ProfileHoverCard handle={post.user!.handle}>
						<Avatar
							radius="xl"
							size="md"
							component={Link}
							to={`/${post.user!.handle}`}
							onClick={stopPropagation}
						/>
					</ProfileHoverCard>
					<ProfileHoverCard handle={post.user!.handle}>
						<Text
							component={Link}
							to={`/${post.user!.handle}`}
							weight="bold"
							onClick={stopPropagation}
						>
							{post.user?.handle}
						</Text>
					</ProfileHoverCard>
				</Group>
				<Text my="xs">{post.content}</Text>
				<PostActions post={post} setPost={setPost} />
			</Paper>
		);
	}
);

Post.displayName = "Post";
export default Post;
