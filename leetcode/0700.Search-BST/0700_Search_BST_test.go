package leetcode

import (
	"fmt"
	"leetcode/structures/tree"
	"testing"
)

type question700 struct {
	para700
	ans700
}

// para 是参数
// one 代表第一个参数
type para700 struct {
	one []int
	two int
}

// ans 是答案
// one 代表第一个答案
type ans700 struct {
	one []int
}

func Test_Problem700(t *testing.T) {

	qs := []question700{

		{
			// root = [4,2,7,1,3], val = 2
			para700{[]int{4, 2, 7, 1, 3}, 2},
			ans700{[]int{2, 1, 3}},
		},
		{
			// root = [4,2,7,1,3], val = 5
			para700{[]int{4, 2, 7, 1, 3}, 5},
			ans700{[]int{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 700------------------------\n")

	for _, q := range qs {
		_, p := q.ans700, q.para700
		root := SearchBST(structures.Ints2Tree(p.one).GetRoot(), p.two)
		output := structures.Tree2Ints(structures.NewBST(root))
		fmt.Printf("SearchBST():\n【input】: root = %v, val = %v    【output】:%v\n",
			p.one,
			p.two,
			output,
		)
		root = SearchBST2(structures.Ints2Tree(p.one).GetRoot(), p.two)
		output = structures.Tree2Ints(structures.NewBST(root))
		fmt.Printf("SearchBST2():\n【input】: root = %v, val = %v    【output】:%v\n",
			p.one,
			p.two,
			output,
		)
	}
	fmt.Printf("\n\n\n")
}
