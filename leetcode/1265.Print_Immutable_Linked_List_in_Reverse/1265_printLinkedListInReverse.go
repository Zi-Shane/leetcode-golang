package leetcode

import (
	"fmt"
	"leetcode/structures/linkedlist"
	Stack "leetcode/structures/stack"
)

func PrintLinkedListInReverse1(head *linkedlist.Node) {
	if head == nil {
		return
	}
	PrintLinkedListInReverse1(head.Next)
	fmt.Print(head.Val)
}

func PrintLinkedListInReverse2(head *linkedlist.Node) {
	count := 0
	tmp := head
	for head != nil {
		count++
		head = head.Next
	}
	for i := count; i > 0; i-- {
		head = tmp
		for j := 0; j < i-1; j++ {
			head = head.Next
		}
		fmt.Print(head.Val)
	}
}

// Follow up: Constant space complexity
func PrintLinkedListInReverse3(head *linkedlist.Node) {
	s := Stack.NewStack()
	for head != nil {
		s.PushInt(head.Val)
		head = head.Next
	}
	for !s.IsEmpty() {
		fmt.Print(s.PopInt())
	}
}

// Follow up: Linear time complexity and less than linear space complexity
func PrintLinkedListInReverse4(head *linkedlist.Node) {
	if head == nil {
		return
	}
	if head.Next == nil {
		fmt.Print(head.Val)
		return
	}
	PrintLinkedListInReverse4(head.Next.Next)
	PrintReverse(head)
}

func PrintReverse(head *linkedlist.Node) {
	fmt.Print(head.Next.Val)
	fmt.Print(head.Val)
}
