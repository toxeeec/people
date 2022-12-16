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
			handle: (value) => (value.length < 5 ? "Invalid Handle" : null),
			password: (value) => (value.length < 12 ? "Invalid Password" : null),
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
			setOpened={setLoginOpened}
			isLoading={isLoading}
			form={form}
			handleSubmit={handleSubmit(mutate, setLoginOpened, setAuth, form)}
			text="Already have an account? "
			handleChange={handleLogin}
			buttonText="Log in"
		/>
	);
};
