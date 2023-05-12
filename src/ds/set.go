package ds

type Hashable interface {
	int | string | rune
}

type Set[T Hashable] map[T]bool

func NewSet[T Hashable]() Set[T] {
	s := make(map[T]bool)
	return s
}

func (s *Set[T]) Add(element T) {
	(*s)[element] = true
}

func ToSet[T Hashable](elements []T) Set[T] {
	s := NewSet[T]()
	for _, val := range elements {
		s.Add(val)
	}
	return s
}

func (s *Set[T]) Contains(element T) bool {
	val, found := (*s)[element]
	if !found {
		return false
	}
	return val
}

func (s *Set[T]) Delete(element T) {
	delete(*s, element)
}
