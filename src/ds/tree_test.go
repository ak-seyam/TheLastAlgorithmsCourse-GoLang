package ds

import (
	"testing"
)

func TestTreeTraversal(t *testing.T) {
	root := BTNode[int, ComparableInt]{
		Inner: NewComparableInt(2),
	}
	left := BTNode[int, ComparableInt]{
		Inner: NewComparableInt(-1),
	}
	right := BTNode[int, ComparableInt]{
		Inner: NewComparableInt(3),
	}
	root.Left = &left
	root.Right = &right
	represenation_pre := []int{2, -1, 3}
	root.ForEach(func(idx, n int) {
		if represenation_pre[idx] != n {
			t.Fatalf("invalid index it should be %d but it is %d\n", represenation_pre[idx], n)
		}
	}, PRE)
	represenation_in := []int{-1, 2, 3}
	root.ForEach(func(idx, n int) {
		if represenation_in[idx] != n {
			t.Fatalf("invalid index it should be %d but it is %d\n", represenation_in[idx], n)
		}
	}, IN)
	represenation_post := []int{-1, 3, 2}
	root.ForEach(func(idx, n int) {
		if represenation_post[idx] != n {
			t.Fatalf("invalid index it should be %d but it is %d\n", represenation_post[idx], n)
		}
	}, POST)
}

func TestBFS(t *testing.T) {
	root := BTNode[int, ComparableInt]{
		Inner: NewComparableInt(2),
	}
	left := BTNode[int, ComparableInt]{
		Inner: NewComparableInt(-1),
	}
	right := BTNode[int, ComparableInt]{
		Inner: NewComparableInt(3),
	}
	root.Left = &left
	root.Right = &right
	rep := []int{2, -1, 3}
	idx := 0
	root.BFS(func(node *BTNode[int, ComparableInt]) {
		if rep[idx] != node.Inner.Value() {
			t.Fatalf("invalid value it should be %d and it is %d", rep[idx], node.Inner)
		}
		idx++
	})
}

func TestEqualityBothAreEqual(t *testing.T) {
	root1 := BTNode[int, ComparableInt]{
		Inner: NewComparableInt(2),
	}
	left1 := BTNode[int, ComparableInt]{
		Inner: NewComparableInt(-1),
	}
	right1 := BTNode[int, ComparableInt]{
		Inner: NewComparableInt(3),
	}
	root1.Left = &left1
	root1.Right = &right1

	root2 := BTNode[int, ComparableInt]{
		Inner: NewComparableInt(2),
	}
	left2 := BTNode[int, ComparableInt]{
		Inner: NewComparableInt(-1),
	}
	right2 := BTNode[int, ComparableInt]{
		Inner: NewComparableInt(3),
	}
	root2.Left = &left2
	root2.Right = &right2

	if !root1.Equals(&root2) {
		t.Fatal("equality check not valid")
	}
}

func TestEqualityBothValueNotEqual(t *testing.T) {
	root1 := BTNode[int, ComparableInt]{
		Inner: NewComparableInt(3),
	}
	left1 := BTNode[int, ComparableInt]{
		Inner: NewComparableInt(-1),
	}
	right1 := BTNode[int, ComparableInt]{
		Inner: NewComparableInt(3),
	}
	root1.Left = &left1
	root1.Right = &right1

	root2 := BTNode[int, ComparableInt]{
		Inner: NewComparableInt(2),
	}
	left2 := BTNode[int, ComparableInt]{
		Inner: NewComparableInt(-1),
	}
	right2 := BTNode[int, ComparableInt]{
		Inner: NewComparableInt(3),
	}
	root2.Left = &left2
	root2.Right = &right2

	if root1.Equals(&root2) {
		t.Fatal("equality check not valid (values not equal)")
	}
}

func TestEqualityBothStructureNotEqual(t *testing.T) {
	root1 := BTNode[int, ComparableInt]{
		Inner: NewComparableInt(3),
	}
	left1 := BTNode[int, ComparableInt]{
		Inner: NewComparableInt(-1),
	}
	right1 := BTNode[int, ComparableInt]{
		Inner: NewComparableInt(3),
	}
	root1.Left = &left1
	root1.Left.Right = &right1

	root2 := BTNode[int, ComparableInt]{
		Inner: NewComparableInt(2),
	}
	left2 := BTNode[int, ComparableInt]{
		Inner: NewComparableInt(-1),
	}
	right2 := BTNode[int, ComparableInt]{
		Inner: NewComparableInt(3),
	}
	root2.Left = &left2
	root2.Right = &right2

	if root1.Equals(&root2) {
		t.Fatal("equality check not valid (structur not equal)")
	}
}

func TestFinding(t *testing.T) {
	root := getTestNode(7)
	left := getTestNode(5)
	right := getTestNode(10)
	root.Left = &left
	root.Right = &right
	comp := NewComparableInt(5)
	node := root.Find(comp)
	if node == nil {
		t.Fatal("finding failed the value existed!")
	}
	if node.Inner.CompareTo(comp) != EQUAL {
		t.Fatal("value is not correct!")
	}
}

func TestFindingDoesNotExist(t *testing.T) {
	root := getTestNode(7)
	left := getTestNode(5)
	right := getTestNode(10)
	root.Left = &left
	root.Right = &right
	node := root.Find(NewComparableInt(9))
	if node != nil {
		t.Fatal("finding failed the value should not exist!")
	}
}
func getTestNode(v int) BTNode[int, ComparableInt] {
	return BTNode[int, ComparableInt]{
		Inner: NewComparableInt(v),
	}
}
