import { createContext, type Dispatch, type SetStateAction, useState, useEffect } from "react";

type RouteContextType = {
	routeName: string;
	setRouteName: Dispatch<SetStateAction<string>>;
};

export const RouteContext = createContext<RouteContextType>({} as RouteContextType);

export function RouteContextProvider({ children }: { children: React.ReactNode }) {
	const [routeName, setRouteName] = useState("");
	useEffect(() => {
		if (!routeName) return;
		document.title = `${routeName} | People`;
	}, [routeName]);
	return (
		<RouteContext.Provider value={{ routeName, setRouteName }}>{children}</RouteContext.Provider>
	);
}
