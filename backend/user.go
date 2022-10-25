package people

type UserService interface {
	Create(AuthUser) (uint, error)
	Exists(handle string) bool
	Delete(handle string) error
	Get(handle string) (User, error)
	Follow(id uint, handle string) error
	Unfollow(id uint, handle string) error
}
