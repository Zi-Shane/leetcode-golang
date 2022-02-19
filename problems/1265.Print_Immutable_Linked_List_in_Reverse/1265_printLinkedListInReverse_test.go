package problems

import (
	"fmt"
	"leetcode/structures/linkedlist"

	// "leetcode/structures"
	"testing"
)

func TestPrintLinkedListInReverse(t *testing.T) {
	l := linkedlist.NewLinkedList()
	l.AddInt(1)
	l.AddInt(2)
	l.AddInt(3)
	l.AddInt(4)

	fmt.Println("PrintLinkedListInReverse1: ")
	PrintLinkedListInReverse1(l.Head)

	fmt.Println("\nPrintLinkedListInReverse2: ")
	PrintLinkedListInReverse2(l.Head)

	fmt.Println("\nPrintLinkedListInReverse3: ")
	PrintLinkedListInReverse3(l.Head)

	fmt.Println("\nPrintLinkedListInReverse4: ")
	PrintLinkedListInReverse4(l.Head)

	fmt.Println()
}
