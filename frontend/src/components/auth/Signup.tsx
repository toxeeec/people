import { Modal } from "@mantine/core";
import { useForm } from "@mantine/form";
import { type Dispatch, type SetStateAction, useContext, useState } from "react";
import { AuthContext } from "@/context/AuthContext";
import { type AuthUser } from "@/models";
import { usePostRegister } from "@/spec.gen";
import { AuthModal } from "@/components/auth/AuthModal";

type SignupProps = {
	signupOpened: boolean;
	setSignupOpened: Dispatch<SetStateAction<boolean>>;
	setLoginOpened: Dispatch<SetStateAction<boolean>>;
};

export function Signup({ signupOpened, setSignupOpened, setLoginOpened }: SignupProps) {
	const form = useForm<AuthUser>({
		initialValues: { handle: "", password: "" },
		validate: {
			handle: (value) => {
				if (value.length < 5) return "Username must contain at least 5 characters";
				if (value.length > 15) return "Username cannot contain more than 15 characters";
			},
			password: (value) => {
				if (value.length < 8) return "Password must contain at least 8 characters";
			},
		},
	});
	const { setIsNewAccount } = useContext(AuthContext);
	const [imagePickerOpened, setImagePickerOpened] = useState(false);
	const { mutate, isLoading } = usePostRegister();
	const handleLogin = () => {
		setSignupOpened(false);
		setLoginOpened(true);
	};

	return (
		<>
			<AuthModal
				title="Sign up"
				opened={signupOpened}
				setOpened={setSignupOpened}
				isLoading={isLoading}
				form={form}
				text="Already have an account?"
				handleChange={handleLogin}
				buttonText="Log in"
				mutate={mutate}
				onSuccess={() => setIsNewAccount(true)}
			/>
			<Modal opened={imagePickerOpened} onClose={() => setImagePickerOpened(false)} />
		</>
	);
}
