import { AuthContext } from "@/context/AuthContext";
import { NotificationsContextProvider } from "@/context/NotificationsContext";
import { RouteContextProvider } from "@/context/RouteContext";
import {
	AXIOS_INSTANCE,
	createRequestInterceptor,
	createResponseInterceptor,
} from "@/custom-instance";
import { Layout } from "@/layout";
import Auth from "@/pages/Auth";
import Follows from "@/pages/Follows";
import Home from "@/pages/Home";
import Messages from "@/pages/Messages";
import NotFound from "@/pages/NotFound";
import Post from "@/pages/Post";
import Search from "@/pages/Search";
import Settings from "@/pages/Settings";
import User from "@/pages/User";
import { MantineProvider } from "@mantine/core";
import { NotificationsProvider } from "@mantine/notifications";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { useContext, useEffect } from "react";
import { createBrowserRouter, redirect, RouterProvider } from "react-router-dom";
const queryClient = new QueryClient({ defaultOptions: { queries: { retry: 1 } } });

export default function App() {
	const { getAuth, setAuth, clearAuth, isAuthenticated } = useContext(AuthContext);

	useEffect(() => {
		const requestInterceptor = createRequestInterceptor(getAuth);
		const responseInterceptor = createResponseInterceptor(getAuth, setAuth, clearAuth);
		return () => {
			AXIOS_INSTANCE.interceptors.request.eject(requestInterceptor);
			AXIOS_INSTANCE.interceptors.response.eject(responseInterceptor);
		};
	}, [clearAuth, getAuth, setAuth]);

	const router = createBrowserRouter([
		{
			index: true,
			element: <Auth />,
			loader: () => isAuthenticated && redirect("/home"),
			errorElement: <NotFound />,
		},
		{
			element: <Layout />,
			children: [
				{
					path: "/:handle",
					element: <User value={"posts"} />,
				},
				{
					path: "/:handle/posts",
					element: <User value={"posts"} />,
				},
				{
					path: "/:handle/likes",
					element: <User value={"likes"} />,
				},
				{
					path: "/posts/:postID",
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
					loader: () => !isAuthenticated && redirect("/"),
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
							path: "/messages/:thread",
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
		{
			path: "/404",
			element: <NotFound />,
		},
	]);

	return (
		<MantineProvider withGlobalStyles withNormalizeCSS theme={{ colorScheme: "dark" }}>
			<NotificationsProvider position="bottom-center">
				<QueryClientProvider client={queryClient}>
					<RouteContextProvider>
						<NotificationsContextProvider>
							<RouterProvider router={router} />
						</NotificationsContextProvider>
					</RouteContextProvider>
				</QueryClientProvider>
			</NotificationsProvider>
		</MantineProvider>
	);
}
