package problems

import "testing"

func TestDecompressRLElist(t *testing.T) {
	testCases := []struct {
		nums []int
		ans  []int
	}{
		{
			nums: []int{1, 2, 3, 4},
			ans:  []int{2, 4, 4, 4},
		},
	}
	for _, tC := range testCases {
		res := decompressRLElist(tC.nums)
		for i, v := range tC.ans {
			if res[i] != v {
				t.Errorf("TestCase %v is Wrong!! res=%v, ans=%v", i, res, tC.ans)
			}
		}
	}
}
