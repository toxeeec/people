package people

type UserService interface {
	Exists(handle string) bool
	Create(AuthUser) (uint, error)
	Delete(handle string) error
	Get(handle string) (User, error)
}
