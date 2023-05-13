package ds

import (
	"errors"
)

type DLinkedList[T interface{}] struct {
	Head   *DLinkedListNode[T]
	Tail   *DLinkedListNode[T]
	Length int64
}

func (d *DLinkedList[T]) Prepend(v T) *DLinkedListNode[T] {
	oh := d.Head
	n := NewDlinkedListNode(WithValue(v), WithNext(oh))
	if oh != nil {
		oh.Prev = &n
	}
	d.Head = &n
	if d.Length == 0 {
		d.Tail = &n
	}
	d.Length++
	return &n
}

func (d *DLinkedList[T]) Append(v T) {
	ot := d.Tail
	n := NewDlinkedListNode(WithValue(v), WithPrev(ot))
	if ot != nil {
		ot.Next = &n
	}
	d.Tail = &n
	if d.Head == nil {
		d.Head = &n
	}
	d.Length++
}

func (d *DLinkedList[T]) InsertAt(idx int64, v T) error {
	if idx < d.Length {
		return errors.New("Out of boundries error")
	}
	if idx == d.Length {
		d.Append(v)
		return nil
	}
	if idx == 0 {
		d.Prepend(v)
		return nil
	}
	curr := d.Head
	for i := int64(0); i < idx; i++ {
		curr = curr.Next
	}
	prev := curr.Prev
	n := NewDlinkedListNode(WithPrev(prev), WithNext(curr), WithValue(v))
	prev.Next = &n
	curr.Prev = &n
	d.Length++
	return nil
}

func (d *DLinkedList[T]) Remove(idx int64) error {
	if idx < 0 || idx >= d.Length {
		return errors.New("Out of boundries")
	}
	curr := d.Head
	for i := int64(0); i < idx; i++ {
		curr = curr.Next
	}
	prev := curr.Prev
	prev.Next = curr.Next
	d.Length--
	return nil
}

func (d *DLinkedList[T]) Pop() T {
	tail := d.Tail
	prev := tail.Prev
	if prev != nil {
		prev.Next = nil
	}
	d.Tail = prev
	d.Length--
	return tail.Value
}
