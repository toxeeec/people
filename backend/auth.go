package people

import (
	"github.com/google/uuid"
	"github.com/toxeeec/people/backend/token"
)

type AuthService interface {
	VerifyCredentials(AuthUser) (uint, error)
	NewTokens(id uint) (Tokens, error)
	UpdateRefreshToken(userID uint, tokenID uuid.UUID) (token.RefreshToken, error)
	CheckRefreshToken(token.RefreshToken) bool
}
