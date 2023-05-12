package ds

type AdjListNode[T interface{}] struct {
	Connections []AdjListConnection[T]
	Value       T
}

type AdjListConnection[T interface{}] struct {
	To     *AdjListNode[T]
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

func (a AdjListNode[T]) DFS(cb AdjListTraversalCallBack[T]) {
	s := NewStack[*AdjListNode[T]]()
	s.Push(&a)
	for !s.IsEmpty() {
		node := s.Pop()
		cb(node)
		for _, con := range node.Connections {
			s.Push(con.To)
		}
	}
}

func (a AdjListNode[T]) BFS(cb AdjListTraversalCallBack[T]) {
	s := NewQueue[*AdjListNode[T]]()
	s.Enqueue(&a)
	for !s.IsEmpty() {
		node := s.Dequeu()
		cb(node)
		for _, con := range node.Connections {
			s.Enqueue(con.To)
		}
	}
}
