package leetcode

import "leetcode/structures/tree"

func preorderTraversal(root *tree.TreeNode) []int {
	output := []int{}
	if root == nil {
		return output
	}
	output = append(output, root.Val)
	output = append(output, preorderTraversal(root.Left)...)
	output = append(output, preorderTraversal(root.Right)...)
	return output
}
