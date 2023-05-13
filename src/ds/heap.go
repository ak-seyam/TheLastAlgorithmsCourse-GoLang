package ds

import (
	"errors"
)

type MinHeap[T interface{}, V Comparable[T]] struct {
	data   []V
	length int
}

func NewMinHeap[T interface{}, V Comparable[T]]() MinHeap[T, V] {
	return MinHeap[T, V]{
		data:   []V{},
		length: 0,
	}
}

func (h *MinHeap[T, V]) Length() int {
	return h.length
}

func parent(idx int) (int, error) {
	if idx < 0 {
		return -1, errors.New("out of boundries")
	}
	return (idx - 1) / 2, nil
}

/*
return children (left, right) and error
*/
func children(idx int) (int, int, error) {
	if idx < 0 {
		return -1, -1, errors.New("out of boundries")
	}
	return (idx * 2) + 1, (idx * 2) + 2, nil
}

func (h *MinHeap[T, V]) HeapifyUp(idx int) (int, error) {
	if idx <= 0 {
		return -1, nil
	}
	parent, err := parent(idx)
	if err != nil {
		return -1, err
	}
	parentValue := h.data[parent]
	if h.data[idx].CompareTo(parentValue) == LESS {
		h.data[idx], h.data[parent] = h.data[parent], h.data[idx]
		h.HeapifyUp(parent)
	}
	return idx, nil
}

func (h *MinHeap[T, V]) HeapifyDown(idx int) error {
	if idx >= h.length {
		return nil
	}
	left, right, err := children(idx)
	if err != nil {
		return err
	}
	if left >= h.length { // right will always be >= h.length if this condition is true so we shouldn't check it
		return nil
	}
	lVal := h.data[left]
	rVal := h.data[right]
	currentValue := h.data[idx]
	if lVal.CompareTo(rVal) == LESS && currentValue.CompareTo(lVal) == MORE {
		h.data[left], h.data[idx] = h.data[idx], h.data[left]
		h.HeapifyDown(left)
	} else if rVal.CompareTo(lVal) == LESS && currentValue.CompareTo(rVal) == MORE {
		h.data[right], h.data[idx] = h.data[idx], h.data[right]
		h.HeapifyDown(right)
	}
	return nil
}

func (h *MinHeap[T, V]) Insert(val V) int {
	h.data = append(h.data, val)
	newIdx, err := h.HeapifyUp(h.length)
	if err != nil {
		panic(err)
	}
	h.length++
	return newIdx
}

func (h *MinHeap[T, V]) Pop() (*V, error) {
	if h.length == 0 {
		return nil, errors.New("empty heap")
	}
	val := h.data[0]
	if h.length == 1 {
		h.data = []V{}
		h.length = 0
		return &val, nil
	}
	h.length--
	h.data[0] = h.data[h.length]
	h.HeapifyDown(0)
	return &val, nil
}

func (h *MinHeap[T, V]) Update(idx int, newVal V) {
	oldVal := h.data[idx]
	h.data[idx] = newVal
	if newVal.CompareTo(oldVal) == MORE {
		h.HeapifyDown(idx)
	} else {
		h.HeapifyUp(idx)
	}
}
