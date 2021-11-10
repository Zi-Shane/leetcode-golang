package leetcode

import (
	"fmt"
	"leetcode/structures"
	"testing"
)

type question144 struct {
	para144
	ans144
}

// para 是参数
// one 代表第一个参数
type para144 struct {
	one *structures.TreeNode
}

// ans 是答案
// one 代表第一个答案
type ans144 struct {
	one []int
}

func Test_Problem144(t *testing.T) {
	tree := structures.NewEmptyBST()
	tree.Root = &structures.TreeNode{1, nil, nil}
	tree.Root.Right = &structures.TreeNode{2, nil, nil}
	tree.Root.Right.Left = &structures.TreeNode{3, nil, nil}

	qs := []question144{

		{
			// root = [1, nil, 2, 3]
			para144{tree.Root},
			ans144{[]int{2, 1, 3}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 144------------------------\n")

	for _, q := range qs {
		_, p := q.ans144, q.para144
		output := preorderTraversal(p.one)
		fmt.Printf("SearchBST():\n【input】: root = %v    【output】:%v\n",
			p.one,
			output,
		)
	}
	fmt.Printf("\n\n\n")
}
