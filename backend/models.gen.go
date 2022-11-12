// Package people provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.2 DO NOT EDIT.
package people

import (
	"database/sql"
	"time"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// AuthResponse defines model for AuthResponse.
type AuthResponse struct {
	Tokens Tokens `json:"tokens"`
	User   User   `db:"user" json:"user"`
}

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

// HandlePaginationMeta defines model for HandlePaginationMeta.
type HandlePaginationMeta = PaginationMeta[string]

// IDPaginationMeta defines model for IDPaginationMeta.
type IDPaginationMeta = PaginationMeta[uint]

// Likes defines model for Likes.
type Likes struct {
	IsLiked bool `db:"is_liked" json:"isLiked"`
	Likes   uint `db:"likes" fake:"skip" json:"likes"`
}

// Post defines model for Post.
type Post struct {
	Content   string         `db:"content" fake:"{sentence}" json:"content"`
	CreatedAt time.Time      `db:"created_at" fake:"skip" json:"createdAt"`
	ID        uint           `db:"post_id" fake:"skip" json:"id"`
	IsLiked   bool           `db:"is_liked" json:"isLiked"`
	Likes     uint           `db:"likes" fake:"skip" json:"likes"`
	Replies   uint           `db:"replies" fake:"skip" json:"replies"`
	RepliesTo *sql.NullInt32 `db:"replies_to" fake:"skip" json:"repliesTo,omitempty"`
	User      *User          `db:"user" json:"user,omitempty"`
}

// Posts defines model for Posts.
type Posts = PaginationResult[Post, uint]

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
}

// Users defines model for Users.
type Users = PaginationResult[User, string]

// AfterHandleParam defines model for afterHandleParam.
type AfterHandleParam = string

// AfterParam defines model for afterParam.
type AfterParam = uint

// BeforeHandleParam defines model for beforeHandleParam.
type BeforeHandleParam = string

// BeforeParam defines model for beforeParam.
type BeforeParam = uint

// HandleParam defines model for handleParam.
type HandleParam = string

// LimitParam defines model for limitParam.
type LimitParam = uint

// PostIDParam defines model for postIDParam.
type PostIDParam = uint

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
	Limit  *LimitParam        `form:"limit,omitempty" json:"limit,omitempty"`
	Before *BeforeHandleParam `form:"before,omitempty" json:"before,omitempty"`
	After  *AfterHandleParam  `form:"after,omitempty" json:"after,omitempty"`
}

// GetMeFollowingParams defines parameters for GetMeFollowing.
type GetMeFollowingParams struct {
	Limit  *LimitParam        `form:"limit,omitempty" json:"limit,omitempty"`
	Before *BeforeHandleParam `form:"before,omitempty" json:"before,omitempty"`
	After  *AfterHandleParam  `form:"after,omitempty" json:"after,omitempty"`
}

// PostPostsJSONBody defines parameters for PostPosts.
type PostPostsJSONBody struct {
	Content string `fake:"{sentence}" json:"content"`
}

// GetPostsPostIDRepliesParams defines parameters for GetPostsPostIDReplies.
type GetPostsPostIDRepliesParams struct {
	Limit  *LimitParam  `form:"limit,omitempty" json:"limit,omitempty"`
	Before *BeforeParam `form:"before,omitempty" json:"before,omitempty"`
	After  *AfterParam  `form:"after,omitempty" json:"after,omitempty"`
}

// PostPostsPostIDRepliesJSONBody defines parameters for PostPostsPostIDReplies.
type PostPostsPostIDRepliesJSONBody struct {
	Content string `fake:"{sentence}" json:"content"`
}

// PostRefreshJSONBody defines parameters for PostRefresh.
type PostRefreshJSONBody struct {
	RefreshToken string `db:"refreshToken" json:"refreshToken"`
}

// GetUsersHandleFollowersParams defines parameters for GetUsersHandleFollowers.
type GetUsersHandleFollowersParams struct {
	Limit  *LimitParam        `form:"limit,omitempty" json:"limit,omitempty"`
	Before *BeforeHandleParam `form:"before,omitempty" json:"before,omitempty"`
	After  *AfterHandleParam  `form:"after,omitempty" json:"after,omitempty"`
}

// GetUsersHandleFollowingParams defines parameters for GetUsersHandleFollowing.
type GetUsersHandleFollowingParams struct {
	Limit  *LimitParam        `form:"limit,omitempty" json:"limit,omitempty"`
	Before *BeforeHandleParam `form:"before,omitempty" json:"before,omitempty"`
	After  *AfterHandleParam  `form:"after,omitempty" json:"after,omitempty"`
}

// GetUsersHandlePostsParams defines parameters for GetUsersHandlePosts.
type GetUsersHandlePostsParams struct {
	Limit  *LimitParam  `form:"limit,omitempty" json:"limit,omitempty"`
	Before *BeforeParam `form:"before,omitempty" json:"before,omitempty"`
	After  *AfterParam  `form:"after,omitempty" json:"after,omitempty"`
}

// PostLoginJSONRequestBody defines body for PostLogin for application/json ContentType.
type PostLoginJSONRequestBody = AuthUser

// PostPostsJSONRequestBody defines body for PostPosts for application/json ContentType.
type PostPostsJSONRequestBody PostPostsJSONBody

// PostPostsPostIDRepliesJSONRequestBody defines body for PostPostsPostIDReplies for application/json ContentType.
type PostPostsPostIDRepliesJSONRequestBody PostPostsPostIDRepliesJSONBody

// PostRefreshJSONRequestBody defines body for PostRefresh for application/json ContentType.
type PostRefreshJSONRequestBody PostRefreshJSONBody

// PostRegisterJSONRequestBody defines body for PostRegister for application/json ContentType.
type PostRegisterJSONRequestBody = AuthUser
