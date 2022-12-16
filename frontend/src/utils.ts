import { MouseEvent } from "react";

export const stopPropagation = (e: MouseEvent) => {
	e.stopPropagation();
};
