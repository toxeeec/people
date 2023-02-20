import { FormErrors, UseFormReturnType } from "@mantine/form";
import { UseMutateFunction } from "@tanstack/react-query";
import { Dispatch, SetStateAction } from "react";
import { SetAuthProps } from "../../context/AuthContext";
import { ErrorType } from "../../custom-instance";
import { AuthResponse, AuthUser, Error } from "../../models";

export const handleSubmit = (
	mutate: UseMutateFunction<
		AuthResponse,
		ErrorType<Error>,
		{
			data: AuthUser;
		},
		unknown
	>,
	setOpened: Dispatch<SetStateAction<boolean>>,
	setAuth: (props: SetAuthProps) => void,
	form: UseFormReturnType<AuthUser, (values: AuthUser) => AuthUser>
) => {
	return (user: AuthUser) => {
		mutate(
			{ data: user },
			{
				onSuccess: (res) => {
					setAuth({ tokens: res.tokens, handle: res.user.handle });
					setOpened(false);
				},
				onError: (error) => {
					const err = error.response?.data.message
						.replaceAll("Handle", "Username")
						.replaceAll("handle", "username");
					const errors: SetStateAction<FormErrors> = {};
					if (err?.toLowerCase().includes("username")) {
						errors.handle = err;
					}
					if (err?.toLowerCase().includes("password")) {
						errors.password = err;
					}
					form.setErrors(errors);
				},
			}
		);
	};
};
