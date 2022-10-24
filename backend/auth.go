package people

import (
	"github.com/google/uuid"
)

type RefreshToken struct {
	ID     uuid.UUID `db:"token_id"`
	Value  string    `db:"value"`
	UserID uint      `db:"user_id"`
}

type AuthService interface {
	VerifyCredentials(AuthUser) (uint, error)
	NewTokens(id uint) (Tokens, error)
	UpdateRefreshToken(userID uint, tokenID uuid.UUID) (RefreshToken, error)
	CheckRefreshToken(RefreshToken) bool
}
