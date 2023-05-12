package ds

import "testing"

func TestHeapInsertion(t *testing.T) {
	h := NewMinHeap[int, ComparableInt]()
	values := []ComparableInt{
		NewComparableInt(4),
		NewComparableInt(5),
		NewComparableInt(6),
		NewComparableInt(8),
		NewComparableInt(3),
	}
	for _, v := range values {
		h.Insert(v)
	}
	expected := []ComparableInt{
		NewComparableInt(3),
		NewComparableInt(4),
		NewComparableInt(6),
		NewComparableInt(8),
		NewComparableInt(5),
	}
	for i, exV := range expected {
		if exV.CompareTo(h.data[i]) != EQUAL {
			t.Fatalf(
				"value in heap is not sorted it should be %d but it is %d\n",
				exV.Value(),
				h.data[i],
			)
		}
	}
}

func TestDelete(t *testing.T) {
	h := NewMinHeap[int, ComparableInt]()
	values := []ComparableInt{
		NewComparableInt(4),
		NewComparableInt(5),
		NewComparableInt(6),
		NewComparableInt(8),
		NewComparableInt(3),
	}
	for _, v := range values {
		h.Insert(v)
	}
	for _, num := range []int{1, 2, 3} {
		v, err := h.Pop()
		if err != nil {
			t.Fatalf("%s\n", err.Error())
		}
		if num == 1 && v.Value() != 3 {
			t.Fatalf("popped value is not correct it should be %d but it is %d", 3, v.Value())
		} else if num == 2 && v.Value() != 4 {
			t.Fatalf("popped value is not correct it should be %d but it is %d", 4, v.Value())
		} else if num == 3 && v.Value() != 5 {
			t.Fatalf("popped value is not correct it should be %d but it is %d", 5, v.Value())
		}
	}
}
