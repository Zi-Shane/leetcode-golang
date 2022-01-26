package leetcode

import (
	"leetcode/structures/tree"
	"testing"
)

func Test(t *testing.T) {
	tree1 := tree.NewTree()
	// root
	tree1.Root = tree.NewTreeNode(1)
	// l2
	tree1.Root.Left = tree.NewTreeNode(2)
	tree1.Root.Right = tree.NewTreeNode(3)
	// l3
	tree1.Root.Left.Left = tree.NewTreeNode(4)
	tree1.Root.Left.Right = tree.NewTreeNode(5)
	tree1.Root.Right.Right = tree.NewTreeNode(6)
	// l4
	tree1.Root.Left.Left.Left = tree.NewTreeNode(7)
	tree1.Root.Right.Right.Right = tree.NewTreeNode(8)

	testCases := []struct {
		tree *tree.TreeNode
		ans  int
	}{
		{
			tree: tree1.Root,
			ans:  15,
		},
	}
	for i, tC := range testCases {
		res := deepestLeavesSum1(tC.tree)
		if res != tC.ans {
			t.Errorf("TestCase %v is Wrong!! res=%v, ans=%v", i, res, tC.ans)
		}
	}
}
