package linkedlist

type LinkedList struct {
	Head *ListNode
	Tail *ListNode
}

// Int ListNode
type ListNode struct {
	// Val  interface{}
	Val  int
	Next *ListNode
}

func NewLinkedList() *LinkedList {
	return new(LinkedList)
}

func NewIntListNode(val int) *ListNode {
	n := new(ListNode)
	n.Val = val
	return n
}

func (l *LinkedList) isEmpty() bool {
	return l.Head == nil
}

func (l *LinkedList) append(n *ListNode) {
	if l.isEmpty() {
		l.Head = n
		l.Tail = n
		return
	}
	l.Tail.Next = n // 將結尾的下個點指向新節點
	l.Tail = n      // 將結尾改成新節點
}

func (l *LinkedList) AddInt(num int) {
	l.append(&ListNode{num, nil})
}
