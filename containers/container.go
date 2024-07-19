package containers

type List[T any] interface {
	Append(T)
	Len() uint
	Prepend(T)
	PopFront() (T, error)
	PopBack() (T, error)
}

type TraverseType int

const (
	TraversePreOrder TraverseType = iota
	TraverseInOrder
	TraversePostOrder
)

type Tree[T any] interface {
	Contains(T) bool
	Insert(T)
	Remove(T)
	Size() uint
	Walk(TraverseType, chan T)
}
