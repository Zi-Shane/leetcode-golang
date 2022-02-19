package problems

import (
	"fmt"
	"testing"
)

func TestBuildArray(t *testing.T) {
	testCases := []struct {
		nums []int
		ans  []int
	}{
		{
			nums: []int{0, 2, 1, 5, 3, 4},
			ans:  []int{0, 1, 2, 4, 5, 3},
		},
		{
			nums: []int{5, 0, 1, 2, 3, 4},
			ans:  []int{4, 5, 0, 1, 2, 3},
		},
	}
	for _, tC := range testCases {
		fmt.Println("nums: ", tC.nums)
		res := buildArray(tC.nums)
		for i := 0; i < len(tC.nums); i++ {
			if res[i] != tC.ans[i] {
				fmt.Println("worng answer")
				break
			}
		}
		fmt.Println("ans: ", res)
	}
}
