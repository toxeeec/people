/**
 * Generated by orval v6.10.3 🍺
 * Do not edit manually.
 * People API
 * OpenAPI spec version: 1.0.0
 */
import type { LikeStatus } from "@/models/likeStatus";

export type PostInner = {
	id: number;
	content: string;
	likes: number;
	replies: number;
	createdAt: string;
	repliesTo?: number;
	status?: LikeStatus;
};
