package ds

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
