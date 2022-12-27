import { Box, Paper, Tabs } from "@mantine/core";
import { useDebouncedValue } from "@mantine/hooks";
import { useCallback, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Posts, Query as PostsQuery } from "../components/Posts";
import { Users, Query as UsersQuery } from "../components/Users";
import { Search as SearchComponent } from "../components/Search";
import { QueryKey } from "../query-key";
import { getPostsSearch, getUsersSearch } from "../spec.gen";

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
		<>
			<SearchComponent value={query} setValue={setQuery} />
			<Tabs
				value={value}
				onTabChange={(value) => navigate(`/search/${value}`, { replace: true })}
			>
				<Paper pos="fixed" w="100%" mt="42px" radius={0} style={{ zIndex: 1 }}>
					<Tabs.List grow position="center">
						<Tabs.Tab value="posts">Posts</Tabs.Tab>
						<Tabs.Tab value="people">People</Tabs.Tab>
					</Tabs.List>
				</Paper>

				<Box pt="78px">
					<Tabs.Panel value="posts">
						<Posts
							enabled={value === "posts" && debounced.length > 0}
							query={postsQuery}
							queryKey={[QueryKey.POSTS_SEARCH, debounced]}
						/>
					</Tabs.Panel>
					<Tabs.Panel value="people">
						<Users
							enabled={value === "people" && debounced.length > 0}
							query={usersQuery}
							queryKey={[QueryKey.USERS_SEARCH, debounced]}
						/>
					</Tabs.Panel>
				</Box>
			</Tabs>
		</>
	);
};

export default Search;
