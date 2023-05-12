package ds

import "testing"

/*
where s is the start of relation

		0      1
	   s|s\    s|
	  10|6 10   3
	    |s   \s |
		3 s-12- 2
*/
func getTestAdjGraph() []*AdjListNode[int] {
	g := NewAdjListGraph[int]()
	n0 := &AdjListNode[int]{
		Value: 1,
		ID:    0,
		Connections: []AdjListConnection{
			{To: 2, Weight: 10},
			{To: 3, Weight: 10},
		},
	}
	g = append(g, n0)
	n1 := &AdjListNode[int]{
		Value: 12,
		ID:    1,
		Connections: []AdjListConnection{
			{To: 2, Weight: 3},
		},
	}
	g = append(g, n1)
	n2 := &AdjListNode[int]{
		Value:       -5,
		ID:          2,
		Connections: []AdjListConnection{},
	}
	g = append(g, n2)
	n3 := &AdjListNode[int]{
		Value: 9,
		ID:    3,
		Connections: []AdjListConnection{
			{To: 0, Weight: 6},
			{To: 2, Weight: 12},
		},
	}
	g = append(g, n3)
	return g
}

func TestDFSAdjList(t *testing.T) {
	g := getTestAdjGraph()
	startingNode := g[0]
	actualIDs := []int{}
	expected := []int{0, 3, 2}
	startingNode.DFS(g, func(n *AdjListNode[int]) {
		actualIDs = append(actualIDs, n.ID)
	})
	for i, id := range expected {
		if actualIDs[i] != id {
			t.Fatalf("invalid order %d expected but %d was found", id, actualIDs[i])
		}
	}
}

func TestBFSAdjList(t *testing.T) {
	g := getTestAdjGraph()
	startingNode := g[0]
	actualIDs := []int{}
	expected := []int{0, 2, 3}
	startingNode.BFS(g, func(n *AdjListNode[int]) {
		actualIDs = append(actualIDs, n.ID)
	})
	for i, id := range expected {
		if actualIDs[i] != id {
			t.Fatalf("invalid order %d expected but %d was found", id, actualIDs[i])
		}
	}
}
