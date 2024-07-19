package linkedlist

import (
	"errors"
	"fmt"
)

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

func (list *LinkedList[T]) PopFront() (T, error) {
	if list.head == nil {
		return *new(T), errors.New("popped empty list")
	}
	ret := list.head.data
	list.head = list.head.next
	return ret, nil
}

func (list *LinkedList[T]) PopBack() (T, error) {
	if list.head == nil {
		return *new(T), errors.New("popped empty list")
	}
	if list.head.next == nil {
		ret := list.head.data
		list.head = nil
		return ret, nil
	}
	cur := list.head
	for cur.next.next != nil {
		cur = cur.next
	}
	ret := cur.data
	cur.next = nil
	return ret, nil
}

func (list *LinkedList[T]) Len() uint {
	cur := list.head
	var i uint = 0
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
