package leetcode

import (
	"leetcode/structures"
)

func reverseList(head *structures.ListNode) *structures.ListNode {
	var pre *structures.ListNode // Type: `Node` and Value: `nil`
	for head != nil {
		/**
		 * head: a index move each loop
		 * help: store previous head
		 * * hint: `head` must move before `help` change next node
		 */
		help := head
		head = head.Next
		help.Next = pre
		pre = help
	}
	return pre
}

func reverseList2(head *structures.ListNode) *structures.ListNode {
	if head.Next == nil {
		return head
	}
	new_head := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return new_head
}

// func main() {
// 	head := []int{1, 2, 3, 4, 5}
// 	list1 := linkedlist.New()
// 	for i := 0; i < len(head); i++ {
// 		list1.AddItem(linkedlist.NewNode(head[i]))
// 	}
// 	var ans *linkedlist.Node
// 	ans = reverseList(list1.Head)
// 	ans = reverseList2(list1.Head)
// 	fmt.Println(ans)
// }
