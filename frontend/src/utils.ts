export const uniqueArray = <T>(arr: T[]) => {
	return [...new Set(arr)];
};
