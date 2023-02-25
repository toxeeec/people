import { createContext, type Dispatch, type SetStateAction, useState } from "react";

type RouteContextType = {
	routeName: string;
	setRouteName: Dispatch<SetStateAction<string>>;
};

export const RouteContext = createContext<RouteContextType>({} as RouteContextType);

export function RouteContextProvider({ children }: { children: React.ReactNode }) {
	const [routeName, setRouteName] = useState("");
	return (
		<RouteContext.Provider value={{ routeName, setRouteName }}>{children}</RouteContext.Provider>
	);
}
