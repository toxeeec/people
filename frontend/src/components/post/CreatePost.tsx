import {
	ActionIcon,
	Button,
	FileButton,
	Flex,
	Group,
	LoadingOverlay,
	Text,
	Textarea,
} from "@mantine/core";
import { showNotification } from "@mantine/notifications";
import { IconPhoto } from "@tabler/icons";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { forwardRef, useCallback, useRef, useState } from "react";
import { type ErrorType } from "@/custom-instance";
import { type ImageResponse, type NewPost, type PostResponse } from "@/models";
import { CreateImages } from "@/components/images/CreateImages";

const MAX_POST_LENGTH = 280;
const MAX_IMAGES_COUNT = 4;

export type MutationFn = (newPost: NewPost) => Promise<PostResponse>;

type CreatePostProps = {
	mutationFn: MutationFn;
	handleClose?: () => void;
	placeholder: string;
};

export const CreatePost = forwardRef<HTMLTextAreaElement, CreatePostProps>(
	({ mutationFn, handleClose, placeholder }, ref) => {
		const [content, setContent] = useState("");
		const [error, setError] = useState("");
		const queryClient = useQueryClient();
		const { mutate } = useMutation({
			onSuccess: () => {
				setContent("");
				queryClient.invalidateQueries();
				handleClose?.();
			},
			mutationFn,
			retry: 1,
			onError: (error) => {
				const msg = (error as ErrorType<Error>).response?.data.message;
				setError(msg ?? "Unknown Error");
			},
		});

		const [files, setFiles] = useState<File[]>([]);
		const [imageResponses, setImageResponses] = useState<Promise<ImageResponse>[]>([]);
		const [loading, setLoading] = useState(false);
		const resetRef = useRef<() => void>(null);
		const handleChange = (newFiles: File[]) => {
			if (newFiles.length + files.length > MAX_IMAGES_COUNT) {
				showNotification({
					message: `You can only select up to ${MAX_IMAGES_COUNT} photos`,
					disallowClose: true,
					color: "red",
				});
				return;
			}
			setFiles((files) => [...files, ...newFiles]);
			resetRef.current?.();
		};

		const trimmed = content.trim();
		const handleSubmit = useCallback(() => {
			setLoading(true);
			Promise.all(imageResponses).then((irs) => {
				mutate(
					{ content: trimmed, images: irs.map((ir) => ir.id) },
					{
						onSettled: () => {
							setFiles([]);
							setImageResponses([]);
							setLoading(false);
						},
					}
				);
			});
		}, [trimmed, mutate, imageResponses]);

		const postDisabled =
			(files.length === 0 && trimmed.length === 0) || trimmed.length > MAX_POST_LENGTH;
		const filesDisabled = files.length >= MAX_IMAGES_COUNT;

		return (
			<>
				<LoadingOverlay visible={loading} />
				<Textarea
					placeholder={placeholder}
					value={content}
					autosize
					minRows={4}
					variant="unstyled"
					onChange={(e) => setContent(e.currentTarget.value)}
					error={error}
					ref={ref}
					data-autofocus
				/>
				<CreateImages files={files} setFiles={setFiles} setImageResponses={setImageResponses} />
				<Flex justify="space-between" align="center" mt="md">
					<Group>
						<FileButton
							resetRef={resetRef}
							onChange={handleChange}
							disabled={filesDisabled}
							accept="image/png,image/jpeg,image/webp"
							multiple
						>
							{(props) => (
								<ActionIcon {...props} disabled={filesDisabled}>
									<IconPhoto />
								</ActionIcon>
							)}
						</FileButton>
					</Group>
					<Group>
						<Text mr="md">{`${trimmed.length}/${MAX_POST_LENGTH}`}</Text>
						<Button onClick={handleSubmit} variant="filled" radius="xl" disabled={postDisabled}>
							Post
						</Button>
					</Group>
				</Flex>
			</>
		);
	}
);

CreatePost.displayName = "CreatePost";
