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
	previous := getInitializedSlice(len(g), -1)
	heapIndexes := make([]int, len(g))
	distancesHeap := getDistancesHeap(len(g), &heapIndexes)
	distances := getInitializedSlice(len(g), math.MaxInt)
	startDist, err := distancesHeap.Pop()
	if err != nil {
		panic(err)
	}
	startDist.value.Distance = 0
	distancesHeap.Insert(*startDist)
	distances[a.ID] = 0
	for shortest := shortestDistance(distancesHeap); shortest != nil; shortest = shortestDistance(distancesHeap) {
		if shortest.value.ID == target { // you reach your goal your shortest is the target (you cannot go any shorter to your dist)
			break
		}
		for _, con := range g[shortest.value.ID].Connections {
			dist, updated := min(distances[con.To], con.Weight+distances[shortest.value.ID])
			if updated {
				distances[con.To] = dist
				previous[con.To] = shortest.value.ID
				idx := heapIndexes[con.To]
				distancesHeap.Update(idx, ComparableDistance{
					value: Distance{
						ID:       con.To,
						Distance: dist,
					},
				})
			}
		}
	}
	src := a.ID
	return collect(previous, src, target)
}

type Distance struct {
	ID       int
	Distance int
}

type ComparableDistance struct {
	value Distance
}

func (cd ComparableDistance) Value() Distance {
	return cd.value
}

func (cd ComparableDistance) CompareTo(cd2 Comparable[Distance]) ComparisonResult {
	if cd.value.Distance > cd2.Value().Distance {
		return MORE
	} else if cd.value.Distance < cd2.Value().Distance {
		return LESS
	} else {
		return EQUAL
	}
}

func getDistancesHeap(gl int, nodes *[]int) MinHeap[Distance, ComparableDistance] {
	h := NewMinHeap[Distance, ComparableDistance]()
	for i := 0; i < gl; i++ {
		node := ComparableDistance{
			value: Distance{
				ID:       i,
				Distance: math.MaxInt,
			},
		}
		idx := h.Insert(node)
		(*nodes)[i] = idx
	}
	return h
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

func shortestDistance(distanceHeap MinHeap[Distance, ComparableDistance]) *ComparableDistance {
	item, err := distanceHeap.Pop()
	if err != nil {
		panic(err)
	}
	return item
}

func getInitializedSlice(length int, targetInit int) []int {
	res := make([]int, length)
	for i := 0; i < length; i++ {
		res[i] = targetInit
	}
	return res
}
