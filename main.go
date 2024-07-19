package main

import (
	"binarytree"
	"containers"
	"fmt"
	"linkedlist"
)
import _ "vec"

func testList(list containers.List[int]) {
	for i := range 10 {
		list.Append(i)
	}
	fmt.Println(list)
	if back, err := list.PopBack(); err == nil {
		fmt.Println("popped from back: ", back)
	}
	fmt.Println(list)

	list.Append(11)
	list.Prepend(0)
	fmt.Println(list)
}

func testTree(tree containers.Tree[int]) {
	fmt.Println(tree)
	for i := range 10 {
		tree.Insert(i)
	}
	fmt.Println(tree)
	fmt.Println("tree size:", tree.Size())
	fmt.Println("tree contains 2:", tree.Contains(2))
	fmt.Println("tree does not contain -1:", !tree.Contains(-1))

}

func main() {
	list := linkedlist.New[int]()
	testList(&list)

	tree := binarytree.New[int]()
	testTree(&tree)
}
