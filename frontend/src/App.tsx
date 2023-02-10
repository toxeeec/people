import { useContext } from "react";
import {
	createBrowserRouter,
	redirect,
	RouterProvider,
} from "react-router-dom";
import { AuthContext } from "./context/AuthContext";
import {
	createRequestInterceptor,
	createResponseInterceptor,
} from "./custom-instance";
import Layout from "./layout";
import Auth from "./pages/Auth";
import Follows from "./pages/Follows";
import Home from "./pages/Home";
import Post from "./pages/Post";
import User from "./pages/User";
import Search from "./pages/Search";
import Settings from "./pages/Settings";
import Messages from "./pages/Messages";

const App = () => {
	const { getAuth, setAuth, clearAuth, isAuthenticated } =
		useContext(AuthContext);

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
					path: "/:handle",
					element: <User value={"posts"} />,
				},
				{
					path: "/:handle/likes",
					element: <User value={"likes"} />,
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
				{
					path: "/search/posts",
					element: <Search value={"posts"} />,
				},
				{
					path: "/search/people",
					element: <Search value={"people"} />,
				},
				{
					loader: () => {
						if (!isAuthenticated) {
							return redirect("/");
						}
						return null;
					},
					children: [
						{
							path: "/home",
							element: <Home />,
						},
						{
							path: "/messages",
							element: <Messages />,
						},
						{
							path: "/messages/:handle",
							element: <Messages />,
						},
						{
							path: "/settings",
							element: <Settings />,
						},
					],
				},
			],
		},
	]);

	return <RouterProvider router={router} />;
};

export default App;
