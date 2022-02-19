package problems

import (
	"leetcode/structures/tree"
)

func constructMaximumBinaryTree(nums []int) *tree.TreeNode {
	stack := []*tree.TreeNode{}
	for _, n := range nums {
		node := &tree.TreeNode{n, nil, nil}
		var last *tree.TreeNode
		for len(stack) > 0 && stack[len(stack)-1].Val < n {
			last = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		node.Left = last

		if len(stack) > 0 && stack[len(stack)-1].Val > n {
			stack[len(stack)-1].Right = node
		}

		stack = append(stack, node)
	}

	return stack[0]
}
