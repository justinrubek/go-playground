package linkedlist

import "fmt"

type Node[T any] struct {
	data T
	next *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
}

func New[T any]() LinkedList[T] {
	return *new(LinkedList[T])
}

func (list *LinkedList[T]) Append(data T) {
	node := &Node[T]{data: data, next: nil}

	if list.head == nil {
		list.head = node
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = node
}

func (list *LinkedList[T]) Prepend(data T) {
	node := &Node[T]{data: data, next: list.head}
	list.head = node
}

func (list *LinkedList[T]) PopFront() {
	list.head = list.head.next
}

func (list *LinkedList[T]) PopBack() {
	if list.head == nil {
		return
	}
	if list.head.next == nil {
		list.head = nil
		return
	}

	cur := list.head
	for cur.next.next != nil {
		cur = cur.next
	}
	cur.next = nil
}

func (list *LinkedList[T]) Len() int {
	cur := list.head
	i := 0
	for cur != nil {
		i++
		cur = cur.next
	}
	return i
}

func (list LinkedList[T]) String() string {
	str := ""
	str += fmt.Sprint("[")
	cur := list.head
	for cur != nil {
		str += fmt.Sprint(" ", cur.data)
		cur = cur.next
	}
	str += fmt.Sprint(" ]")
	return str
}
