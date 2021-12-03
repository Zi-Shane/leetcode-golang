package leetcode

import (
	"fmt"
	"leetcode/structures"
	"testing"
)

// Input: list1 = [1,2,4], list2 = [1,3,4]
// Output: [1,1,2,3,4,4]

/*
step1
list1 = [1,2,4], list2 = [1,3,4]

step2
list1 = [1,1,2,4], list2 = [3,4]

step3
list1 = [1,1,2,3,4], list2 = [4]

step4
list1 = [1,1,2,3,4,4], list2 = []
*/

type para21 struct {
	one *structures.ListNode
	two *structures.ListNode
}

func Test_mergeTwoLists(t *testing.T) {
	// list1 = [1,2,4], list2 = [1,3,4]
	l1_1 := structures.NewLinkedList()
	l1_1.AddNode(structures.NewNode(1))
	l1_1.AddNode(structures.NewNode(2))
	l1_1.AddNode(structures.NewNode(4))

	l2_1 := structures.NewLinkedList()
	l2_1.AddNode(structures.NewNode(1))
	l2_1.AddNode(structures.NewNode(3))
	l2_1.AddNode(structures.NewNode(4))

	// list1 = [4,5], list2 = [7]
	l1_2 := structures.NewLinkedList()
	l1_2.AddNode(structures.NewNode(4))
	l1_2.AddNode(structures.NewNode(5))

	l2_2 := structures.NewLinkedList()
	l2_2.AddNode(structures.NewNode(7))

	qs := []para21{
		{nil, nil},
		{l1_1.Head, nil},
		{nil, l2_1.Head},
		{l1_1.Head, l2_1.Head},
		{l1_2.Head, l2_2.Head},
	}

	for _, v := range qs {
		tmp := MergeTwoLists(v.one, v.two)
		for tmp != nil {
			fmt.Print(tmp.Val, "->")
			tmp = tmp.Next
		}
		fmt.Println("nil")
	}
}
