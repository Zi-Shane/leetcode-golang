package problems

import (
	"leetcode/structures/linkedlist"
)

func pairSum(head *linkedlist.ListNode) int {
	slow, fast := head, head
	n := 0
	for fast != nil && slow.Next != nil {
		n++
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *linkedlist.ListNode
	if slow.Next == nil {
		prev = slow
	} else {
		for slow != nil {
			tmp := slow.Next
			slow.Next = prev
			prev = slow
			slow = tmp
		}
	}

	maxSum := 0
	head2 := prev
	for head != nil && head2 != nil {
		sum := head.Val + head2.Val
		if maxSum < sum {
			maxSum = sum
		}
		head = head.Next
		head2 = head2.Next
	}

	return maxSum
}
