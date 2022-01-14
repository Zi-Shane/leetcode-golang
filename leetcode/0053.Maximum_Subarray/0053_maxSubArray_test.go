package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	// https://leetcode.com/submissions/detail/604888155/testcase/
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	res := maxSubArray2(nums)
	fmt.Println(res)
}
