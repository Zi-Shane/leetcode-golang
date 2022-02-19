package problems

import (
	"leetcode/structures"
)

func addTwoNumbers(l1 *structures.ListNode, l2 *structures.ListNode) *structures.ListNode {
	head := new(structures.ListNode)
	cur := head
	tmp := 0
	for l1 != nil || l2 != nil || tmp > 0 {
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}
		cur.Next = &structures.ListNode{Val: tmp % 10, Next: nil}
		cur = cur.Next
		tmp = tmp / 10
	}

	return head.Next
}
