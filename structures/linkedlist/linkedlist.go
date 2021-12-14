package linkedlist

type LinkedList struct {
	Head *Node
	Tail *Node
}

// Int Node
type Node struct {
	// Val  interface{}
	Val  int
	Next *Node
}

func NewLinkedList() *LinkedList {
	return new(LinkedList)
}

func NewIntListNode(val int) *Node {
	n := new(Node)
	n.Val = val
	return n
}

func (l *LinkedList) isEmpty() bool {
	return l.Head == nil
}

func (l *LinkedList) append(n *Node) {
	if l.isEmpty() {
		l.Head = n
		l.Tail = n
		return
	}
	l.Tail.Next = n // 將結尾的下個點指向新節點
	l.Tail = n      // 將結尾改成新節點
}

func (l *LinkedList) AddInt(num int) {
	l.append(&Node{num, nil})
}
