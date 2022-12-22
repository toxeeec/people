import {
	ActionIcon,
	Button,
	Container,
	FileButton,
	Flex,
	Group,
	LoadingOverlay,
	Text,
	Textarea,
} from "@mantine/core";
import { showNotification } from "@mantine/notifications";
import { IconPhoto } from "@tabler/icons";
import { QueryKey, useMutation, useQueryClient } from "@tanstack/react-query";
import { MouseEvent, useCallback, useContext, useRef, useState } from "react";
import { PostsContext } from "../../context/PostsContext";
import { UsersContext } from "../../context/UsersContext";
import { ErrorType } from "../../custom-instance";
import { ImageResponse, NewPost, PostResponse } from "../../models";
import { CreateImages } from "../images/CreateImages";

export type MutationFn = (newPost: NewPost) => Promise<PostResponse>;
interface CreatePostProps {
	mutationFn: MutationFn;
	queryKey: QueryKey;
}

export const CreatePost = ({ mutationFn, queryKey }: CreatePostProps) => {
	const [content, setContent] = useState("");
	const [error, setError] = useState("");
	const queryClient = useQueryClient();
	const { setUser } = useContext(UsersContext);
	const { setPost } = useContext(PostsContext);
	const { mutate } = useMutation({
		mutationFn,
		retry: 1,
		onSuccess: (postResponse) => {
			setPost(postResponse.data);
			setUser(postResponse.user);
			setContent("");
			queryClient.invalidateQueries({ queryKey });
		},
		onError: (error) => {
			const err = (error as ErrorType<Error>).response?.data.message;
			setError(err!);
		},
	});

	const [loading, setLoading] = useState(false);
	const [files, setFiles] = useState<File[]>([]);
	const [imageResponses, setImageResponses] = useState<
		Promise<ImageResponse>[]
	>([]);

	const resetRef = useRef<() => void>(null);
	const onChange = (newFiles: File[]) => {
		if (newFiles.length + files.length > 4) {
			showNotification({
				message: "You can only select up to 4 photos",
				disallowClose: true,
				color: "red",
			});
			return;
		}
		setFiles((files) => [...files, ...newFiles]);
		resetRef.current?.();
	};

	const handleSubmit = useCallback(
		(e: MouseEvent) => {
			e.stopPropagation();
			setLoading(true);
			Promise.all(imageResponses)
				.then((irs) => {
					mutate({ content: content.trim(), images: irs.map((ir) => ir.id) });
				})
				.finally(() => {
					setFiles([]);
					setImageResponses([]);
					setLoading(false);
				});
		},
		[content, mutate, imageResponses]
	);

	const postDisabled =
		(files.length === 0 && content.trim().length === 0) ||
		content.trim().length > 280;
	const filesDisabled = files.length >= 4;
	return (
		<Container p="lg" pos="relative">
			<LoadingOverlay visible={loading} />
			<Textarea
				placeholder="Create new post"
				value={content}
				variant="unstyled"
				onChange={(e) => setContent(e.currentTarget.value)}
				error={error}
			/>
			<CreateImages
				files={files}
				setFiles={setFiles}
				setImageResponses={setImageResponses}
			/>
			<Flex justify="space-between" align="center" mt="md">
				<Group>
					<FileButton
						resetRef={resetRef}
						onChange={onChange}
						disabled={filesDisabled}
						accept="image/png,image/jpeg,image/webp"
						multiple
					>
						{(props) => (
							<ActionIcon {...props} disabled={filesDisabled}>
								<IconPhoto size={24} />{" "}
							</ActionIcon>
						)}
					</FileButton>
				</Group>
				<Group>
					<Text mr="md">{`${content.trim().length}/280`}</Text>
					<Button
						onClick={handleSubmit}
						variant="filled"
						radius="xl"
						disabled={postDisabled}
					>
						Post
					</Button>
				</Group>
			</Flex>
		</Container>
	);
};
