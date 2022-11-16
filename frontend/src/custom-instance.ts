import Axios, { AxiosError, AxiosRequestConfig } from "axios";
import { AuthValues, SetAuthProps } from "./context/AuthContext";
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

export function createRequestInterceptor(getAuth: () => AuthValues) {
	return AXIOS_INSTANCE.interceptors.request.use((config) => {
		console.log(config.url);
		const { accessToken } = getAuth();
		if (accessToken) {
			config.headers = {
				...config.headers,
				Authorization: `Bearer ${accessToken}`,
			};
		}
		return config;
	});
}

export function createResponseInterceptor(
	getAuth: () => AuthValues,
	setAuth: (_props: SetAuthProps) => void,
	clearAuth: () => void
) {
	return AXIOS_INSTANCE.interceptors.response.use(
		(res) => {
			return res;
		},
		(err: AxiosError) => {
			if (err.response?.status === 403 && err.config?.url !== "/refresh") {
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
			return Promise.reject(err);
		}
	);
}

export type ErrorType<Error> = AxiosError<Error>;
