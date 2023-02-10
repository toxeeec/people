import { Box, Container, Paper, Tabs } from "@mantine/core";
import { useDebouncedValue } from "@mantine/hooks";
import { useCallback, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Posts, Query as PostsQuery } from "../components/Posts";
import { Users, Query as UsersQuery } from "../components/Users";
import { Search as SearchComponent } from "../components/Search";
import { getPostsSearch, getUsersSearch } from "../spec.gen";
import { Wrapper } from "../components/Wrapper";

export type SearchPage = "posts" | "people";

const Search = ({ value }: { value: SearchPage }) => {
	const navigate = useNavigate();
	const [query, setQuery] = useState("");
	const [debounced] = useDebouncedValue(query, 200);
	const postsQuery: PostsQuery = useCallback(
		(params) => getPostsSearch({ query: debounced, ...params }),
		[debounced]
	);
	const usersQuery: UsersQuery = useCallback(
		(params) => getUsersSearch({ query: debounced, ...params }),
		[debounced]
	);

	return (
		<Tabs
			value={value}
			onTabChange={(value) => navigate(`/search/${value}`, { replace: true })}
		>
			<Wrapper>
				<Container pos="fixed" w="100%" p={0} style={{ zIndex: 1 }} m={-1}>
					<Paper radius={0} withBorder>
						<SearchComponent value={query} setValue={setQuery} />
						<Tabs.List grow position="center">
							<Tabs.Tab value="posts">Posts</Tabs.Tab>
							<Tabs.Tab value="people">People</Tabs.Tab>
						</Tabs.List>
					</Paper>
				</Container>

				<Box pt="78px">
					<Tabs.Panel value="posts">
						<Posts
							queryKey={["posts", debounced]}
							enabled={value === "posts" && debounced.length > 0}
							query={postsQuery}
						/>
					</Tabs.Panel>
					<Tabs.Panel value="people">
						<Users
							queryKey={["users", debounced]}
							enabled={value === "people" && debounced.length > 0}
							query={usersQuery}
						/>
					</Tabs.Panel>
				</Box>
			</Wrapper>
		</Tabs>
	);
};

export default Search;
