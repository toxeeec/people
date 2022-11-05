package people

func (u User) Identify() string {
	return u.Handle
}

type UserService interface {
	Create(AuthUser) (uint, error)
	Exists(handle string) bool
	Delete(handle string) error
	Get(handle string) (AuthUser, error)
	Follow(id uint, handle string) error
	Unfollow(id uint, handle string) error
	IsFollowing(id uint, handle string) (bool, error)
	IsFollowed(id uint, handle string) (bool, error)
	Following(id uint, p HandlePagination) (Users, error)
	Followers(id uint, p HandlePagination) (Users, error)
}
