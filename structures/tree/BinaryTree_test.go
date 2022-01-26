package tree

import (
	"fmt"
	"testing"
)

func Test_BinaryTree(t *testing.T) {
	input := []int{4, 2, 7, 1, 3}
	tree := Ints2Tree(input)
	fmt.Print(tree.Root.Val)
	ints := Tree2Ints(tree)
	fmt.Print(ints)
}
