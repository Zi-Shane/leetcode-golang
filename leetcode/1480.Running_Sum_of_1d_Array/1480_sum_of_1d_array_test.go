package leetcode

import "testing"

func TestRunningSum(t *testing.T) {
	testCases := []struct {
		nums []int
		ans  []int
	}{
		{
			nums: []int{1, 2, 3, 4},
			ans:  []int{1, 3, 6, 10},
		},
		{
			nums: []int{1, 1, 1, 1, 1},
			ans:  []int{1, 2, 3, 4, 5},
		},
	}
	for _, tC := range testCases {
		res := runningSum(tC.nums)
		for i := 0; i < len(tC.ans); i++ {
			if res[i] != tC.ans[i] {
				t.Error("Wrong Answer!")
				break
			}
		}
	}
}
