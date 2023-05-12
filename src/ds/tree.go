package ds

type ComparisonResult int8

const (
	LESS ComparisonResult = iota
	MORE
	EQUAL
)

type Ordering int8

const (
	PRE Ordering = iota
	IN
	POST
)

type Comparable[T interface{}] interface {
	CompareTo(c Comparable[T]) ComparisonResult
	Value() T
}

type ComparableInt struct {
	value int
}

func (c ComparableInt) Value() int {
	return c.value
}

func (c ComparableInt) CompareTo(c1 Comparable[int]) ComparisonResult {
	if c.Value() > c1.Value() {
		return ComparisonResult(MORE)
	} else if c.Value() == c1.Value() {
		return ComparisonResult(EQUAL)
	} else {
		return ComparisonResult(LESS)
	}
}

func NewComparableInt(v int) ComparableInt {
	return ComparableInt{value: v}
}

type BTNode[T interface{}, V Comparable[T]] struct {
	Left  *BTNode[T, V]
	Right *BTNode[T, V]
	Inner V
}

func (b *BTNode[T, V]) ForEach(f func(idx int, n T), o Ordering) {
	idx := 0
	traverse(b, &idx, f, o)
}

func traverse[T interface{}, V Comparable[T]](n *BTNode[T, V], idx *int, f func(idx int, n T), o Ordering) {
	if n == nil {
		return
	}
	if o == PRE {
		f(*idx, n.Inner.Value())
		(*idx)++
	}
	traverse(n.Left, idx, f, o)
	if o == IN {
		f(*idx, n.Inner.Value())
		(*idx)++
	}
	traverse(n.Right, idx, f, o)
	if o == POST {
		f(*idx, n.Inner.Value())
		(*idx)++
	}
}

func (b *BTNode[T, V]) BFS(cb func(node *BTNode[T, V])) {
	q := NewQueue[*BTNode[T, V]]()
	q.Enqueue(b)
	for !q.IsEmpty() {
		curr := q.Dequeu()
		cb(curr)
		if curr.Left != nil {
			q.Enqueue(curr.Left)
		}
		if curr.Right != nil {
			q.Enqueue(curr.Right)
		}
	}
}

func (b *BTNode[T, V]) Equals(bTree *BTNode[T, V]) bool {
	return _equals(b, bTree)
}

func _equals[T interface{}, V Comparable[T]](b1 *BTNode[T, V], b2 *BTNode[T, V]) bool {
	if b1 == nil && b2 == nil {
		return true
	} else if b1 == nil || b2 == nil {
		return false
	} else if b1.Inner.CompareTo(b2.Inner) != EQUAL {
		return false
	}
	left := _equals(b1.Left, b2.Left)
	right := _equals(b1.Right, b2.Right)
	return left && right
}

func (b *BTNode[T, V]) Find(v Comparable[T]) *BTNode[T, V] {
	return _find(b, v)
}

func _find[T interface{}, V Comparable[T]](b *BTNode[T, V], v Comparable[T]) *BTNode[T, V] {
	if b == nil {
		return nil
	}
	if b.Inner.CompareTo(v) == EQUAL {
		return b
	} else if b.Inner.CompareTo(v) == MORE {
		return _find(b.Left, v)
	} else {
		return _find(b.Right, v)
	}
}
