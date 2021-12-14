package structures

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
	Tail *ListNode
}

// List2Ints convert List to []int
func List2Ints(head *ListNode) []int {
	// 链条深度限制，链条深度超出此限制，会 panic
	limit := 100

	times := 0

	res := []int{}
	for head != nil {
		times++
		if times > limit {
			msg := fmt.Sprintf("链条深度超过%d，可能出现环状链条。请检查错误，或者放宽 l2s 函数中 limit 的限制。", limit)
			panic(msg)
		}

		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

// Ints2List convert []int to List
func Ints2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}

func NewLinkedList() *LinkedList {
	return new(LinkedList)
}

func NewNode(val int) *ListNode {
	n := new(ListNode)
	n.Val = val
	return n
}

func (l *LinkedList) AddNode(n *ListNode) {
	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
		l.Tail = l.Tail.Next
	}
}

// func (l *LinkedList) AddItem(n int) {
// 	if l.Head == nil {
// 		l.Head = NewNode(n)
// 		l.Tail = l.Head
// 	} else {
// 		l.Tail.Next = NewNode(n)
// 		l.Tail = l.Tail.Next
// 	}
// }

func (l *LinkedList) RemoveItem() *ListNode {
	if l.Head == nil {
		return nil
	}
	var pre_node *ListNode
	for current := l.Head; current != l.Tail; current = current.Next {
		pre_node = current
	}
	tmp := pre_node.Next
	pre_node.Next = nil
	l.Tail = pre_node
	return tmp
}

func (node ListNode) printValue() {
	fmt.Print(node.Val)
}

func (node ListNode) getNext() ListNode {
	return *node.Next
}
