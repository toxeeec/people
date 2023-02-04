import { useForm } from "@mantine/form";
import { Dispatch, SetStateAction, useContext } from "react";
import { AuthContext } from "../../context/AuthContext";
import { AuthUser } from "../../models";
import { usePostRegister } from "../../spec.gen";
import { AuthModal } from "./AuthModal";
import { handleSubmit } from "./handleSubmit";

interface SignupProps {
	signupOpened: boolean;
	setSignupOpened: Dispatch<SetStateAction<boolean>>;
	setLoginOpened: Dispatch<SetStateAction<boolean>>;
}

export const Signup = ({
	signupOpened,
	setSignupOpened,
	setLoginOpened,
}: SignupProps) => {
	const form = useForm<AuthUser>({
		initialValues: { handle: "", password: "" },
		validate: {
			handle: (value) =>
				value.length < 5
					? "Handle must have at least 5 characters"
					: value.length > 15
					? "Handle cannot have more than 15 characters"
					: null,
			password: (value) =>
				value.length < 12 ? "Password must have at least 12 characters" : null,
		},
	});
	const { setAuth } = useContext(AuthContext);
	const { mutate, isLoading } = usePostRegister();
	const handleLogin = () => {
		setSignupOpened(false);
		setLoginOpened(true);
	};

	return (
		<AuthModal
			title="Sign up"
			opened={signupOpened}
			setOpened={setSignupOpened}
			isLoading={isLoading}
			form={form}
			handleSubmit={handleSubmit(mutate, setLoginOpened, setAuth, form)}
			text="Already have an account? "
			handleChange={handleLogin}
			buttonText="Log in"
		/>
	);
};
