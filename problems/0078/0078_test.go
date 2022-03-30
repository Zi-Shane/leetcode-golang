package problems

import (
	"fmt"
	"testing"
)

func TestSubsets(t *testing.T) {
	nums := []int{9, 0, 3, 5, 7}
	// nums := []int{0}

	res := subsets(nums)

	fmt.Println(res)
}
