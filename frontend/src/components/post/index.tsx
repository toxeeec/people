import { Avatar, Text, Group, Paper } from "@mantine/core";
import { forwardRef } from "react";
import { Link, useNavigate } from "react-router-dom";
import { Post as PostsData } from "../../models";
import { stopPropagation } from "../../utils";
import PostActions from "./PostActions";
import ProfileHoverCard from "./ProfileHoverCard";

interface PostProps {
	data: PostsData;
}

const Post = forwardRef<HTMLDivElement, PostProps>(({ data }, ref) => {
	const { content, user, id } = data;
	const navigate = useNavigate();
	return (
		<Paper
			p="xs"
			radius="xs"
			withBorder
			ref={ref}
			onClick={() => {
				navigate(`/${user!.handle}/${id}`);
			}}
			style={{ cursor: "pointer" }}
		>
			<Group align="center">
				<ProfileHoverCard handle={user!.handle}>
					<Avatar
						radius="xl"
						size="md"
						component={Link}
						to={`/${user!.handle}`}
						onClick={stopPropagation}
					/>
				</ProfileHoverCard>
				<ProfileHoverCard handle={user!.handle}>
					<Text
						component={Link}
						to={`/${user!.handle}`}
						weight="bold"
						onClick={stopPropagation}
					>
						{user?.handle}
					</Text>
				</ProfileHoverCard>
			</Group>
			<Text my="xs">{content}</Text>
			<PostActions {...data} />
		</Paper>
	);
});

Post.displayName = "Post";
export default Post;
