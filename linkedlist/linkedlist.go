package linkedlist

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func New() LinkedList {
    return LinkedList { nil }
}

func (list *LinkedList) Append(data int) {
	node := &Node{data: data, next: nil}

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

func (list *LinkedList) Prepend(data int) {
	node := &Node{data: data, next: list.head}
	list.head = node
}

func (list *LinkedList) PopFront() {
	list.head = list.head.next
}

func (list *LinkedList) PopBack() {
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

func (list *LinkedList) Len() int {
	cur := list.head
	i := 0
	for cur != nil {
		i++
        cur = cur.next
	}
	return i
}

func (list LinkedList) String() string {
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
