package leetcode

import (
	"leetcode/structures"
)

func SearchBST(root *structures.TreeNode, val int) *structures.TreeNode {
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

func SearchBST2(root *structures.TreeNode, val int) *structures.TreeNode {
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
