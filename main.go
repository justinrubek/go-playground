package main

import (
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

func main() {
	list := linkedlist.New[int]()
	testList(&list)
}
