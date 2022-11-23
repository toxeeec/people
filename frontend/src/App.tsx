import { useQueryClient } from "@tanstack/react-query";
import { useContext, useEffect } from "react";
import {
	createBrowserRouter,
	redirect,
	RouterProvider,
} from "react-router-dom";
import AuthContext from "./context/AuthContext";
import UsersContext from "./context/UsersContext";
import {
	createRequestInterceptor,
	createResponseInterceptor,
} from "./custom-instance";
import useAuth from "./hooks/useAuth";
import Layout from "./layout";
import Auth from "./pages/Auth";
import Follows, { FollowsPage } from "./pages/Follows";
import Home from "./pages/Home";
import MainPost from "./pages/MainPost";
import Profile from "./pages/Profile";
import { getPostsPostID, getUsersHandle } from "./spec.gen";

export default function App() {
	const usersCtx = useContext(UsersContext);
	const { getAuth, setAuth, clearAuth, isAuthenticated } = useAuth();
	const queryClient = useQueryClient();

	useEffect(() => {
		if (isAuthenticated) {
			getUsersHandle(getAuth().handle!).then((user) =>
				usersCtx?.setUser(user!.handle!, user!)
			);
		}
	}, [isAuthenticated, getAuth, usersCtx]);

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
					},
				},

				{
					path: "/:handle",
					element: <Profile />,
					loader: async ({ params }) => {
						return queryClient.fetchQuery({
							queryKey: ["user", params.handle],
							queryFn: () =>
								getUsersHandle(params.handle!).then((u) => {
									usersCtx?.setUser(u.handle, u);
									return u;
								}),
						});
					},
				},

				{
					path: "/:handle/:postID",
					element: <MainPost />,
					loader: async ({ params }) => {
						return queryClient.fetchQuery({
							queryKey: ["post", params.postID],
							queryFn: () => getPostsPostID(parseInt(params.postID!)),
						});
					},
				},
				{
					path: "/:handle/following",
					element: <Follows value={FollowsPage.Following} />,
				},
				{
					path: "/:handle/followers",
					element: <Follows value={FollowsPage.Followers} />,
				},
			],
		},
	]);

	return (
		<AuthContext.Provider
			value={{ getAuth, setAuth, clearAuth, isAuthenticated }}
		>
			<RouterProvider router={router} />
		</AuthContext.Provider>
	);
}
