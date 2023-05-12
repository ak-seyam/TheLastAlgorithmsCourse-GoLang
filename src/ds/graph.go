package ds

type AdjListConnection struct {
	To     int
	Weight int
}

type AdjListNode[T interface{}] struct {
	Connections []AdjListConnection
	Value       T
}

type AdjMatrixConnection struct {
	Weight int
}

type AdjMatrixNode[T interface{}] struct {
	Connection AdjMatrixConnection
	Value      T
}
