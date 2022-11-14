import { createContext } from "react";
import { User } from "../models";

type Users = { [key: string]: Partial<User> };

interface UsersContextType {
	users: Users;
	setUser: (_handle: string, _user: Partial<User>) => void;
}

export const initialUsersContext: UsersContextType = {
	users: {},
	setUser(handle, user) {
		this.users = {
			...this.users,
			[handle]: { ...this.users[handle], ...user },
		};
	},
};

export default createContext<UsersContextType | null>(null);
