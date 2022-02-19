package problems

import "testing"

func TestSmallerNumbersThanCurrent(t *testing.T) {
	testCases := []struct {
		nums []int
		ans  []int
	}{
		{
			nums: []int{8, 1, 2, 2, 3},
			ans:  []int{4, 0, 1, 1, 3},
		},
	}
	for _, tC := range testCases {
		res := smallerNumbersThanCurrent(tC.nums)
		for i, v := range tC.ans {
			if res[i] != v {
				t.Errorf("TestCase %v is Wrong!! res=%v, ans=%v", i, res, tC.ans)
			}
		}
	}
}
