package binarytree

import (
	"containers"
	"fmt"
	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	left  *Node[T]
	right *Node[T]
	data  T
}

type BinaryTree[T constraints.Ordered] struct {
	root *Node[T]
}

func New[T constraints.Ordered]() BinaryTree[T] {
	return BinaryTree[T]{}
}

func (tree BinaryTree[T]) String() string {
	ch := make(chan T)
	go tree.Walk(containers.TraverseInOrder, ch)
	ret := "{"
	for val := range ch {
		ret += fmt.Sprint(" ", val)
	}
	ret += " }"
	return ret
}

func (tree *BinaryTree[T]) Insert(data T) {
	node := &Node[T]{data: data}
	if tree.root == nil {
		root := node
		tree.root = root
		return
	}
	insertNode(tree.root, node)
}

func insertNode[T constraints.Ordered](cur *Node[T], node *Node[T]) {
	if cur.data > node.data {
		if cur.left == nil {
			cur.left = node
			return
		}
		insertNode(cur.left, node)
	} else {
		if cur.right == nil {
			cur.right = node
			return
		}
		insertNode(cur.right, node)
	}
}

func (tree *BinaryTree[T]) Remove(data T) {
}

func (tree *BinaryTree[T]) Size() (ret uint) {
	ch := make(chan T)
	go tree.Walk(containers.TraverseInOrder, ch)
	for range ch {
		ret++
	}
	return
}

func (tree *BinaryTree[T]) Contains(val T) bool {
	ch := make(chan T)
	go tree.Walk(containers.TraverseInOrder, ch)
	for data := range ch {
		if data == val {
			return true
		}
	}
	return false
}

func (tree *BinaryTree[T]) Walk(traverse containers.TraverseType, ch chan T) {
	switch traverse {
	case containers.TraverseInOrder:
		tree.root.TraverseInOrder(ch)
	case containers.TraversePreOrder:
		tree.root.TraversePreOrder(ch)
	case containers.TraversePostOrder:
		tree.root.TraversePostOrder(ch)
	}
	close(ch)
}

func (node *Node[T]) TraverseInOrder(ch chan T) {
	if node == nil {
		return
	}
	node.left.TraverseInOrder(ch)
	ch <- node.data
	node.right.TraverseInOrder(ch)
}

func (node *Node[T]) TraversePreOrder(ch chan T) {
	if node == nil {
		return
	}
	ch <- node.data
	node.left.TraversePreOrder(ch)
	node.right.TraversePreOrder(ch)
}

func (node *Node[T]) TraversePostOrder(ch chan T) {
	if node == nil {
		return
	}
	node.left.TraversePostOrder(ch)
	node.right.TraversePostOrder(ch)
	ch <- node.data
}
