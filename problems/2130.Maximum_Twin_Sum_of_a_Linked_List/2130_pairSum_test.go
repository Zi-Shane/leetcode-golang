package problems

import (
	"leetcode/structures/linkedlist"
	"testing"
)

func TestPairSum(t *testing.T) {
	root := linkedlist.NewIntListNode(5)
	root.Next = linkedlist.NewIntListNode(4)
	root.Next.Next = linkedlist.NewIntListNode(2)
	root.Next.Next.Next = linkedlist.NewIntListNode(1)

	testCases := []struct {
		node *linkedlist.ListNode
		ans  int
	}{
		{
			node: root,
			ans:  6,
		},
	}
	for i, tC := range testCases {
		res := pairSum(tC.node)
		if tC.ans != res {
			t.Errorf("TestCase %d, Wrong answer res=%d, ans=%d", i, res, tC.ans)
		}
	}
}
