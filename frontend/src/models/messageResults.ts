/**
 * Generated by orval v6.10.3 🍺
 * Do not edit manually.
 * People API
 * OpenAPI spec version: 1.0.0
 */
import type { ServerMessage } from "./serverMessage";
import type { IDPaginationMeta } from "./iDPaginationMeta";

export interface MessageResults {
	data: ServerMessage[];
	meta?: IDPaginationMeta;
}
