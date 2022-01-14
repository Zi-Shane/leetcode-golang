package leetcode

import (
	"leetcode/structures/tree"
)

func rangeSumBST1(root *tree.TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	if root.Val < low {
		return rangeSumBST1(root.Right, low, high)
	}
	if root.Val > high {
		return rangeSumBST1(root.Left, low, high)
	}
	return rangeSumBST1(root.Right, low, high) + rangeSumBST1(root.Left, low, high) + root.Val
}

func rangeSumBST2(root *tree.TreeNode, low int, high int) int {
	res := 0
	a := []tree.TreeNode{}
	a = append(a, *root)
	for len(a) > 0 {
		node := a[0]
		a = a[1:]
		if node.Val < low {
			if node.Right != nil {
				a = append(a, *node.Right)
			}
		} else if node.Val > high {
			if node.Left != nil {
				a = append(a, *node.Left)
			}
		} else {
			if node.Left != nil {
				a = append(a, *node.Left)
			}
			if node.Right != nil {
				a = append(a, *node.Right)
			}
			res += node.Val
		}
	}
	return res
}
