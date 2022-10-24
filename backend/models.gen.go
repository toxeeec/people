// Package people provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package people

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// AuthUser defines model for AuthUser.
type AuthUser struct {
	Handle   string   `fake:"{lettern:10}" json:"handle"`
	Password Password `fake:"{password:true,true,true,true,false,12}" json:"password"`
}

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
}

// Tokens defines model for Tokens.
type Tokens struct {
	AccessToken  *string `json:"accessToken,omitempty"`
	RefreshToken string  `json:"refreshToken"`
}

// User defines model for User.
type User struct {
	Followers uint   `db:"followers" fake:"skip" json:"followers"`
	Following uint   `db:"following" fake:"skip" json:"following"`
	Handle    string `db:"handle" json:"handle"`
	Hash      string `db:"hash" json:"-"`
	ID        uint   `db:"user_id" fake:"skip" json:"-"`
	Password  string `json:"-"`
}

// HandleParam defines model for handleParam.
type HandleParam = string

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

// TokensBody defines model for TokensBody.
type TokensBody = Tokens

// PostLoginJSONRequestBody defines body for PostLogin for application/json ContentType.
type PostLoginJSONRequestBody = AuthUserBody

// PostRefreshJSONRequestBody defines body for PostRefresh for application/json ContentType.
type PostRefreshJSONRequestBody = TokensBody

// PostRegisterJSONRequestBody defines body for PostRegister for application/json ContentType.
type PostRegisterJSONRequestBody = AuthUserBody
