package problems

import (
	"leetcode/structures"
	"testing"
)

func Test_Problem2(t *testing.T) {
	// l1 = [2, 4, 3], l2 = [5, 6, 4]
	// l1 := &structures.ListNode{Val: 2, Next: nil}
	// l1.Next = &structures.ListNode{Val: 4, Next: nil}
	// l1.Next.Next = &structures.ListNode{Val: 3, Next: nil}
	// l2 := &structures.ListNode{Val: 5, Next: nil}
	// l2.Next = &structures.ListNode{Val: 6, Next: nil}
	// l2.Next.Next = &structures.ListNode{Val: 4, Next: nil}

	// l1 = [9999], l2 = [999]
	l1 := &structures.ListNode{Val: 9, Next: nil}
	l1.Next = &structures.ListNode{Val: 9, Next: nil}
	l1.Next.Next = &structures.ListNode{Val: 9, Next: nil}
	l1.Next.Next.Next = &structures.ListNode{Val: 9, Next: nil}
	l2 := &structures.ListNode{Val: 9, Next: nil}
	l2.Next = &structures.ListNode{Val: 9, Next: nil}
	l2.Next.Next = &structures.ListNode{Val: 9, Next: nil}

	ans := addTwoNumbers(l1, l2)
	print(ans.Val)
	print("Done.")
}
