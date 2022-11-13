import Axios, { AxiosError, AxiosRequestConfig } from "axios";
import { SetAuthProps } from "./context/AuthContext";
import { postRefresh } from "./spec.gen";

export const AXIOS_INSTANCE = Axios.create({
	baseURL: "http://localhost:8000",
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

export function createRequestInterceptor(accessToken: string) {
	return AXIOS_INSTANCE.interceptors.request.use((config) => {
		config.headers = {
			...config.headers,
			Authorization: `Bearer ${accessToken}`,
		};
		return config;
	});
}
export function createResponseInterceptor(
	refreshToken: string,
	setAuth: (_props: SetAuthProps) => void,
	clearAuth: () => void
) {
	return AXIOS_INSTANCE.interceptors.response.use(
		(res) => res,
		(err: AxiosError) => {
			if (err.response?.status === 403 && err.config?.url !== "/refresh") {
				postRefresh({ refreshToken: refreshToken })
					.then((tokens) => {
						setAuth({ tokens });
					})
					.catch(() => {
						clearAuth();
					});
			}
			return Promise.reject(err);
		}
	);
}

export type ErrorType<Error> = AxiosError<Error>;
