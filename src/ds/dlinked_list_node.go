package ds

type NodeDefFunc[T interface{}] func(d *DLinkedListNode[T])

type DLinkedListNode[T interface{}] struct {
	Prev  *DLinkedListNode[T]
	Next  *DLinkedListNode[T]
	Value T
}

func WithPrev[T interface{}](prev *DLinkedListNode[T]) NodeDefFunc[T] {
	return func(d *DLinkedListNode[T]) {
		d.Prev = prev
	}
}

func WithNext[T interface{}](nxt *DLinkedListNode[T]) NodeDefFunc[T] {
	return func(d *DLinkedListNode[T]) {
		d.Next = nxt
	}
}

func WithValue[T interface{}](v T) NodeDefFunc[T] {
	return func(d *DLinkedListNode[T]) {
		d.Value = v
	}
}

func NewDlinkedListNode[T interface{}](opts ...NodeDefFunc[T]) DLinkedListNode[T] {
	d := DLinkedListNode[T]{}
	for _, o := range opts {
		o(&d)
	}
	return d
}

func (d DLinkedListNode[T]) Remove() {
	prev := d.Prev
	nxt := d.Next
	if prev != nil {
		prev.Next = nxt
	}
	if nxt != nil {
		nxt.Prev = prev
	}
}
