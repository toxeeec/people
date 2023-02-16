package set

type Set[T comparable] struct {
	values map[T]struct{}
	slice  []T
}

func New[T comparable]() Set[T] {
	set := Set[T]{values: make(map[T]struct{})}
	return set
}

func FromSlice[T comparable](slice []T) Set[T] {
	set := Set[T]{values: make(map[T]struct{})}
	for _, v := range slice {
		set.Add(v)
	}
	return set
}

func (s *Set[T]) Add(val T) {
	if _, ok := s.values[val]; ok {
		return
	}
	s.values[val] = struct{}{}
	s.slice = append(s.slice, val)
}

func (s Set[T]) Slice() []T {
	return s.slice
}
