// Package people provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package people

import (
	"database/sql"
	"time"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// AuthUser defines model for AuthUser.
type AuthUser struct {
	Handle   string   `db:"handle" fake:"{lettern:10}" json:"handle"`
	Hash     *string  `db:"hash" json:"hash,omitempty"`
	ID       *uint    `db:"user_id" fake:"skip" json:"id,omitempty"`
	Password Password `fake:"{password:true,true,true,true,false,12}" json:"password"`
}

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
}

// Post defines model for Post.
type Post struct {
	Content   string         `db:"content" fake:"{sentence}" json:"content"`
	CreatedAt time.Time      `db:"created_at" fake:"skip" json:"createdAt"`
	ID        uint           `db:"post_id" fake:"skip" json:"id"`
	Replies   uint           `db:"replies" fake:"skip" json:"replies"`
	RepliesTo *sql.NullInt32 `db:"replies_to" fake:"skip" json:"repliesTo,omitempty"`
	User      *User          `db:"user" json:"user,omitempty"`
}

// Posts defines model for Posts.
type Posts = SeekPaginationResult[Post]

// SeekPaginationMeta defines model for SeekPaginationMeta.
type SeekPaginationMeta struct {
	NewestID uint `json:"newestID"`
	OldestID uint `json:"oldestID"`
}

// Tokens defines model for Tokens.
type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// User defines model for User.
type User struct {
	Followers uint   `db:"followers" fake:"skip" json:"followers"`
	Following uint   `db:"following" fake:"skip" json:"following"`
	Handle    string `db:"handle" json:"handle"`
	ID        *uint  `db:"user_id" fake:"skip" json:"id,omitempty"`
}

// Users defines model for Users.
type Users = []User

// AfterParam defines model for afterParam.
type AfterParam uint

// BeforeParam defines model for beforeParam.
type BeforeParam uint

// HandleParam defines model for handleParam.
type HandleParam = string

// LimitParam defines model for limitParam.
type LimitParam uint

// PageParam defines model for pageParam.
type PageParam uint

// PostIDParam defines model for postIDParam.
type PostIDParam uint

// BadRequest defines model for BadRequest.
type BadRequest = Error

// Forbidden defines model for Forbidden.
type Forbidden = Error

// NoContent defines model for NoContent.
type NoContent = Error

// NotFound defines model for NotFound.
type NotFound = Error

// Unauthorized defines model for Unauthorized.
type Unauthorized = Error

// AuthUserBody defines model for AuthUserBody.
type AuthUserBody = AuthUser

// PostBody defines model for PostBody.
type PostBody struct {
	Content string `fake:"{sentence}" json:"content"`
}

// TokensBody defines model for TokensBody.
type TokensBody struct {
	RefreshToken string `db:"refreshToken" json:"refreshToken"`
}

// GetMeFeedParams defines parameters for GetMeFeed.
type GetMeFeedParams struct {
	Limit  *LimitParam  `form:"limit,omitempty" json:"limit,omitempty"`
	Before *BeforeParam `form:"before,omitempty" json:"before,omitempty"`
	After  *AfterParam  `form:"after,omitempty" json:"after,omitempty"`
}

// GetMeFollowersParams defines parameters for GetMeFollowers.
type GetMeFollowersParams struct {
	Page  *PageParam  `form:"page,omitempty" json:"page,omitempty"`
	Limit *LimitParam `form:"limit,omitempty" json:"limit,omitempty"`
}

// GetMeFollowingParams defines parameters for GetMeFollowing.
type GetMeFollowingParams struct {
	Page  *PageParam  `form:"page,omitempty" json:"page,omitempty"`
	Limit *LimitParam `form:"limit,omitempty" json:"limit,omitempty"`
}

// GetUsersHandleFollowersParams defines parameters for GetUsersHandleFollowers.
type GetUsersHandleFollowersParams struct {
	Page  *PageParam  `form:"page,omitempty" json:"page,omitempty"`
	Limit *LimitParam `form:"limit,omitempty" json:"limit,omitempty"`
}

// GetUsersHandleFollowingParams defines parameters for GetUsersHandleFollowing.
type GetUsersHandleFollowingParams struct {
	Page  *PageParam  `form:"page,omitempty" json:"page,omitempty"`
	Limit *LimitParam `form:"limit,omitempty" json:"limit,omitempty"`
}

// GetUsersHandlePostsParams defines parameters for GetUsersHandlePosts.
type GetUsersHandlePostsParams struct {
	Limit  *LimitParam  `form:"limit,omitempty" json:"limit,omitempty"`
	Before *BeforeParam `form:"before,omitempty" json:"before,omitempty"`
	After  *AfterParam  `form:"after,omitempty" json:"after,omitempty"`
}

// PostLoginJSONRequestBody defines body for PostLogin for application/json ContentType.
type PostLoginJSONRequestBody = AuthUserBody

// PostPostsJSONRequestBody defines body for PostPosts for application/json ContentType.
type PostPostsJSONRequestBody PostBody

// PostPostsPostIDRepliesJSONRequestBody defines body for PostPostsPostIDReplies for application/json ContentType.
type PostPostsPostIDRepliesJSONRequestBody PostBody

// PostRefreshJSONRequestBody defines body for PostRefresh for application/json ContentType.
type PostRefreshJSONRequestBody TokensBody

// PostRegisterJSONRequestBody defines body for PostRegister for application/json ContentType.
type PostRegisterJSONRequestBody = AuthUserBody
