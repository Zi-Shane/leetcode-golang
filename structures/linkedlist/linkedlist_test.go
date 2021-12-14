package linkedlist

import "testing"

func TestListnode(t *testing.T) {
	list := NewLinkedList()
	list.AddInt(1)
	list.AddInt(2)
}
