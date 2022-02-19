package problems

import (
	"fmt"
	"testing"
)

func TestSumOddLengthSubarrays(t *testing.T) {
	ans := sumOddLengthSubarrays([]int{1, 4, 2, 5, 3})
	fmt.Println(ans)
}
