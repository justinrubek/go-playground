package binarytree

import "testing"

func TestInsert(t *testing.T) {
	tree := New[int]()
	for i := range 5 {
		tree.Insert(i)
	}

	str := tree.String()
	if str != "{ 0 1 2 3 4 }" {
		t.Fatal("insert failed, got:", str)
	}
}

func TestRemove(t *testing.T) {
	tree := New[int]()
	for i := range 5 {
		tree.Insert(i)
	}
	tree.Remove(2)

	str := tree.String()
	if str != "{ 0 1 3 4 }" {
		t.Fatal("remove failed, got:", str)
	}
}
