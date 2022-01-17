package leetcode

import (
	"fmt"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	testCases := []struct {
		candies      []int
		extraCandies int
		ans          []bool
	}{
		{
			candies:      []int{2, 3, 5, 1, 3},
			extraCandies: 3,
			ans:          []bool{true, true, true, false, true},
		},
	}
	for _, tC := range testCases {
		res := kidsWithCandies(tC.candies, tC.extraCandies)
		fmt.Println(res)
		for i, v := range tC.ans {
			if res[i] != v {
				t.Errorf("Wrong answer!!")
				break
			}
		}
	}
}
