import { useMemo } from "react";
import {
	createBrowserRouter,
	redirect,
	RouterProvider,
} from "react-router-dom";
import { CenterLoader } from "./components/CenterLoader";
import { AuthContext } from "./context/AuthContext";
import { PostsContextProvider } from "./context/PostsContext";
import { Users, UsersContextProvider } from "./context/UsersContext";
import {
	createRequestInterceptor,
	createResponseInterceptor,
} from "./custom-instance";
import { useAuth } from "./hooks/useAuth";
import Layout from "./layout";
import Auth from "./pages/Auth";
import Follows from "./pages/Follows";
import Home from "./pages/Home";
import Post from "./pages/Post";
import Profile from "./pages/Profile";
import { useGetUsersHandle } from "./spec.gen";

const App = () => {
	const { getAuth, setAuth, clearAuth, isAuthenticated } = useAuth();
	const users: Users = useMemo(() => ({}), []);
	const { isLoading } = useGetUsersHandle(getAuth().handle!, {
		query: {
			enabled: isAuthenticated,
			onSuccess: (u) => (users[u.handle] = u),
			onError: () => clearAuth(),
		},
	});

	createRequestInterceptor(getAuth);
	createResponseInterceptor(getAuth, setAuth, clearAuth);

	const router = createBrowserRouter([
		{
			index: true,
			element: <Auth />,
			loader: () => {
				if (isAuthenticated) {
					return redirect("/home");
				}
				return null;
			},
		},
		{
			element: <Layout />,
			children: [
				{
					path: "/home",
					element: <Home />,
					loader: () => {
						if (!isAuthenticated) {
							return redirect("/");
						}
						return null;
					},
				},
				{
					path: "/:handle",
					element: <Profile />,
				},
				{
					path: "/:handle/:postID",
					element: <Post />,
				},
				{
					path: "/:handle/following",
					element: <Follows value={"following"} />,
				},
				{
					path: "/:handle/followers",
					element: <Follows value={"followers"} />,
				},
			],
		},
	]);

	return isLoading && isAuthenticated ? (
		<CenterLoader />
	) : (
		<AuthContext.Provider
			value={{ getAuth, setAuth, clearAuth, isAuthenticated }}
		>
			<UsersContextProvider initialUsers={users}>
				<PostsContextProvider>
					<RouterProvider router={router} />
				</PostsContextProvider>
			</UsersContextProvider>
		</AuthContext.Provider>
	);
};

export default App;
