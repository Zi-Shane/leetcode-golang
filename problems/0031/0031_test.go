package problems

import (
	"fmt"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	nums := []int{1, 3, 2}
	nextPermutation(nums)

	fmt.Println(nums)
}
