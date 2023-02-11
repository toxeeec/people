// Package people provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package people

import (
	"time"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// AuthResponse defines model for AuthResponse.
type AuthResponse struct {
	Tokens Tokens `json:"tokens"`
	User   User   `json:"user"`
}

// AuthUser defines model for AuthUser.
type AuthUser struct {
	Handle   string `db:"handle" fake:"{lettern:10}" json:"handle"`
	Password string `fake:"{password:true,true,true,true,false,8}" json:"password"`
}

// Error defines model for Error.
type Error struct {
	Kind    *ErrorKind `json:"kind,omitempty"`
	Message string     `json:"message"`
}

// FollowStatus defines model for FollowStatus.
type FollowStatus struct {
	IsFollowed  bool `db:"is_followed" json:"isFollowed"`
	IsFollowing bool `db:"is_following" json:"isFollowing"`
}

// Handle defines model for Handle.
type Handle struct {
	Handle string `db:"handle" fake:"{lettern:10}" json:"handle"`
}

// HandlePaginationMeta defines model for HandlePaginationMeta.
type HandlePaginationMeta = PaginationMeta[string]

// IDPaginationMeta defines model for IDPaginationMeta.
type IDPaginationMeta = PaginationMeta[uint]

// ImageResponse defines model for ImageResponse.
type ImageResponse struct {
	ID uint `db:"image_id" fake:"skip" json:"id"`
}

// LikeStatus defines model for LikeStatus.
type LikeStatus struct {
	IsLiked bool `db:"is_liked" json:"isLiked"`
}

// NewImage defines model for NewImage.
type NewImage struct {
	Image openapi_types.File `json:"image"`
}

// NewPost defines model for NewPost.
type NewPost struct {
	Content string  `fake:"{sentence}" json:"content"`
	Images  *[]uint `fake:"skip" json:"images,omitempty"`
}

// Post defines model for Post.
type Post struct {
	Content   string      `db:"content" fake:"{sentence}" json:"content"`
	CreatedAt time.Time   `db:"created_at" fake:"skip" json:"createdAt"`
	ID        uint        `db:"post_id" fake:"skip" json:"id"`
	Images    *[]string   `fake:"skip" json:"images,omitempty"`
	Likes     uint        `db:"likes" fake:"skip" json:"likes"`
	Replies   uint        `db:"replies" fake:"skip" json:"replies"`
	RepliesTo *uint       `db:"replies_to" fake:"skip" json:"repliesTo,omitempty"`
	Status    *LikeStatus `json:"status,omitempty"`
	UserID    uint        `db:"user_id" json:"-"`
}

// PostResponse defines model for PostResponse.
type PostResponse struct {
	Data Post `json:"data"`
	User User `json:"user"`
}

// PostsResponse defines model for PostsResponse.
type PostsResponse = PaginatedResults[PostResponse, uint]

// Tokens defines model for Tokens.
type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// User defines model for User.
type User struct {
	Followers uint          `db:"followers" fake:"skip" json:"followers"`
	Following uint          `db:"following" fake:"skip" json:"following"`
	Handle    string        `db:"handle" json:"handle"`
	ID        uint          `db:"user_id" json:"-"`
	Status    *FollowStatus `json:"status,omitempty"`
}

// Users defines model for Users.
type Users = PaginatedResults[User, string]

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

// QueryParam defines model for queryParam.
type QueryParam = string

// AuthUserBody defines model for AuthUserBody.
type AuthUserBody = AuthUser

// LogoutBody defines model for LogoutBody.
type LogoutBody struct {
	LogoutFromAll *bool  `json:"logoutFromAll,omitempty"`
	RefreshToken  string `json:"refreshToken"`
}

// NewPostBody defines model for NewPostBody.
type NewPostBody = NewPost

// RefreshTokenBody defines model for RefreshTokenBody.
type RefreshTokenBody struct {
	RefreshToken string `json:"refreshToken"`
}

// PostLogoutJSONBody defines parameters for PostLogout.
type PostLogoutJSONBody struct {
	LogoutFromAll *bool  `json:"logoutFromAll,omitempty"`
	RefreshToken  string `json:"refreshToken"`
}

// DeleteMeJSONBody defines parameters for DeleteMe.
type DeleteMeJSONBody struct {
	Password string `json:"password"`
}

// PutMeJSONBody defines parameters for PutMe.
type PutMeJSONBody = Handle

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

// GetPostsSearchParams defines parameters for GetPostsSearch.
type GetPostsSearchParams struct {
	Query  QueryParam   `form:"query" json:"query"`
	Limit  *LimitParam  `form:"limit,omitempty" json:"limit,omitempty"`
	Before *BeforeParam `form:"before,omitempty" json:"before,omitempty"`
	After  *AfterParam  `form:"after,omitempty" json:"after,omitempty"`
}

// GetPostsPostIDLikesParams defines parameters for GetPostsPostIDLikes.
type GetPostsPostIDLikesParams struct {
	Limit  *LimitParam        `form:"limit,omitempty" json:"limit,omitempty"`
	Before *BeforeHandleParam `form:"before,omitempty" json:"before,omitempty"`
	After  *AfterHandleParam  `form:"after,omitempty" json:"after,omitempty"`
}

// GetPostsPostIDRepliesParams defines parameters for GetPostsPostIDReplies.
type GetPostsPostIDRepliesParams struct {
	Limit  *LimitParam  `form:"limit,omitempty" json:"limit,omitempty"`
	Before *BeforeParam `form:"before,omitempty" json:"before,omitempty"`
	After  *AfterParam  `form:"after,omitempty" json:"after,omitempty"`
}

// PostRefreshJSONBody defines parameters for PostRefresh.
type PostRefreshJSONBody struct {
	RefreshToken string `json:"refreshToken"`
}

// GetUsersSearchParams defines parameters for GetUsersSearch.
type GetUsersSearchParams struct {
	Query  QueryParam         `form:"query" json:"query"`
	Limit  *LimitParam        `form:"limit,omitempty" json:"limit,omitempty"`
	Before *BeforeHandleParam `form:"before,omitempty" json:"before,omitempty"`
	After  *AfterHandleParam  `form:"after,omitempty" json:"after,omitempty"`
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

// GetUsersHandleLikesParams defines parameters for GetUsersHandleLikes.
type GetUsersHandleLikesParams struct {
	Limit  *LimitParam  `form:"limit,omitempty" json:"limit,omitempty"`
	Before *BeforeParam `form:"before,omitempty" json:"before,omitempty"`
	After  *AfterParam  `form:"after,omitempty" json:"after,omitempty"`
}

// GetUsersHandlePostsParams defines parameters for GetUsersHandlePosts.
type GetUsersHandlePostsParams struct {
	Limit  *LimitParam  `form:"limit,omitempty" json:"limit,omitempty"`
	Before *BeforeParam `form:"before,omitempty" json:"before,omitempty"`
	After  *AfterParam  `form:"after,omitempty" json:"after,omitempty"`
}

// PostImagesMultipartRequestBody defines body for PostImages for multipart/form-data ContentType.
type PostImagesMultipartRequestBody = NewImage

// PostLoginJSONRequestBody defines body for PostLogin for application/json ContentType.
type PostLoginJSONRequestBody = AuthUser

// PostLogoutJSONRequestBody defines body for PostLogout for application/json ContentType.
type PostLogoutJSONRequestBody PostLogoutJSONBody

// DeleteMeJSONRequestBody defines body for DeleteMe for application/json ContentType.
type DeleteMeJSONRequestBody DeleteMeJSONBody

// PutMeJSONRequestBody defines body for PutMe for application/json ContentType.
type PutMeJSONRequestBody = PutMeJSONBody

// PostPostsJSONRequestBody defines body for PostPosts for application/json ContentType.
type PostPostsJSONRequestBody = NewPost

// PostPostsPostIDRepliesJSONRequestBody defines body for PostPostsPostIDReplies for application/json ContentType.
type PostPostsPostIDRepliesJSONRequestBody = NewPost

// PostRefreshJSONRequestBody defines body for PostRefresh for application/json ContentType.
type PostRefreshJSONRequestBody PostRefreshJSONBody

// PostRegisterJSONRequestBody defines body for PostRegister for application/json ContentType.
type PostRegisterJSONRequestBody = AuthUser
