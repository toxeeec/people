package people

type AuthService interface {
	NewTokens(id uint) (Tokens, error)
}
