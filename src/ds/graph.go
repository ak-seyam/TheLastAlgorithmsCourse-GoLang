package ds

import "math"

type AdjListNode[T interface{}] struct {
	Connections []AdjListConnection
	Value       T
	ID          int
}

type AdjListConnection struct {
	To     int
	Weight int
}

type AdjMatrixConnection struct {
	Weight int
}

type AdjMatrixNode[T interface{}] struct {
	Connection AdjMatrixConnection
	Value      T
}

type AdjListTraversalCallBack[T interface{}] func(n *AdjListNode[T])

func (a AdjListNode[T]) DFS(g []*AdjListNode[T], cb AdjListTraversalCallBack[T]) {
	s := NewStack[*AdjListNode[T]]()
	vs := NewSet[int]()
	s.Push(&a)
	for !s.IsEmpty() {
		node := s.Pop()
		vs.Add(node.ID)
		cb(node)
		for _, con := range node.Connections {
			if !vs.Contains(g[con.To].ID) {
				s.Push(g[con.To])
			}
		}
	}
}

func (a AdjListNode[T]) BFS(g []*AdjListNode[T], cb AdjListTraversalCallBack[T]) {
	s := NewQueue[*AdjListNode[T]]()
	vs := NewSet[int]()
	s.Enqueue(&a)
	for !s.IsEmpty() {
		node := s.Dequeu()
		vs.Add(a.ID)
		cb(node)
		for _, con := range node.Connections {
			if !vs.Contains(g[con.To].ID) {
				s.Enqueue(g[con.To])
			}
		}
	}
}

func NewAdjListGraph[T interface{}]() []*AdjListNode[T] {
	return []*AdjListNode[T]{}
}

func (a AdjListNode[T]) ShortestPathDijkstra(g []*AdjListNode[T], target int) []int {
	seenSet := NewSet[int]()
	previous := getInitializedSlice(len(g), -1)
	distances := getInitializedSlice(len(g), math.MaxInt)
	distances[a.ID] = 0
	for haveUnseen(seenSet, distances) {
		shortest := shortestUnseen(seenSet, distances)
		if shortest == -1 {
			break
		}
		seenSet.Add(shortest)
		if shortest == target { // you reach your goal your shortest is the target (you cannot go any shorter to your dist)
			break
		}
		for _, con := range g[shortest].Connections {
			if seenSet.Contains(con.To) { // because it was the local minimum at some point if this was the case it cannot be longest
				continue
			}
			dist, updated := min(distances[con.To], con.Weight+distances[shortest])
			if updated {
				distances[con.To] = dist
				previous[con.To] = shortest
			}
		}
	}
	src := a.ID
	return collect(previous, src, target)
}

func collect(previous []int, src, target int) []int {
	curr := target
	path := []int{}
	path = append(path, curr)
	for curr != src {
		curr = previous[curr]
		if curr == -1 {
			break
		}
		path = append(path, curr)
		if curr == src {
			reverse(path)
			return path
		}
	}
	return []int{}
}

func reverse(x []int) {
	for i := 0; i < len(x)/2; i++ {
		l := i
		r := len(x) - 1 - i
		x[l], x[r] = x[r], x[l]
	}
}

func min(old, new int) (int, bool) {
	if old <= new {
		return old, false
	} else {
		return new, true
	}
}

func shortestUnseen(visitedSet Set[int], distances []int) int {
	minDist := math.MaxInt
	shortestID := -1
	for ID, d := range distances {
		if d < minDist && !visitedSet.Contains(ID) {
			shortestID = ID
			minDist = d
		}
	}
	return shortestID
}

func haveUnseen(visitedSet Set[int], distances []int) bool {
	for ID, d := range distances {
		if d != math.MaxInt && !visitedSet.Contains(ID) { // d != math.MaxInt means it is reachable from our starting point
			return true
		}
	}
	return false
}

func getInitializedSlice(length int, targetInit int) []int {
	res := make([]int, length)
	for i := 0; i < length; i++ {
		res[i] = targetInit
	}
	return res
}
