package leetcode

import (
	"fmt"
	"testing"
)

func TestDotProduct(t *testing.T) {
	testCases := []struct {
		nums1 []int
		nums2 []int
		ans   int
	}{
		{
			nums1: []int{1, 0, 0, 2, 3},
			nums2: []int{0, 3, 0, 4, 0},
			ans:   8,
		},
		{
			nums1: []int{0, 1, 0, 0, 0},
			nums2: []int{0, 0, 0, 0, 2},
			ans:   0,
		},
		{
			nums1: []int{0, 1, 0, 0, 2, 0, 0},
			nums2: []int{1, 0, 0, 0, 3, 0, 4},
			ans:   6,
		},
	}
	for i, tC := range testCases {
		vec1 := NewSparseVector(tC.nums1)
		vec2 := NewSparseVector(tC.nums2)
		res := vec1.dotProduct(vec2)
		if res != tC.ans {
			t.Errorf("Case %v: Wrong Answer (res=%v)\n", i, res)
		} else {
			fmt.Printf("Case %v: Correct\n", i)
		}
	}
}
