import { createContext, useState } from "react";
import { User } from "../models";

export type Users = { [handle: string]: User | undefined };
interface UsersContextType {
	users: Users;
	setUser: (user: User) => void;
}

export const UsersContext = createContext<UsersContextType>(null!);

interface UsersContextProviderProps {
	children: React.ReactNode;
	initialUsers: Users;
}

export const UsersContextProvider = ({
	children,
	initialUsers,
}: UsersContextProviderProps) => {
	const [users, setUsers] = useState(initialUsers);
	const setUser = (user: User) => {
		setUsers((users) => ({ ...users, [user.handle]: user }));
	};
	return (
		<UsersContext.Provider value={{ users, setUser }}>
			{children}
		</UsersContext.Provider>
	);
};
