package ds

type queue[T interface{}] struct {
	topIdx int
	inner  []T
}

func (q *queue[T]) Enqueue(value T) {
	if q.topIdx == -1 {
		q.topIdx = 0
	}
	q.inner = append(q.inner, value)
}

func NewQueue[T interface{}]() queue[T] {
	return queue[T]{topIdx: -1, inner: []T{}}
}

func (q *queue[T]) Dequeu() T {
	res := q.inner[q.topIdx]
	q.topIdx++
	return res
}

func (q *queue[T]) IsEmpty() bool {
	if q.topIdx == -1 || q.topIdx == len(q.inner) {
		return true
	}
	return false
}
