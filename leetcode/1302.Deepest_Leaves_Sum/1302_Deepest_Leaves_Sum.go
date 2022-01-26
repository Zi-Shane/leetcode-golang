package leetcode

import (
	"leetcode/structures/tree"
)

func deepestLeavesSum(root *tree.TreeNode) int {
	q := []*tree.TreeNode{root}
	var sum int
	for len(q) > 0 {
		sum = 0
		n := len(q)
		// length of each layer
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:]
			sum += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}

	return sum
}

// Solution 1 - recursive DFS
func deepestLeavesSum1(root *tree.TreeNode) int {
	sum := 0
	deepest := 0
	dfs1(root, 1, &deepest, &sum)

	return sum
}

func dfs1(node *tree.TreeNode, layer int, deepest *int, sum *int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		if layer > *deepest {
			*sum = node.Val
			*deepest = layer
		} else if layer == *deepest {
			*sum += node.Val
		}
	}
	dfs1(node.Left, layer+1, deepest, sum)
	dfs1(node.Right, layer+1, deepest, sum)
}
