import { MouseEvent } from "react";

export function stopPropagation(e: MouseEvent) {
	e.stopPropagation();
}
