package algos

type ComparableState int8

const (
	LESS ComparableState = iota
	MORE
	EQUAL
)

type Comparable[T interface{}] interface {
	ComparedTo(element T) ComparableState
	Value() T
}

type MergeSorter[T interface{}] struct {
	Input []Comparable[T]
}

func (m MergeSorter[T]) MergeSort() {
	m._mergeSort(0, len(m.Input)-1)
}

func (m MergeSorter[T]) _mergeSort(start, end int) {
	if start >= end {
		return
	}
	p := m.pivoting(start, end)
	m._mergeSort(start, p-1)
	m._mergeSort(p+1, end)
}

func (m MergeSorter[T]) pivoting(start, end int) int {
	l, r := start, end
	pivot := l
	pivotVal := m.Input[pivot].Value()
	if l == r {
		return start
	}
	for l < r {
		for ; (m.Input[l].ComparedTo(pivotVal) == LESS || m.Input[l].ComparedTo(pivotVal) == EQUAL) && l < end; l++ {
		}
		for ; (m.Input[r].ComparedTo(pivotVal) == MORE) && r > start; r-- {
		}
		if l < r {
			m.Input[l], m.Input[r] = m.Input[r], m.Input[l]
		}
	}
	m.Input[pivot], m.Input[r] = m.Input[r], m.Input[pivot]
	return r
}

func (m MergeSorter[T]) swap(a, b int) {
	tmp := m.Input[a]
	m.Input[a] = m.Input[b]
	m.Input[b] = tmp
}
