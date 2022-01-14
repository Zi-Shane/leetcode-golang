package leetcode

import (
	"fmt"
	"leetcode/structures/tree"
	"testing"
)

func Test(t *testing.T) {
	bst1 := tree.NewTree()
	bst1.Root = tree.NewTreeNode(10)
	bst1.Root.Left = tree.NewTreeNode(5)
	bst1.Root.Left.Left = tree.NewTreeNode(3)
	bst1.Root.Left.Right = tree.NewTreeNode(7)
	bst1.Root.Right = tree.NewTreeNode(15)
	bst1.Root.Right.Right = tree.NewTreeNode(18)

	bst2 := tree.NewTree()
	// layer1
	bst2.Root = tree.NewTreeNode(10)
	// layer2
	bst2.Root.Left = tree.NewTreeNode(5)
	bst2.Root.Right = tree.NewTreeNode(15)
	// layer3
	bst2.Root.Left.Left = tree.NewTreeNode(3)
	bst2.Root.Left.Right = tree.NewTreeNode(7)
	bst2.Root.Right.Left = tree.NewTreeNode(13)
	bst2.Root.Right.Right = tree.NewTreeNode(18)
	// layer4
	bst2.Root.Left.Left.Left = tree.NewTreeNode(1)
	bst2.Root.Left.Right.Left = tree.NewTreeNode(6)

	r1 := 0
	r1 = rangeSumBST1(bst1.Root, 7, 15)
	fmt.Println(r1)
	r1 = rangeSumBST1(bst2.Root, 6, 10)
	fmt.Println(r1)

	r2 := 0
	r2 = rangeSumBST2(bst1.Root, 7, 15)
	fmt.Println(r2)
	r2 = rangeSumBST2(bst2.Root, 6, 10)
	fmt.Println(r2)
}
