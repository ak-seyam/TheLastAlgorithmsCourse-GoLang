package ds

type Stack[T interface{}] struct {
	topIdx int
	data   []T
}

func (s *Stack[T]) Push(item T) {
	s.data = append(s.data, item)
	s.topIdx = len(s.data) - 1
}

func (s *Stack[T]) Pop() T {
	res := s.data[s.topIdx]
	s.topIdx--
	return res
}

func NewStack[T interface{}]() Stack[T] {
	return Stack[T]{
		topIdx: -1,
		data:   []T{},
	}
}

func (s Stack[T]) IsEmpty() bool {
	return s.topIdx == -1
}
