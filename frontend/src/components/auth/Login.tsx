import { useForm } from "@mantine/form";
import { Dispatch, SetStateAction, useContext } from "react";
import { AuthUser } from "../../models";
import { usePostLogin } from "../../spec.gen";
import { AuthContext } from "../../context/AuthContext";
import { handleSubmit } from "./handleSubmit";
import { AuthModal } from "./AuthModal";

interface LoginProps {
	loginOpened: boolean;
	setLoginOpened: Dispatch<SetStateAction<boolean>>;
	setSignupOpened: Dispatch<SetStateAction<boolean>>;
}

export const Login = ({
	loginOpened,
	setLoginOpened,
	setSignupOpened,
}: LoginProps) => {
	const form = useForm<AuthUser>({
		initialValues: { handle: "", password: "" },
		validate: {
			handle: (value) =>
				value.length < 5 || value.length > 15 ? "Invalid Handle" : null,
			password: (value) => (value.length < 12 ? "Invalid Password" : null),
		},
	});
	const { setAuth } = useContext(AuthContext);
	const { mutate, isLoading } = usePostLogin();
	const handleSignup = () => {
		setLoginOpened(false);
		setSignupOpened(true);
	};
	return (
		<AuthModal
			title="Log in"
			opened={loginOpened}
			setOpened={setLoginOpened}
			isLoading={isLoading}
			form={form}
			handleSubmit={handleSubmit(mutate, setLoginOpened, setAuth, form)}
			text="Don't have an account? "
			handleChange={handleSignup}
			buttonText="Sign up"
		/>
	);
};
