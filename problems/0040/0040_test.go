package problems

import (
	"fmt"
	"testing"
)

func TestCombine(t *testing.T) {
	res := combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)

	fmt.Println(res)
}
