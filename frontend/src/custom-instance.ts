import Axios, { AxiosError, AxiosRequestConfig } from "axios";
import { AuthValues, SetAuthProps } from "./context/AuthContext";
import { postRefresh } from "./spec.gen";

const port = BACKEND_PORT;
export const baseURL = `${location.protocol}//${location.hostname}:${port}`;
export const wsURL = `ws://${location.hostname}:${port}/ws`;
export const AXIOS_INSTANCE = Axios.create({
	baseURL: `${baseURL}/api`,
});

export const customInstance = <T>(
	config: AxiosRequestConfig,
	options?: AxiosRequestConfig
): Promise<T> => {
	const source = Axios.CancelToken.source();
	const promise = AXIOS_INSTANCE({
		...config,
		...options,
		cancelToken: source.token,
	}).then(({ data }) => data);

	// eslint-disable-next-line
	(promise as any).cancel = () => {
		source.cancel("Query was cancelled");
	};

	return promise;
};

export const createRequestInterceptor = (getAuth: () => AuthValues) => {
	return AXIOS_INSTANCE.interceptors.request.use((config) => {
		const { accessToken } = getAuth();
		if (accessToken) {
			config.headers = {
				Authorization: `Bearer ${accessToken}`,
			};
		}
		return config;
	});
};

export const createResponseInterceptor = (
	getAuth: () => AuthValues,
	setAuth: (props: SetAuthProps) => void,
	clearAuth: () => void
) => {
	return AXIOS_INSTANCE.interceptors.response.use(
		(res) => {
			return res;
		},
		(error: AxiosError) => {
			if (error.response?.status === 403 && error.config?.url !== "/refresh") {
				const err = (error as ErrorType<Error>).response?.data.message;
				if (!err?.includes("token")) return;
				const { refreshToken } = getAuth();
				if (refreshToken) {
					postRefresh({ refreshToken: refreshToken })
						.then((tokens) => {
							setAuth({ tokens });
						})
						.catch(() => {
							clearAuth();
						});
				} else {
					clearAuth();
				}
			}
			return Promise.reject(error);
		}
	);
};

export type ErrorType<Error> = AxiosError<Error>;
