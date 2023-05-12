package ds

type AdjListNode[T interface{}] struct {
	Connections []AdjListConnection[T]
	Value       T
}

type AdjListConnection[T interface{}] struct {
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
	s.Push(&a)
	for !s.IsEmpty() {
		node := s.Pop()
		cb(node)
		for _, con := range node.Connections {
			s.Push(g[con.To])
		}
	}
}

func (a AdjListNode[T]) BFS(g []*AdjListNode[T], cb AdjListTraversalCallBack[T]) {
	s := NewQueue[*AdjListNode[T]]()
	s.Enqueue(&a)
	for !s.IsEmpty() {
		node := s.Dequeu()
		cb(node)
		for _, con := range node.Connections {
			s.Enqueue(g[con.To])
		}
	}
}

func NewAdjListGraph[T interface{}]() []AdjListNode[T] {
	return []AdjListNode[T]{}
}
