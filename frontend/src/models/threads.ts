/**
 * Generated by orval v6.10.3 🍺
 * Do not edit manually.
 * People API
 * OpenAPI spec version: 1.0.0
 */
import type { Thread } from "./thread";
import type { IDPaginationMeta } from "./iDPaginationMeta";

export interface Threads {
	data: Thread[];
	meta?: IDPaginationMeta;
}