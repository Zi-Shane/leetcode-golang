package problems

import (
	"fmt"
	"testing"
)

func TestSuffle(t *testing.T) {
	testCases := []struct {
		arr []int
		n   int
		ans []int
	}{
		{
			arr: []int{2, 5, 1, 3, 4, 7},
			n:   3,
			ans: []int{2, 3, 5, 4, 1, 7},
		},
	}
	for _, tC := range testCases {
		res := shuffle(tC.arr, tC.n)
		for i := 0; i < len(res); i++ {
			if tC.ans[i] != res[i] {
				fmt.Printf("i=%v is wrong!", i)
			}
		}
	}
}
