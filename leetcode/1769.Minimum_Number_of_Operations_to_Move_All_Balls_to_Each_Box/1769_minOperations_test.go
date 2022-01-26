package leetcode

import "testing"

func TestMinOperations(t *testing.T) {
	testCases := []struct {
		boxes string
		ans   []int
	}{
		{
			boxes: "110",
			ans:   []int{1, 1, 3},
		},
		{
			boxes: "001011",
			ans:   []int{11, 8, 5, 4, 3, 4},
		},
	}
	for i, tC := range testCases {
		res := minOperations(tC.boxes)
		for j, v := range res {
			if v != tC.ans[j] {
				t.Errorf("TestCase%v is wrong!! res=%v, ans=%v", i, res, tC.ans)
			}
		}
	}
}
