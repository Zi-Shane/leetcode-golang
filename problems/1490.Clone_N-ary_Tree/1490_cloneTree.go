package problems

import "leetcode/structures/graph"

func cloneTree(root *graph.Node) *graph.Node {
	newNode := dfs1490_2(root)
	return newNode
}

func dfs1490(node *graph.Node) *graph.Node {
	newNode := &graph.Node{node.Val, []graph.Node{}}

	// stop
	if len(node.Children) == 0 {
		return newNode
	}

	// traverse
	for _, n := range node.Children {
		newNode.Children = append(newNode.Children, *dfs1490(&n))
	}

	return newNode
}

func dfs1490_2(node *graph.Node) *graph.Node {

	// stop
	if node == nil {
		return nil
	}

	// newNode := &graph.Node{Val: node.Val, Children: []graph.Node{}}
	newNode := new(graph.Node)
	newNode.Val = node.Val

	// traverse
	for _, n := range node.Children {
		newNode.Children = append(newNode.Children, *dfs1490_2(&n))
	}

	return newNode
}
