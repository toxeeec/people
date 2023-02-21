import { Modal } from "@mantine/core";
import { useForm } from "@mantine/form";
import { Dispatch, SetStateAction, useContext, useState } from "react";
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
					? "Username must have at least 5 characters"
					: value.length > 15
					? "Username cannot have more than 15 characters"
					: null,
			password: (value) =>
				value.length < 8 ? "Password must have at least 8 characters" : null,
		},
	});
	const { setAuth, setIsNewAccount } = useContext(AuthContext);
	const { mutate, isLoading } = usePostRegister();
	const handleLogin = () => {
		setSignupOpened(false);
		setLoginOpened(true);
	};
	const [imagePickerOpened, setImagePickerOpened] = useState(false);

	return (
		<>
			<AuthModal
				title="Sign up"
				opened={signupOpened}
				setOpened={setSignupOpened}
				isLoading={isLoading}
				form={form}
				handleSubmit={handleSubmit(mutate, setLoginOpened, setAuth, form, () =>
					setIsNewAccount(true)
				)}
				text="Already have an account? "
				handleChange={handleLogin}
				buttonText="Log in"
			/>
			<Modal
				opened={imagePickerOpened}
				onClose={() => setImagePickerOpened(false)}
			/>
		</>
	);
};
