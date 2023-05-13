package ds

import "testing"

func TestLRUPut(t *testing.T) {
	lru := NewLRU[string, int](4)
	lru.Put("number1", 1)
	lru.Put("number2", 2)
	lru.Put("number3", 3)
	lru.Put("number4", 4)
	lru.Put("number5", 5)
	if len(lru.m) != 4 {
		t.Fatalf("error in capacity maintenance the cap should be %d but it is %d\n", 4, len(lru.m))
	}
	expectedKeys := []string{
		"number2",
		"number3",
		"number4",
		"number5",
	}
	for _, v := range expectedKeys {
		_, mFound := lru.m[v]
		if !mFound {
			t.Fatalf("keys should have %s but it doesn't", v)
		}
	}
	dllOrderedValues := []int{5, 4, 3, 2}
	curr := lru.dll.Head
	for _, v := range dllOrderedValues {
		if v != curr.Value.value {
			t.Fatalf("order error it should be %d but it is %d", v, curr.Value.value)
		}
		curr = curr.Next
	}
}

func TestGetting(t *testing.T) {
	lru := NewLRU[string, int](4)
	lru.Put("number1", 1)
	lru.Put("number2", 2)
	lru.Put("number3", 3)
	lru.Put("number4", 4)
	lru.Put("number5", 5)

	val, found := lru.Get("number1")
	if found {
		t.Fatalf("%d should not be found\n", *val)
	}

	val5, found := lru.Get("number5")
	if !found {
		t.Fatalf("%d should be found\n", *val5)
	}

	// messing the order
	lru.Get("number3")

	expectedOrderedValues := []int{3, 5, 4, 2}
	curr := lru.dll.Head
	for _, v := range expectedOrderedValues {
		if v != curr.Value.value {
			t.Fatalf("order error it should be %d but it is %d", v, curr.Value.value)
		}
		curr = curr.Next
	}
}
