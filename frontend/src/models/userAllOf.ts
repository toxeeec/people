/**
 * Generated by orval v6.10.3 🍺
 * Do not edit manually.
 * People API
 * OpenAPI spec version: 1.0.0
 */
import type { FollowStatus } from "@/models/followStatus";

export type UserAllOf = {
	handle: string;
	following: number;
	followers: number;
	status?: FollowStatus;
};
