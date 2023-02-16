package people

type Identifier[T any] interface {
	Identify() T
}

func IntoMap[T Identifier[U], U comparable](slice []T) map[U]T {
	m := make(map[U]T, len(slice))
	for _, v := range slice {
		m[v.Identify()] = v
	}
	return m
}

func IntoIDs[T Identifier[U], U comparable](slice []T) []U {
	ids := make([]U, len(slice))
	for i, v := range slice {
		ids[i] = v.Identify()
	}
	return ids
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

func (m Message) Identify() uint {
	return m.ID
}

func (t Thread) Identify() uint {
	return t.ID
}

func (u ThreadUser) Identify() uint {
	return u.ID
}
