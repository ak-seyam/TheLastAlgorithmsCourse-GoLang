package ds

import "testing"

func TestDlinkedListAppending(t *testing.T) {
	headNode := &DLinkedListNode[int]{Value: 1}
	dll := DLinkedList[int]{
		Head: headNode,
		Tail: headNode,
	}
	dll.Append(5)
	curr := dll.Head
	listRepresentation := []int{1, 5}
	for _, v := range listRepresentation {
		if curr.Value != v {
			t.Fatalf("invalid value it should be %d but it is %d", v, curr.Value)
		}
		curr = curr.Next
	}
	if curr != nil {
		t.Fatal("list didn't match representation")
	}
}

func TestDlinkedListAppendingIfListIsEmpty(t *testing.T) {
	dll := DLinkedList[int]{}
	dll.Append(5)
	curr := dll.Head
	listRepresentation := []int{5}
	for _, v := range listRepresentation {
		if curr.Value != v {
			t.Fatalf("invalid value it should be %d but it is %d", v, curr.Value)
		}
		curr = curr.Next
	}
	if curr != nil {
		t.Fatal("list didn't match representation")
	}
}
