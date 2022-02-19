package problems

import (
	"leetcode/structures"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func MergeTwoLists(list1 *structures.ListNode, list2 *structures.ListNode) (res *structures.ListNode) {
	if list1 == nil {
		res = list2
	} else if list2 == nil {
		res = list1
	} else {
		if list1.Val < list2.Val {
			res = list1
			merge2(list1, list2)
		} else {
			res = list2
			merge2(list2, list1)
		}
	}

	return res
}

func merge(list1 *structures.ListNode, list2 *structures.ListNode) {
	for list2 != nil && list1.Next != nil {
		if list1.Next.Val > list2.Val {
			// connect
			tmp := list1.Next
			list1.Next = list2    // connect list2 to list1
			list2 = list2.Next    // update list2
			list1.Next.Next = tmp // connect new node to old list1
			list1 = list1.Next    // update list1
		} else {
			// goto next
			list1 = list1.Next
		}
	}
	if list2 != nil {
		list1.Next = list2
	}
}

func merge2(list1 *structures.ListNode, list2 *structures.ListNode) {
	// list1 = tail node
	// list2 = node is not added
	for list2 != nil {
		if list1.Next == nil {
			list1.Next = list2
			break
		}
		tmp := list1.Next
		list1.Next = list2
		for list2 != nil && list2.Val < tmp.Val {
			list2 = list2.Next
			list1 = list1.Next
		}
		list1.Next = tmp
		list1 = list1.Next
	}

}

func MergeTwoLists2(list1 *structures.ListNode, list2 *structures.ListNode) (res *structures.ListNode) {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = MergeTwoLists2(list1.Next, list2)
		return list1
	} else {
		list2.Next = MergeTwoLists2(list1, list2.Next)
		return list2
	}
}
