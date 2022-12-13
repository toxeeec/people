package people

type Identifier[T any] interface {
	Identify() T
}

func (u User) Identify() string {
	return u.Handle
}

func (p Post) Identify() uint {
	return p.ID
}

func (p PostResponse) Identify() uint {
	return p.Data.ID
}
