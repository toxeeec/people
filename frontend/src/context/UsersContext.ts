import { createContext } from "react";
import { User } from "../models";

interface UsersContextType {
	users: Map<string, Partial<User>>;
	setUser: (_handle: string, _user: Partial<User>) => Partial<User>;
}

export const initialUsersContext: UsersContextType = {
	users: new Map(),
	setUser(handle, user) {
		this.users.set(handle, { ...this.users.get(handle), ...user });
		return this.users.get(handle)!;
	},
};

export default createContext<UsersContextType | null>(null);
