package containers

type List[T any] interface {
    Append(T)
    Len() uint
    Prepend(T)
    PopFront() (T, error)
    PopBack() (T, error)
}
