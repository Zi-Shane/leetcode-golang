package problems

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	// nums := []int{4, 5, 6, 7, 0, 1, 2}
	nums := []int{0}
	target := 1

	res := search2(nums, target)

	fmt.Println(res)
}
