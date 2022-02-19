package problems

import (
	"leetcode/structures/tree"
)

func SearchBST(root *tree.TreeNode, val int) *tree.TreeNode {
	if root == nil {
		return nil
	}
	if val < root.Val {
		return SearchBST(root.Left, val)
	} else if val > root.Val {
		return SearchBST(root.Right, val)
	}
	// val == root.Val
	return root
}

func SearchBST2(root *tree.TreeNode, val int) *tree.TreeNode {
	for root != nil {
		if val == root.Val {
			return root
		} else if val < root.Val {
			return SearchBST2(root.Left, val)
		} else if val > root.Val {
			return SearchBST2(root.Right, val)
		}
	}
	return nil
}
