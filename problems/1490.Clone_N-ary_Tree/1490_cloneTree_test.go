package problems

import (
	"fmt"
	"leetcode/structures/graph"
	"testing"
)

func TestCloneTree(t *testing.T) {
	node2 := graph.Node{2, nil}

	node6 := graph.Node{6, nil}
	node14 := graph.Node{14, nil}
	node11 := graph.Node{11, []graph.Node{node14}}
	node7 := graph.Node{7, []graph.Node{node11}}
	node3 := graph.Node{3, []graph.Node{node6, node7}}

	node12 := graph.Node{12, nil}
	node8 := graph.Node{8, []graph.Node{node12}}
	node4 := graph.Node{4, []graph.Node{node8}}

	node13 := graph.Node{13, nil}
	node9 := graph.Node{9, []graph.Node{node13}}
	node10 := graph.Node{10, nil}
	node5 := graph.Node{5, []graph.Node{node9, node10}}

	root1 := graph.Node{1, []graph.Node{node2, node3, node4, node5}}

	newRoot := cloneTree(&root1)
	fmt.Println(newRoot.Val)

}
