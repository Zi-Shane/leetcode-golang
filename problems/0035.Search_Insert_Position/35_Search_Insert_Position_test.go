package problems

import (
	"fmt"
	"testing"
)

type question35 struct {
	para35
	ans35
}

type para35 struct {
	one []int
	two int
}

type ans35 struct {
	one int
}

func Test_Problem35(t *testing.T) {

	qs := []question35{

		// nums = [1,3,5,6], target = 5
		{
			para35{[]int{1, 3, 5, 6}, 5},
			ans35{2},
		},

		// nums = [1,3,5,6], target = 2
		{
			para35{[]int{1, 3, 5, 6}, 2},
			ans35{1},
		},

		// nums = [1,3,5,6], target = 7
		{
			para35{[]int{1, 3, 5, 6}, 7},
			ans35{4},
		},

		// nums = [1,3,5,6], target = 0
		{
			para35{[]int{1, 3, 5, 6}, 0},
			ans35{0},
		},

		// nums = [1], target = 0
		{
			para35{[]int{1}, 0},
			ans35{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 26------------------------\n")

	for i, q := range qs {
		_, p := q.ans35, q.para35
		ans := searchInsert(p.one, p.two)
		fmt.Printf("【input】:%v %v   【output】:%v\n", p.one, p.two, ans)
		if ans != q.ans35.one {
			t.Errorf("Test case %v, wrong answer!", i+1)
		}
	}
	fmt.Printf("\n\n\n")
}
