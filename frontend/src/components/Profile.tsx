import { Paper } from "@mantine/core";
import {
	Dispatch,
	forwardRef,
	SetStateAction,
	useContext,
	useEffect,
	useState,
} from "react";
import AuthContext from "../context/AuthContext";
import UsersContext from "../context/UsersContext";
import { User } from "../models";
import AccountInfo from "./AccountInfo";
import FollowButton from "./FollowButton";

interface ProfileProps {
	user: Partial<User>;
	setUser?: Dispatch<SetStateAction<Partial<User>>>;
}

const Profile = forwardRef<HTMLDivElement, ProfileProps>(
	({ user, setUser }, ref) => {
		const usersCtx = useContext(UsersContext)!;
		const { getAuth } = useContext(AuthContext)!;
		const handle = getAuth().handle;
		const [childUser, setChildUser] = useState(user);

		useEffect(() => {
			setUser && setUser(childUser);
		}, [childUser, setUser, usersCtx]);

		const updateUser = (u: Partial<User>) => {
			setChildUser(usersCtx.setUser(user.handle!, u));
		};
		return (
			<Paper ref={ref} p="md">
				<AccountInfo user={childUser!}>
					{handle === childUser.handle ? null : (
						<FollowButton user={childUser} updateUser={updateUser} />
					)}
				</AccountInfo>
			</Paper>
		);
	}
);

Profile.displayName = "Profile";
export default Profile;
