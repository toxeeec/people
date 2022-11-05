package people

type Identifier[T any] interface {
	Identify() T
}
