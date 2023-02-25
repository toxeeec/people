import { AuthModal } from "@/components/auth/AuthModal";
import { type AuthUser } from "@/models";
import { usePostLogin } from "@/spec.gen";
import { useForm } from "@mantine/form";
import { type Dispatch, type SetStateAction } from "react";

type LoginProps = {
	loginOpened: boolean;
	setLoginOpened: Dispatch<SetStateAction<boolean>>;
	setSignupOpened: Dispatch<SetStateAction<boolean>>;
};

export function Login({ loginOpened, setLoginOpened, setSignupOpened }: LoginProps) {
	const form = useForm<AuthUser>({
		initialValues: { handle: "", password: "" },
		validate: {
			handle: (value) => (value.length < 5 || value.length > 15) && "Invalid Username",
			password: (value) => value.length < 8 && "Invalid Password",
		},
	});
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
			text="Don't have an account?"
			handleChange={handleSignup}
			buttonText="Sign up"
			mutate={mutate}
		/>
	);
}
