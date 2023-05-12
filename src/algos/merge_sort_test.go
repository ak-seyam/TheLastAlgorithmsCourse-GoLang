package algos

import "testing"

type comparableInteger struct {
	value int
}

func (c comparableInteger) ComparedTo(c2Value int) ComparableState {
	if c.value > c2Value {
		return ComparableState(MORE)
	} else if c.value == c2Value {
		return ComparableState(EQUAL)
	} else {
		return ComparableState(LESS)
	}
}

func (c comparableInteger) Value() int {
	return c.value
}

func TestSorting(t *testing.T) {
	input := []Comparable[int]{
		comparableInteger{value: 1},
		comparableInteger{value: 2},
		comparableInteger{value: 4},
		comparableInteger{value: 5},
		comparableInteger{value: 3},
		comparableInteger{value: 6},
	}
	m := MergeSorter[int]{
		Input: input,
	}
	m.MergeSort()
	for i := 1; i < len(m.Input); i++ {
		if m.Input[i].ComparedTo(m.Input[i-1].Value()) == LESS {
			t.Fatal("not sorted")
		}
	}
}

func TestSortingWhenArrayIsEmpty(t *testing.T) {
	input := []Comparable[int]{}
	m := MergeSorter[int]{
		Input: input,
	}
	m.MergeSort()
	for i := 1; i < len(m.Input); i++ {
		if m.Input[i].ComparedTo(m.Input[i-1].Value()) == LESS {
			t.Fatal("not sorted")
		}
	}
}
