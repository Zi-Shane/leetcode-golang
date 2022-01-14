package tree

type Tree struct {
	Root *TreeNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTree() *Tree {
	return new(Tree)
}

func NewTreeNode(num int) *TreeNode {
	t := new(TreeNode)
	t.Val = num
	t.Left, t.Right = nil, nil
	return t
}
